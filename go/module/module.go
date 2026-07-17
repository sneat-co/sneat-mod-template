package module

import (
	"github.com/sneat-co/sneat-go-core/coretypes"
	"github.com/sneat-co/sneat-go-core/extension"
	"github.com/sneat-co/sneat-mod-module/go/module/api4module"
	"github.com/sneat-co/sneat-mod-module/go/module/const4module"
)

func Module() extension.Config {
	return extension.NewExtension(coretypes.ExtID(const4module.ModuleID), extension.RegisterRoutes(api4module.RegisterHttpRoutes))
}
