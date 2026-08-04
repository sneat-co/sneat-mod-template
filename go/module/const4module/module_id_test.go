package const4module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleID(t *testing.T) {
	assert.NotEmpty(t, ModuleID)
}
