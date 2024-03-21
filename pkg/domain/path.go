package domain

import (
	"errors"
	"math"
)

// Path is a subset of Digraph, where for each point, or node, there can only exist at most one inbound connection, and
// at most one outbound connection.
//
// All points in the path must be connected to form one path, i.e., there should exist no partitions. Also, there
// should exist no cycles or branches in the path.
type Path[T comparable] struct {
	points     map[T]bool
	outboundOf map[T]T
	inboundOf  map[T]T
}

func NewPath[T comparable]() *Path[T] {
	return &Path[T]{
		points:     make(map[T]bool),
		outboundOf: make(map[T]T),
		inboundOf:  make(map[T]T),
	}
}

func (p *Path[T]) AddConnection(from, to T) error {
	if from == to {
		return errors.New(`invalid connection - "from" and "to" are the same`)
	}

	var nullValue T

	if outbound := p.outboundOf[from]; outbound != nullValue {
		return errors.New(`invalid connection - "from" already has an outbound connection`)
	}

	if p.inboundOf[to] != nullValue {
		return errors.New(`invalid connection - "to" already has an inbound connection`)
	}

	p.points[from] = true
	p.points[to] = true

	p.outboundOf[from] = to
	p.inboundOf[to] = from

	return nil
}

func (p *Path[T]) FindStart() (T, error) {
	var nullValue T

	for i := range p.points {
		if p.inboundOf[i] == nullValue {
			return i, nil
		}
	}

	return nullValue, errors.New("unable to find start of path - there's a loop")
}

func (p *Path[T]) FindEnd() (T, error) {
	var nullValue T

	for i := range p.points {
		if p.outboundOf[i] == nullValue {
			return i, nil
		}
	}

	return nullValue, errors.New("unable to find end of path - there's a loop")
}

func (p *Path[T]) GetNext(a T) T {
	return p.outboundOf[a]
}

func (p *Path[T]) Length() int {
	return int(
		math.Max(
			float64(len(p.inboundOf)),
			float64(len(p.outboundOf)),
		),
	)
}
