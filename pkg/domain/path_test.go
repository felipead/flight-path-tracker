package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath_AddConnection_FailsIfPointsAreTheSame(t *testing.T) {
	p := NewPath[string]()
	err := p.AddConnection("foo", "foo")
	assert.EqualError(t, err, `invalid connection - "from" and "to" are the same`)
}

func TestPath_AddConnection_FailsIfFromAlreadyHasAnOutboundConnection(t *testing.T) {
	p := NewPath[string]()

	err := p.AddConnection("foo", "bar")
	assert.NoError(t, err)

	err = p.AddConnection("foo", "baz")

	assert.EqualError(t, err, `invalid connection - "from" already has an outbound connection`)
}

func TestPath_AddConnection_FailsIfToAlreadyHasAnInboundConnection(t *testing.T) {
	p := NewPath[string]()

	err := p.AddConnection("foo", "bar")
	assert.NoError(t, err)

	err = p.AddConnection("baz", "bar")

	assert.EqualError(t, err, `invalid connection - "to" already has an inbound connection`)
}

func TestPath_AddConnection_FailsIfAddingTheSameConnectionTwice(t *testing.T) {
	p := NewPath[string]()

	err := p.AddConnection("foo", "bar")
	assert.NoError(t, err)

	err = p.AddConnection("foo", "bar")

	assert.EqualError(t, err, `invalid connection - "from" already has an outbound connection`)
}

func TestPath_FindStart(t *testing.T) {
	p := NewPath[string]()

	assert.NoError(t, p.AddConnection("c", "d"))
	assert.NoError(t, p.AddConnection("a", "b"))
	assert.NoError(t, p.AddConnection("b", "c"))

	start, err := p.FindStart()
	assert.NoError(t, err)
	assert.Equal(t, start, "a")
}

func TestPath_FindStart_FailsIfTheresALoop(t *testing.T) {
	p := NewPath[string]()

	assert.NoError(t, p.AddConnection("a", "b"))
	assert.NoError(t, p.AddConnection("b", "c"))
	assert.NoError(t, p.AddConnection("c", "d"))
	assert.NoError(t, p.AddConnection("d", "a"))

	_, err := p.FindStart()
	assert.EqualError(t, err, "unable to find start of path - there's a loop")
}

func TestPath_FindEnd(t *testing.T) {
	p := NewPath[string]()

	assert.NoError(t, p.AddConnection("b", "c"))
	assert.NoError(t, p.AddConnection("d", "e"))
	assert.NoError(t, p.AddConnection("a", "b"))
	assert.NoError(t, p.AddConnection("c", "d"))

	end, err := p.FindEnd()
	assert.NoError(t, err)
	assert.Equal(t, end, "e")
}

func TestPath_FindEnd_FailIfTheresALoop(t *testing.T) {
	p := NewPath[string]()

	assert.NoError(t, p.AddConnection("a", "b"))
	assert.NoError(t, p.AddConnection("b", "c"))
	assert.NoError(t, p.AddConnection("c", "d"))
	assert.NoError(t, p.AddConnection("d", "e"))
	assert.NoError(t, p.AddConnection("e", "a"))

	_, err := p.FindEnd()
	assert.EqualError(t, err, "unable to find end of path - there's a loop")
}

func TestPath_NavigateThePath(t *testing.T) {
	p := NewPath[string]()

	assert.NoError(t, p.AddConnection("d", "e"))
	assert.NoError(t, p.AddConnection("b", "c"))
	assert.NoError(t, p.AddConnection("c", "d"))
	assert.NoError(t, p.AddConnection("e", "f"))
	assert.NoError(t, p.AddConnection("a", "b"))

	start, err := p.FindStart()
	assert.NoError(t, err)
	assert.Equal(t, start, "a")

	assert.Equal(t, p.GetNext("a"), "b")
	assert.Equal(t, p.GetNext("b"), "c")
	assert.Equal(t, p.GetNext("c"), "d")
	assert.Equal(t, p.GetNext("d"), "e")
	assert.Equal(t, p.GetNext("e"), "f")
	assert.Equal(t, p.GetNext("f"), "")

	end, err := p.FindEnd()
	assert.NoError(t, err)
	assert.Equal(t, end, "f")
}
