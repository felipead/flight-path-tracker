package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAirportCode_IsValid(t *testing.T) {
	assert.True(t, AirportCode("SFO").IsValid())
	assert.True(t, AirportCode("ORD").IsValid())
	assert.True(t, AirportCode("MIA").IsValid())
}

func TestAirportCode_IsNotValid(t *testing.T) {
	assert.False(t, AirportCode("").IsValid())
	assert.False(t, AirportCode("  ").IsValid())
	assert.False(t, AirportCode("  ORD").IsValid())
	assert.False(t, AirportCode("O5D").IsValid())
	assert.False(t, AirportCode("FOOO").IsValid())
}
