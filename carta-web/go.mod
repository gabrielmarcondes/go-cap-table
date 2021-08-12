module example.com/carta-web

go 1.16

require (
	example.com/captable v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.3
	github.com/google/uuid v1.3.0
)

replace example.com/captable => ../captable
