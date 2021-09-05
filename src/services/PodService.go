package services

//@Service
type PodService struct {
	PodMap *PodMapStruct `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}
func(this *PodService) ListByNs(ns string ) interface{}{
	return this.PodMap.ListByNs(ns)
}