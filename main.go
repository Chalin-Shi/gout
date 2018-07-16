package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/robfig/cron"

	"gout/libs/setting"
	"gout/routers"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.Port)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	c := cron.New()
	spec := "*/5 * * * * ?"
	i := 0
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select {}
}
