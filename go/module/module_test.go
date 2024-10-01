package module

import (
	"github.com/sneat-co/sneat-go-core/module"
	"testing"
)

func TestModule(t *testing.T) {
	m := Module()
	module.AssertModule(t, m, module.Expected{
		ModuleID:      "{MODULE_ID}",
		HandlersCount: 1,
		DelayersCount: 0,
	})
}
