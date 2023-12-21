package api

import (
	"pnb/service/store"
)

func listSourceParamsFrom(req ListSourceReq) store.ListSourceParams {
	return store.ListSourceParams{
		CategoriesSet: len(req.Categories) > 0,
		Categories:    req.Categories,
		CountriesSet:  len(req.Countries) > 0,
		Countries:     req.Countries,
		LanguagesSet:  len(req.Languages) > 0,
		Languages:     req.Languages,
	}
}
