// Package utf7 implements modified UTF-7 encoding defined in RFC 3501 section 5.1.3
package utf7

import (
	"encoding/base64"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

const (
	min = 0x20 // Minimum self-representing UTF-7 value
	max = 0x7E // Maximum self-representing UTF-7 value

	repl = '\uFFFD' // Unicode replacement code point
)

var b64Enc = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+,")

type enc struct{}

func (e enc) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{
		Transformer: transform.Chain(encoding.UTF8Validator, &decoder{ascii: true}),
	}
}

func (e enc) NewEncoder() *encoding.Encoder {
	return &encoding.Encoder{
		Transformer: &encoder{},
	}
}

// Encoding is the modified UTF-7 encoding.
//
// Note, raw UTF-8 is accepted when decoding.
var Encoding encoding.Encoding = enc{}
