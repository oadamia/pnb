package worker

import (
	"log/slog"
	"os"
	"pnb/newsapi"
	"pnb/service"

	"github.com/robfig/cron/v3"
)

type Worker struct {
	service       *service.Service
	cronClient    *cron.Cron
	newsapiClient *newsapi.Client
}

func Init(s *service.Service) error {
	logWrapper := cron.VerbosePrintfLogger(slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo))
	recoveryWrapper := cron.Recover(logWrapper)

	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(logWrapper),
		cron.WithChain(recoveryWrapper))

	w := &Worker{
		service:    s,
		cronClient: c,
	}

	err := w.start()
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) start() error {
	_, err := w.cronClient.AddFunc("0 0 * * * *", w.collectorJob)
	if err != nil {
		return err
	}

	_, err = w.cronClient.AddFunc("30 * * * * *", w.refreshSources)
	if err != nil {
		return err
	}

	w.cronClient.Start()
	return nil
}

func (w *Worker) collectorJob() {
	slog.Info("Collecting...")
}
