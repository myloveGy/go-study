package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	err := Open()
	assert.NoError(t, err)
}
