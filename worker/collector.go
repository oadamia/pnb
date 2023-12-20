package worker

import (
	"log/slog"
	"pnb/service"

	"github.com/robfig/cron/v3"
)

var counter int

func collectJob(s *service.Service) cron.FuncJob {
	return func() {
		counter++
		slog.Info("Collecting...")
		if counter == 3 {
			panic("test")
		}
	}
}
