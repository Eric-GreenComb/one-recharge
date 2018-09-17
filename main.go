package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/one-recharge/config"
	"github.com/Eric-GreenComb/one-recharge/handler"
	"github.com/Eric-GreenComb/one-recharge/persist"
)

var (
	g errgroup.Group
)

func main() {
	if config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDatabase()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)
	}

	// api
	r1 := router.Group("/api/v1")
	{
		// 兑换比例
		r1.POST("/rate/create", handler.CreateRechargeRate)
		r1.GET("/rate", handler.RechargeRateInfo)
		r1.GET("/rate/list/:page/:limit", handler.ListRechargeRate)

	}

	for _, _port := range config.Server.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return server.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
