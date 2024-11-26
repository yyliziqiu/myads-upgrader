package main

import (
	"io"
	"log"
	"os"

	"github.com/kardianos/service"
	"github.com/yyliziqiu/zlib/zlog"

	"myads-upgrater/conf"
	"myads-upgrater/svc"
)

func main() {
	setLogger()
	doService()
}

func setLogger() {
	config := zlog.Config{
		Console: true,
		Path:    conf.BasePath("logs"),
		Name:    "upgrader",
	}

	err := zlog.Init(config)
	if err != nil {
		log.Fatalln(err)
	}
}

func doService() {
	ws, err := service.New(&Upgrader{}, &service.Config{
		Name:        "MyadsUpgrader",
		DisplayName: "Myads Upgrader",
		Description: "Myads Upgrader",
	})
	if err != nil {
		zlog.Errorf("New service failed, error: %v", err)
		return
	}

	if len(os.Args) > 1 {
		err = service.Control(ws, os.Args[1])
	} else {
		err = ws.Run()
	}
	if err != nil {
		zlog.Errorf("Run service failed, error: %v", err)
		return
	}
}

type Upgrader struct {
	writer io.Writer
}

func (m *Upgrader) Start(_ service.Service) error {
	go svc.Boot()
	return nil
}

func (m *Upgrader) Stop(_ service.Service) error {
	return nil
}
