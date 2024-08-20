package api4module

import (
	"github.com/sneat-co/sneat-go-core/module"
	"net/http"
)

func RegisterHttpRoutes(handle module.HTTPHandleFunc) {
	handle("POST", "/api4module/about", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("api4module"))
	})
}
