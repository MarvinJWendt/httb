package assets

import (
	"embed"
)

//go:embed openapi.yaml
var OpenAPISpec string

//go:embed swagger-ui
var Swagger embed.FS

//go:embed landing.html
var LandingPage string
