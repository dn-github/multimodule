package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func GreetWorldTest(t *testing.T) {
	t.Run("Successful", func(t *testing.T) {
		require.Equal(t, "Hello, world.", GreetWorld())
	})
	t.Run("Failure", func(t *testing.T) {
		require.NotEqual(t, "Hello, world.", GreetWorld())
	})
}

func GreetUserTest(t *testing.T) {
	t.Run("Successful", func(t *testing.T) {
		require.Equal(t, "Hello, dev.", GreetUser("dev"))
	})
	t.Run("EmptyCase", func(t *testing.T) {
		require.Equal(t, "Hello, world.", GreetUser(""))
	})
}
