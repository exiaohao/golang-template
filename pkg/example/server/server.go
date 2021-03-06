package server

import (
	"context"
	"fmt"
	"github.com/exiaohao/golang-template/pkg/db"
	"github.com/exiaohao/golang-template/pkg/example/controller"
	"github.com/exiaohao/golang-template/pkg/example/recovery"
	"github.com/exiaohao/golang-template/pkg/example/router"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"time"
)

type InitOptions struct {
	Port       uint16
	Address    string
	KubeConfig string
}

type HttpServer struct {
	Ctx     context.Context
	Port    uint16
	Address string

	//kubeClient	*kubernetes.Clientset
}

func (hs *HttpServer) Initialize(opts InitOptions) {
	//var err error

	hs.Address = opts.Address
	hs.Port = opts.Port

	// Optional: Initialize kubernetes clients if required.
	//if hs.kubeClient, err = common.InitializeKubeClient(opts.KubeConfig); err != nil {
	//	glog.Fatalf("initialize kubeclient failed by using %s: %v", opts.KubeConfig, err)
	//}
}

func (hs *HttpServer) Run(stopCh <-chan struct{}) {
	glog.Infof("Server started, Listen %s:%d", hs.Address, hs.Port)

	r := gin.New()
	r.Use(recovery.Recovery(controller.RecoveryHandler))

	router.RegisterRouter(r)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", hs.Address, hs.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Fatalf("Error occured: %s", err)
		}
	}()

	<-stopCh
	glog.Info("stopping server...")
	db.CloseDB()
	if err := srv.Shutdown(hs.Ctx); err != nil {
		glog.Fatal("Server Shutdown:", err)
	}
	time.Sleep(2 * time.Second)
	glog.Fatal("Server exiting")
}
