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
)

/**
优雅停机
*/
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {

		//回调方法设置一个超时5s
		time.Sleep(10 * time.Second)
		c.String(200, "hello licslan")
	})
	srv := &http.Server{
		Addr:    ":8085",
		Handler: r,
	}
	//将其放入到一个协程里面
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//退出信号channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//阻塞channel
	<-quit
	log.Println("shutdown server.....")
	//创建一个超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown :", err)
	}
	log.Println("server exiting ......")
}
