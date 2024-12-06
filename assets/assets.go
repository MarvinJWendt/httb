package assets

import (
	"embed"
	_ "embed"
)

//go:embed openapi.yaml
var OpenAPISpec string

//go:embed docs.html
var Docs string

//go:embed swagger-ui
var Swagger embed.FS
