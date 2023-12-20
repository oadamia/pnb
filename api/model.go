package api

type ListSourceReq struct {
	Categories []string `query:"category"`
	Countries  []string `queryjson:"country"`
}
