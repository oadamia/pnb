package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pnb/mock"
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
	mockQuerier := store.NewMockQuerier(ctrl)
	s = service.NewService(mockQuerier)

	t.Run("ListSources", func(t *testing.T) {
		listQueryParams := "?category=general&category=sport&language=ge&country=ge"

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, listSourcePath+listQueryParams, nil)
		ctx := e.NewContext(req, rec)
		exp := mock.ListSource()

		mockQuerier.
			EXPECT().
			ListSource(ctx.Request().Context(), gomock.Any()).
			Return(exp, nil)

		assert := assert.New(t)
		if assert.NoError(ListSource(ctx)) {
			var actual []store.Source
			err := json.Unmarshal(rec.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(http.StatusOK, rec.Code)
			assert.Equal(exp, actual)
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
