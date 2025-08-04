package identity

import (
	"log/slog"
	"math/rand"
	"strings"

	"github.com/bdreece/ephemera/internal/security"
)

type Password []byte

func (password Password) Hash() (security.Hash, security.Salt) {
	salt := security.NewSalt()
	hash := security.NewHash(password, salt)
	return hash, salt
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (password *Password) UnmarshalText(text []byte) error {
	*password = text
	return nil
}

// LogValue implements [slog.LogValuer].
func (password Password) LogValue() slog.Value {
	length := 8 + rand.Intn(4)
	masked := strings.Repeat("*", length)
	return slog.StringValue(masked)
}
