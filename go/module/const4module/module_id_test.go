package const4module

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModuleID(t *testing.T) {
	assert.NotEmpty(t, ModuleID)
}
