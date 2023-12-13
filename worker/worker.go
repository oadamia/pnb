package worker

import "pnb/service"

type Worker struct {
	service service.Service
}

func New(s service.Service) *Worker {
	return &Worker{
		service: s,
	}
}

func (w *Worker) Start() {

}
