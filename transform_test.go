package vips_test

import (
	"testing"

	"github.com/davidbyttow/govips"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	buf, err := vips.NewTransform().
		LoadFile("fixtures/canyon.jpg").
		Scale(0.25).
		Apply()

	require.NoError(t, err)
	require.True(t, len(buf) > 0)

	image, err := vips.NewImageFromBuffer(buf)
	require.NoError(t, err)

	assert.Equal(t, 640, image.Width())
	assert.Equal(t, 400, image.Height())

	image.Close()
	vips.ShutdownThread()
	vips.Shutdown()

	vips.PrintObjectReport("Final")
}
