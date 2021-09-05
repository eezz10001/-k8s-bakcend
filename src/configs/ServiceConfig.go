package configs

import "k8s-bakcend/src/services"

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) CommonService() *services.CommonService {
	return services.NewCommonService()
}

func (this *ServiceConfig) DeploymentService() *services.DeploymentService {

	return services.NewDeploymentService()
}

func (this *ServiceConfig) PodtService() *services.PodService {

	return services.NewPodService()
}