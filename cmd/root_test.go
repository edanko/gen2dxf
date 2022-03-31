package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	t.Run("root cmd test", func(t *testing.T) {
		root := RootCmd()
		err := root.Execute()
		assert.NoError(t, err)
	})
}
