package structs

import "net/http"

type EndpointRequirements struct {
	Handler          http.HandlerFunc
	Required         []string
	CustomMiddleware []func(http.Handler) http.Handler
}
