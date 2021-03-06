package cron

import (
	"github.com/robfig/cron"
	"log"
	"time"
	srv "zuanren/service"
)

func Check() {
	log.Println("Starting...")

	c := cron.New()
	// 每天9点执行一次定时任务
	_ = c.AddFunc("0 0 9 * * *", func() {
		log.Println("Run Check Count...")
		srv.SendCheckCount()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
