package module

import (
	"github.com/sneat-co/sneat-go-core/module"
	"github.com/sneat-co/sneat-mod-module/go/module/api4module"
	"github.com/sneat-co/sneat-mod-module/go/module/const4module"
)

func Module() module.Module {
	return module.NewModule(const4module.ModuleID, module.RegisterRoutes(api4module.RegisterHttpRoutes))
}
