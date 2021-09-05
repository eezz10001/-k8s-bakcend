package configs

import (
	"k8s-bakcend/src/services"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)
type K8sConfig struct {
	DepHandler *services.DepHandler `inject:"-"`
	PodHandler *services.PodHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (this *K8sConfig)InitClient() *kubernetes.Clientset {
	config ,err := clientcmd.BuildConfigFromFlags("","src/configs/K8sConfig")
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	return c
}

//初始化Informer
func(this *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact:=informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer:=fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	podInformer:=fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(this.PodHandler)
	fact.Start(wait.NeverStop)

	return fact
}