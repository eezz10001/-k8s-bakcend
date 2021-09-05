package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-bakcend/src/services"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClinet *kubernetes.Clientset `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl) Name() string {
	return "DeploymentCtl"
}

func (this *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/deployments", this.GetList)
}

func (this *DeploymentCtl) GetList(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns","default")
	return  this.DepService.ListAll(ns)
}
