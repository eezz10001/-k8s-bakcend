package configs

import (
	"k8s-bakcend/src/services"
)

//注入 回调handler
type K8sHandler struct {}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}
// deployment handler
func(this *K8sHandler) DepHandlers() *services.DepHandler {
	return &services.DepHandler{}
}
func(this *K8sHandler) PodHandlers() *services.PodHandler {
	return &services.PodHandler{}
}

func (this *K8sHandler)NsHandlers() *services.NsHandler  {
	return &services.NsHandler{}
}