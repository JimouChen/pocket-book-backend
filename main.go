package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"pocket-book/comm"
	"pocket-book/dao/mysql"
	"pocket-book/routes"
	"syscall"
	"time"
)

func main() {

	// 初始化读取配置
	if err := comm.InitViperCfg(); err != nil {
		return
	}
	cfgLoader := comm.CfgLoader

	// 初始化日志
	if err := comm.InitLogger(); err != nil {
		return
	}
	log := comm.Logger

	//- 初始化mysql连接
	if err := mysql.InitMysqlCfg(); err != nil {
		log.Error().Msgf("init mysql failed: %s!", err.Error())
		return
	}
	defer mysql.Close()

	// 初始化路由
	r := routes.Init()
	log.Info().Msg("Router init...")

	env := cfgLoader.GetString("app.env") + ".url"
	//if err := r.Run(cfgLoader.GetString(env)); err != nil {
	//	log.Error().Msg("App Run Error")
	//	return
	//}

	srv := &http.Server{
		Addr:    cfgLoader.GetString(env),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info().Msgf("listen: %s\n", err.Error())
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Info().Msg("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Info().Msgf("Server Shutdown: ", err.Error())
	}

	log.Info().Msg("Server exiting")
}
