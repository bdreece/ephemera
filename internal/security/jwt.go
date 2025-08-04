package security

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// A JWTHandler is a [TokenHandler] implemented using JSON web tokens.
type JWTHandler[C jwt.Claims] struct {
	keyfn      jwt.Keyfunc
	tokenOpts  []jwt.TokenOption
	parserOpts []jwt.ParserOption
}

// Sign implements [TokenSigner].
func (handler *JWTHandler[C]) Sign(claims C) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		claims,
		handler.tokenOpts...,
	)

	secret, _ := handler.keyfn(token)
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("security: failed to sign token: %w", err)
	}

	return signed, nil
}

// Verify implements [TokenVerifier].
func (handler *JWTHandler[C]) Verify(token string) (*C, error) {
	tok, err := jwt.ParseWithClaims(
		token,
		any(new(C)).(jwt.Claims),
		handler.keyfn,
		handler.parserOpts...,
	)
	if err != nil {
		return nil, fmt.Errorf("security: failed to verify token: %w", err)
	}

	claims, ok := any(tok.Claims).(*C)
	if !ok {
		return nil, fmt.Errorf("security: invalid claims type")
	}

	return claims, nil
}

// A JWTConfig provides the signing and parsing options used by [JWTHandler].
type JWTConfig struct {
	TokenOptions  []jwt.TokenOption
	ParserOptions []jwt.ParserOption
}

// A JWTOption is used to modify the [JWTConfig].
type JWTOption func(cfg *JWTConfig)

// Creates a new [JWTHandler] with the provided secret.
func NewJWTHandler[C jwt.Claims](secret Key, opts ...JWTOption) *JWTHandler[C] {
	cfg := JWTConfig{
		TokenOptions:  []jwt.TokenOption{},
		ParserOptions: []jwt.ParserOption{},
	}

	for _, fn := range opts {
		fn(&cfg)
	}

	keyfn := func(*jwt.Token) (any, error) {
		return secret, nil
	}

	return &JWTHandler[C]{keyfn, cfg.TokenOptions, cfg.ParserOptions}
}

// Configures the signing behavior of the [JWTHandler].
func SignJWTWith(opts ...jwt.TokenOption) JWTOption {
	return func(cfg *JWTConfig) {
		cfg.TokenOptions = append(cfg.TokenOptions, opts...)
	}
}

// Configures the verifying behavior of the [JWTHandler].
func VerifyJWTWith(opts ...jwt.ParserOption) JWTOption {
	return func(cfg *JWTConfig) {
		cfg.ParserOptions = append(cfg.ParserOptions, opts...)
	}
}
