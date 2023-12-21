package api

import (
	"net/http"
	"net/http/httptest"
	"pnb/service"
	"pnb/service/store"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSources(t *testing.T) {
	e = echo.New()
	ctrl := gomock.NewController(t)
	querier := store.NewMockQuerier(ctrl)
	s = service.NewService(querier)

	t.Run("ListSources", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, listSourcePath, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, rec)

		querier.
			EXPECT().
			ListSource(ctx.Request().Context(), gomock.Any()).
			Return([]store.Source{}, nil)

		assert := assert.New(t)
		if assert.NoError(ListSource(ctx)) {
			assert.Equal(http.StatusOK, rec.Code)
		}

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
