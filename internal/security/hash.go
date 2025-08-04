package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"golang.org/x/crypto/argon2"
)

// A Hash is a [Key] derived from a password and [Salt].
//
// See [crypto/argon2] for more details.
//
// [crypto/argon2]: https://golang.org/x/crypto/argon2
type Hash []byte

var HashLength int

// Returns the key as a byte-array.
func (key Hash) Bytes() []byte { return []byte(key) }

// String implements [fmt.Stringer].
func (key Hash) String() string {
	return base64.StdEncoding.EncodeToString(key)
}

// LogValue implements [slog.LogValuer].
func (Hash) LogValue() slog.Value {
	return slog.StringValue(strings.Repeat("*", 64))
}

// UnmarshalText implements [encoding.TextUnmarshaler]
func (key *Hash) UnmarshalText(text []byte) error {
	if _, err := base64.StdEncoding.Decode(*key, text); err != nil {
		return fmt.Errorf("security: failed to decode secret key: %w", err)
	}

	return nil
}

func (hash Hash) Compare(value []byte, salt Salt) bool {
	return slices.Equal(hash, NewHash(value, salt))
}

func NewHash(value []byte, salt Salt) Hash {
	const (
		time    uint32 = 1
		memory  uint32 = 64 * 1024
		threads uint8  = 4
		keyLen  uint32 = 32
	)

	key := argon2.IDKey(value, salt[:], time, memory, threads, keyLen)
	return Hash(key)
}

// A Salt is a byte-array generated with a cryptographic RNG.
//
// See [crypto/argon2] for more details.
//
// [crypto/argon2]: https://golang.org/x/crypto/argon2
type Salt []byte

var SaltLength int

func (salt Salt) String() string {
	return base64.StdEncoding.EncodeToString(salt)
}

func (salt *Salt) UnmarshalText(text []byte) error {
	if _, err := base64.StdEncoding.Decode(*salt, text); err != nil {
		return fmt.Errorf("security: failed to unmarshal password salt: %w", err)
	}

	return nil
}

func NewSalt() Salt {
	salt := make([]byte, SaltLength)
	_, _ = rand.Read(salt)

	return salt
}
