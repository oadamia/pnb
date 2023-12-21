package worker

import (
	"log/slog"
	"os"
	"pnb/service"

	"github.com/robfig/cron/v3"
)

type Worker struct {
	service *service.Service
	c       *cron.Cron
}

func Init(s *service.Service) error {
	logWrapper := cron.VerbosePrintfLogger(slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo))
	recoveryWrapper := cron.Recover(logWrapper)

	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(logWrapper),
		cron.WithChain(recoveryWrapper))

	w := &Worker{
		service: s,
		c:       c,
	}

	err := w.start()
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) start() error {
	_, err := w.c.AddFunc("* * * * * *", w.collectorJob)
	if err != nil {
		return err
	}

	w.c.Start()
	return nil
}

func (w *Worker) collectorJob() {
	slog.Info("Collecting...")
}
