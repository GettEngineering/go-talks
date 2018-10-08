package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetting(t *testing.T) {
	settings := SomeSettings{
		Extra:            1,
		Type:                   "text",
		UnavailablePetaTextEnabled: false,
	}

	err := InsertUpdateSettings(settings)
	assert.Equal(t, nil, err)

}