package simplegoflake

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// We'll use January 01, 2011, as the epoch, to match Rubyflake.
var EPOCH = time.Unix(1293858000, 0)

// We'll use 41 bits for the timestamp.
const TIMESTAMP_BITS = 41

// 41 bits for the timestamp leaves 23 bits for the random data.
const RANDOM_BITS = 23
const RANDOM_BYTES = 3

// We'll shift the timestamp left enough to make room for the random data.
const TIMESTAMP_SHIFT = RANDOM_BITS

// The maximum value we can store in 23 bits of random data.
const RANDOM_MAX = 0x7fffff

// A Generator will generate IDs based on a particular epoch.
type Generator struct {
	Epoch time.Time
}

// Generate() will return a more or less ordered, more or less unique ID.
//
// It will return an error if it cannot read random data using crypto/rand.
func (g Generator) Generate() (flake uint64, err error) {
	milliseconds := uint64(time.Since(g.Epoch).Nanoseconds() / 1000000)

	buf := make([]byte, RANDOM_BYTES)
	_, err = rand.Read(buf)
	if err != nil {
		return
	}

	random, _ := binary.Uvarint(buf)

	flake = (milliseconds << TIMESTAMP_SHIFT) + (random & RANDOM_MAX)
	return
}

var defaultGenerator = Generator{
	Epoch: EPOCH,
}

// Generate() will return an ID using a default Generator based on EPOCH.
func Generate() (uint64, error) {
	return defaultGenerator.Generate()
}
