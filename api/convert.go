package api

import "pnb/service"

func listSourceParamsFrom(req ListSourceReq) service.ListSourceParams {
	return service.ListSourceParams{
		CategoriesSet: len(req.Categories) > 0,
		Categories:    req.Categories,
		CountriesSet:  len(req.Countries) > 0,
		Countries:     req.Countries,
	}
}
