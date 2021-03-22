package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiamond_ReturnsSingleLetter(t *testing.T) {
	require.Equal(t, "A", XShape('A'))
}
