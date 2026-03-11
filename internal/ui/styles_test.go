package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylesInitialization(t *testing.T) {
	// styles.go defines variables and style constants,
	// a basic test ensures they are loaded/rendered without panic.

	// Check if colors are not empty internally
	assert.NotNil(t, ColorPrimary)
	assert.NotNil(t, ColorSuccess)
	assert.NotNil(t, ColorError)

	// Verify status constant rendering
	assert.Contains(t, StatusOkText, "ONLINE", "Success constant should contain ONLINE text")
	assert.Contains(t, StatusFailText, "OFFLINE", "Error constant should contain OFFLINE text")

	// This test is mainly for syntax coverage and preventing architecture regressions.
}
