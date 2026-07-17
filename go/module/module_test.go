package module

import (
	"github.com/sneat-co/sneat-go-core/coretypes"
	"github.com/sneat-co/sneat-go-core/extension"
	"testing"
)

func TestModule(t *testing.T) {
	m := Module()
	extension.AssertExtension(t, m, extension.Expected{
		ExtID:         coretypes.ExtID("{MODULE_ID}"),
		HandlersCount: 1,
		DelayersCount: 0,
	})
}
