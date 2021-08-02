package kcp

import (
	"sync"
	"time"

	"github.com/xhaoh94/gox/app"
	"github.com/xhaoh94/gox/engine/network/service"
	"github.com/xhaoh94/gox/engine/xlog"
	"github.com/xtaci/kcp-go/v5"
)

var channelPool *sync.Pool

func init() {
	channelPool = &sync.Pool{
		New: func() interface{} {
			return &KChannel{}
		},
	}
}

type (
	//KChannel TCP信道
	KChannel struct {
		service.Channel
		connGuard sync.RWMutex
		conn      *kcp.UDPSession
	}
)

func (k *KChannel) init(conn *kcp.UDPSession) {
	k.conn = conn
	k.Init(k.write, k.Conn().RemoteAddr().String(), k.Conn().LocalAddr().String())
}

//Conn 获取通信体
func (k *KChannel) Conn() *kcp.UDPSession {
	k.connGuard.RLock()
	defer k.connGuard.RUnlock()
	return k.conn
}

//Start 开启异步接收数据
func (k *KChannel) Start() {
	k.Wg.Add(1)
	go func() {
		defer k.OnStop()
		k.Wg.Wait()
	}()
	k.IsRun = true
	go k.recvAsync()
}
func (k *KChannel) recvAsync() {
	defer k.Wg.Done()
	readTimeout := app.GetAppCfg().Network.ReadTimeout
	if readTimeout > 0 {
		if err := k.Conn().SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
			xlog.Info("kpc addr[%s] 接受数据超时[%v]", k.RemoteAddr(), err)
			k.Stop()
		}
	}
	for k.Conn() != nil && k.IsRun {
		k.Read(k.Conn(), k.Stop)
		if k.IsRun && readTimeout > 0 {
			if err := k.Conn().SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
				xlog.Info("kpc addr[%s] 接受数据超时[%v]", k.RemoteAddr(), err)
				k.Stop()
			}
		}
	}
}

func (k *KChannel) write(buf []byte) {
	_, err := k.Conn().Write(buf)
	if err != nil {
		xlog.Error("kcp addr[%s]信道写入失败: [%v]", k.RemoteAddr(), err)
	}
}

//Stop 停止信道
func (k *KChannel) Stop() {
	if !k.IsRun {
		return
	}
	k.Conn().Close()
	k.IsRun = false
}

//OnStop 关闭
func (k *KChannel) OnStop() {
	k.Channel.OnStop()
	k.conn = nil
	channelPool.Put(k)
}
