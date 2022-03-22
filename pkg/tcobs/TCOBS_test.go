package src

import (
	"testing"

	"github.com/tj/assert"
)

type testTable []struct {
	raw []byte
	exp []byte
}

var testData = testTable{
	{[]byte{}, []byte{}},
	{[]byte{0}, []byte{0x20}},
	{[]byte{0, 0}, []byte{0x40}},
	{[]byte{0, 0xAA}, []byte{0x20, 0xAA, 0xA1}},
	{[]byte{0xAA, 0}, []byte{0xAA, 0x21}},
	{[]byte{0xAA, 0xAA}, []byte{0xAA, 0xAA, 0xA2}},
	{[]byte{0xAA, 0xBB}, []byte{0xAA, 0xBB, 0xA2}},
	{[]byte{0, 0, 0}, []byte{0x60}},
	{[]byte{0, 0, 0, 0}, []byte{0x40, 0x40}},
	{[]byte{0xFF}, []byte{0xFF, 0xA1}},
	{[]byte{0xFF, 0xFF}, []byte{0xC0}},

	// no code {[]byte{0xFF, 0xFF, 0xFF}, []byte{0xE0}},
	// no code {[]byte{0xFF, 0xFF, 0xFF, 0xFF}, []byte{0x80}},
	// no code {[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, []byte{0x80, 0xFF, 0xA1}},
	// {[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, []byte{0x80, 0xC0}},
	// {[]byte{0xAA}, []byte{0xAA, 0xA1}},
	// {[]byte{0xAA, 0xAA}, []byte{0xAA, 0xAA, 0xA2}},
	// {[]byte{0xAA, 0xAA, 0xAA}, []byte{0xAA, 0x09}},
	// {[]byte{0xAA, 0xAA, 0xAA, 0xAA}, []byte{0xAA, 0x11}},
	// {[]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}, []byte{0xAA, 0x19}},
	// {[]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA}, []byte{0xAA, 0x01}},
	// {[]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA}, []byte{0xAA, 0x01, 0xAA, 0xA1}},
}

func TestTCOBSEncode(t *testing.T) {
	for _, k := range testData {
		enc := make([]byte, 100)
		n := TCOBSEncodeC(enc, k.raw)
		enc = enc[:n]
		assert.Equal(t, k.exp, enc)
	}
}
