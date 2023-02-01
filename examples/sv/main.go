package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/xhaoh94/gox"

	"github.com/xhaoh94/gox/engine/helper/codechelper"
	"github.com/xhaoh94/gox/engine/network/service/kcp"
	"github.com/xhaoh94/gox/engine/network/service/ws"
	"github.com/xhaoh94/gox/examples/sv/game"
	"github.com/xhaoh94/gox/examples/sv/mods"
)

func main() {

	// var sid uint
	// flag.UintVar(&sid, "sid", uint(strhelper.StringToHash(commonhelper.GetUUID())), "uuid")
	// var sType, iAddr, oAddr string
	// flag.StringVar(&sType, "type", "all", "服务类型")
	// flag.StringVar(&iAddr, "iAddr", "127.0.0.1:10001", "服务地址")
	// flag.StringVar(&oAddr, "oAddr", "127.0.0.1:10002", "服务地址")
	appConfPath := *flag.String("appConfPath", "", "grpc服务地址")
	flag.Parse()
	if appConfPath == "" {
		log.Fatalf("需要启动配置文件路径")
	}
	engine := gox.NewEngine(appConfPath)
	game.Engine = engine
	engine.SetModule(new(mods.MainModule))
	engine.SetInteriorService(new(kcp.KService), codechelper.Json)
	engine.SetOutsideService(new(ws.WService), codechelper.Json)
	engine.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan
	engine.Shutdown()
	os.Exit(1)
}
