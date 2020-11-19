package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zuanren/dao"
	"zuanren/models"
	"zuanren/routes"
	"zuanren/utils"
)

func main() {

	// 初始化配置变量
	utils.GetConf()

	// MySQL连接池
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	// 结构体与数据表关联
	_ = dao.DB.AutoMigrate(&models.XtpAccounts{}) // xtp_accounts表
	_ = dao.DB.AutoMigrate(&models.Content{})     // content表

	// 注册路由
	router := gin.Default()
	routes.SetupV0Router(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 定时任务
	//go cron.Check()

	// 优雅退出
	go func() {
		// 服务监听启动
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	defer dao.Close() // 关闭数据库连接
	log.Println("Server exiting")
}
