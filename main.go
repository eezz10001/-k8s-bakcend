package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-bakcend/src/configs"
	"k8s-bakcend/src/controllers"
	"k8s-bakcend/src/middlewares"
)



func main() {
	goft.Ignite().Config(
		configs.NewK8sHandler(),    //1
		configs.NewK8sConfig(),     //2
		configs.NewK8sMaps(),       //3
		configs.NewServiceConfig(), //4
	).
		Mount("",
			controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
			controllers.NewWsCtl(),
	).Attach(middlewares.NewCrosMiddleware()).Launch()
}
