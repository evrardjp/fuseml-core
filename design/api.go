package design

import (
	"fmt"

	. "goa.design/goa/v3/dsl"
)

func buildServer(host string) {
	address := "0.0.0.0"
	httpPort := "8000"
	grpcPort := "8080"

	if host == "dev" {
		address = "localhost"
	}
	Description(fmt.Sprintf("Run the service listening on %s.", address))
	URI(fmt.Sprintf("http://%s:%s", address, httpPort))
	URI(fmt.Sprintf("grpc://%s:%s", address, grpcPort))
}

var _ = API("fuseml", func() {
	Title("FuseML core API")
	Description("Provides an API for the core operations of FuseML")

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("fuseml-core", func() {
		Description("fuseml-core services")

		// List the services hosted by this server.
		Services("runnable", "openapi")

		// List the Hosts and their transport URLs.
		Host("prod", func() { buildServer("prod") })
		Host("dev", func() { buildServer("dev") })
	})
})
