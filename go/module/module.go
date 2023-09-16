package module

import (
	"github.com/sneat-co/sneat-go-core/modules"
	"github.com/sneat-co/sneat-mod-module/go/module/api4module"
	"github.com/sneat-co/sneat-mod-module/go/module/const4module"
)

func Module() modules.Module {
	return modules.NewModule(const4module.ModuleID, api4module.RegisterHttpRoutes)
}
