package configs

import (
	"k8s-bakcend/src/services"
	"k8s.io/client-go/kubernetes"
)

type K8sMaps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}
func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}
//初始化 deploymentmap
func(this *K8sMaps) InitDepMap() *services.DeploymentMap {
	return &services.DeploymentMap{}
}


func(this *K8sMaps) InitPodMap() *services.PodMapStruct {
	return &services.PodMapStruct{}
}
