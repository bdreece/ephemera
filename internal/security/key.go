package security

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"strings"
)

// A Key is a secret byte-array.
type Key []byte

// Returns the key as a byte-array.
func (key Key) Bytes() []byte { return []byte(key) }

// Encodes the key to a base-64 string
func (key Key) Encode() string {
	return base64.StdEncoding.EncodeToString(key)
}

// String implements [fmt.Stringer].
func (key Key) String() string { return key.Encode() }

// LogValue implements [slog.LogValuer].
func (Key) LogValue() slog.Value {
	return slog.StringValue(strings.Repeat("*", 64))
}

// UnmarshalText implements [encoding.TextUnmarshaler]
func (key *Key) UnmarshalText(text []byte) error {
	var err error
	if *key, err = base64.StdEncoding.DecodeString(string(text)); err != nil {
		return fmt.Errorf("security: failed to decode secret key: %w", err)
	}

	return nil
}
