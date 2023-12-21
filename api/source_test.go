package api

import (
	"bytes"
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
		exp := mock.ListSource()
		listQueryParams := "?category=general&category=sport&language=ge&country=ge"

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, listSourcePath+listQueryParams, nil)
		ctx := e.NewContext(req, rec)

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
		createParams := mock.CreateSourceParams()
		createBody, err := json.Marshal(createParams)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, createSourcePath, bytes.NewReader(createBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, rec)
		exp := mock.GetSource()

		mockQuerier.
			EXPECT().
			CreateSource(ctx.Request().Context(), gomock.Any()).
			Return(exp, nil)

		assert := assert.New(t)
		if assert.NoError(CreateSource(ctx)) {
			var actual store.Source
			err := json.Unmarshal(rec.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(http.StatusCreated, rec.Code)
			assert.Equal(exp, actual)
		}
	})

	t.Run("GetSource", func(t *testing.T) {
		exp := mock.GetSource()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, getSourcePath, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, rec)

		mockQuerier.
			EXPECT().
			GetSource(ctx.Request().Context(), gomock.Any()).
			Return(exp, nil)

		assert := assert.New(t)
		if assert.NoError(GetSource(ctx)) {
			var actual store.Source
			err := json.Unmarshal(rec.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(http.StatusOK, rec.Code)
			assert.Equal(exp, actual)
		}
	})

	t.Run("UpdateSource", func(t *testing.T) {
		exp := mock.GetSource()
		updateParams := mock.UpdateSourceParams()
		updateBody, err := json.Marshal(updateParams)
		if err != nil {
			t.Fatal(err)
		}

		updatedExp := exp
		updatedExp.Name = updateParams.Name
		updatedExp.Url = updateParams.Url
		updatedExp.Description = updateParams.Description
		updatedExp.Category = updateParams.Category
		updatedExp.Language = updateParams.Language
		updatedExp.Country = updateParams.Country

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, updateSourcePath, bytes.NewReader(updateBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, rec)

		mockQuerier.
			EXPECT().
			UpdateSource(ctx.Request().Context(), gomock.Any()).
			Return(updatedExp, nil)

		assert := assert.New(t)
		if assert.NoError(UpdateSource(ctx)) {
			var actual store.Source
			err := json.Unmarshal(rec.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(http.StatusOK, rec.Code)
			assert.Equal(updatedExp, actual)
		}
	})

	t.Run("DeleteSource", func(t *testing.T) {
		exp := mock.GetSource()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, deleteSourcePath, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, rec)

		mockQuerier.
			EXPECT().
			DeleteSource(ctx.Request().Context(), gomock.Any()).
			Return(exp, nil)

		assert := assert.New(t)
		if assert.NoError(DeleteSource(ctx)) {
			var actual store.Source
			err := json.Unmarshal(rec.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(http.StatusOK, rec.Code)
			assert.Equal(exp, actual)
		}
	})
}
