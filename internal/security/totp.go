package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TOTPHandler struct {
	key Key
	cfg *TOTPConfig
}

var (
	ErrInvalidOTP = errors.New("security: failed to verify OTP")
)

// Sign implements [TokenHandler].
func (handler *TOTPHandler) Sign(time time.Time) (string, error) {
	code, err := totp.GenerateCode(handler.key.String(), time)
	if err != nil {
		return "", fmt.Errorf("security: failed to sign OTP: %w", err)
	}

	return code, nil
}

// Verify implements [TokenHandler].
func (handler *TOTPHandler) Verify(token string) (*time.Time, error) {
	t := time.Now().UTC()
	valid, err := totp.ValidateCustom(token, handler.key.String(), t, handler.cfg.ValidateOptions)
	if err != nil {
		return nil, errors.Join(ErrInvalidOTP, err)
	}
	if !valid {
		return nil, ErrInvalidOTP
	}

	return &t, nil
}

var _ TokenHandler[time.Time] = (*TOTPHandler)(nil)

type TOTPConfig struct {
	GenerateOptions totp.GenerateOpts
	ValidateOptions totp.ValidateOpts
}

var DefaultTOTPConfig = TOTPConfig{
	GenerateOptions: totp.GenerateOpts{
		Period:     30,
		SecretSize: 20,
		Digits:     otp.DigitsSix,
		Algorithm:  otp.AlgorithmSHA1,
	},
	ValidateOptions: totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	},
}

type TOTPOption func(*TOTPConfig)

func NewTOTPHandler(key Key, opts ...TOTPOption) *TOTPHandler {
	cfg := DefaultTOTPConfig
	for _, fn := range opts {
		fn(&cfg)
	}

	return &TOTPHandler{key, &cfg}
}
