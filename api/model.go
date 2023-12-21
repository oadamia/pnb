package api

type ListSourceReq struct {
	Categories []string `query:"category"`
	Countries  []string `query:"country"`
	Languages  []string `query:"language"`
}
