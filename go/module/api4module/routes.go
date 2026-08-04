package api4module

import (
	"net/http"

	"github.com/sneat-co/sneat-go-core/extension"
)

func RegisterHttpRoutes(handle extension.HTTPHandleFunc) {
	handle("POST", "/api4module/about", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("api4module"))
	})
}
