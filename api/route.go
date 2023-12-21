package api

const (
	healthCheckPath = "/healthcheck"

	listSourcePath   = "/sources"
	createSourcePath = "/sources"
	getSourcePath    = "/sources/:id"
	updateSourcePath = "/sources/:id"
	deleteSourcePath = "/sources/:id"
)

func registerRoutes() {
	group := e.Group("/pnb/v1")
	group.GET(healthCheckPath, HealthCheck)

	group.GET(listSourcePath, ListSource)
	group.POST(createSourcePath, CreateSource)
	group.GET(getSourcePath, GetSource)
	group.PUT(updateSourcePath, UpdateSource)
	group.DELETE(deleteSourcePath, DeleteSource)

}
