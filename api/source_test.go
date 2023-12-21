package api

import (
	"pnb/service"
	"pnb/service/mock"
	"testing"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestSources(t *testing.T) {
	e = echo.New()
	ctrl := gomock.NewController(t)
	querier := mock.NewMockQuerier(ctrl)
	s = service.NewService(querier)

	t.Run("ListSources", func(t *testing.T) {

		// Your test code here
	})

	t.Run("CreateSource", func(t *testing.T) {
		// Your test code here
	})

	t.Run("GetSource", func(t *testing.T) {
		// Your test code here
	})

	t.Run("UpdateSource", func(t *testing.T) {
		// Your test code here
	})

	t.Run("DeleteSource", func(t *testing.T) {
		// Your test code here
	})
}
