// duid means Distributed Unique Id.
// This package is a framework for generating unique id numbers such that the
// generation is fast and can be distributed across a network so that no two
// ids generated from different machines will be the same.
package duid

import (
	"time"
)

const (
	zeroLowShort int = ^0xFFFF
)

// IdGenerator produce distributed unique ids.
//
// NewId produces the next new id from this IdGenerator.
type IdGenerator interface {
	NewId() int
}

// Generator is a generator of distributed unique ids.
type generator struct {
	// uniqueKey is a machine-specific identifier to keep ids unique across a network.
	// The type here is set specifically to uint8 as is must be bounded my 255 in order
	// the generation algorithm to work, as there are only 8 bits reserved in the final
	// id for storing this uniqueKey.
	uniqueKey uint8
	// seqNum is the sequence number for ticking up sub-millisecond ids.
	seqNum uint8
}

// NewId for will generate a new id based on the current time, internal uniqueKey
// and sequence number currently set within the generator.
func (g *generator) NewId() int {
	var id = 0
	var now = time.Now().UnixNano()
	id |= int(now) & zeroLowShort
	id |= int(g.uniqueKey) << 8
	id |= int(g.seqNum)
	g.seqNum += 1
	return id
}

// NewGenerator creates a new Generator with the given uniqueKey for generating
// Ids across a distributed network.
func NewGenerator(uniqueKey uint) IdGenerator {
	return &generator{uniqueKey: uint8(uniqueKey), seqNum: 0}
}
