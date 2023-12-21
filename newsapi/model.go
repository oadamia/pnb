package newsapi

import "pnb/service/store"

type SourcesResult struct {
	Status  string                   `json:"status"`
	Sources store.CreateSourceParams `json:"source"`
}
