package scene

import (
	"context"

	"github.com/xhaoh94/gox"
	"github.com/xhaoh94/gox/engine/network"
	"github.com/xhaoh94/gox/engine/xlog"
	"github.com/xhaoh94/gox/examples/netpack"
)

type (
	Unit struct {
		network.Actor
		Id uint
	}
)

func newUnit(id uint) *Unit {
	unit := &Unit{Id: id}
	unit.OnInit()
	gox.ActorSystem.Add(unit) //添加到Actor
	return unit
}

func (unit *Unit) ActorID() uint32 {
	return uint32(unit.Id)
}

func (unit *Unit) OnInit() {
	unit.AddActorFn(unit.SayHello) //添加Actor回调
}

func (unit *Unit) SayHello(ctx context.Context, req *netpack.L2S_SayHello) *netpack.S2L_SayHello {
	xlog.Debug("收到sayHello:%s", req.Txt)
	return &netpack.S2L_SayHello{BackTxt: req.Txt + "返回"}
}
