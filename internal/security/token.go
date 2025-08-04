package security

// A TokenSigner signs identity tokens.
type TokenSigner[C any] interface {
	// Creates and signs a new token using the provided claims.
	Sign(claims C) (string, error)
}

// A TokenVerifier parses and verifies identity tokens.
type TokenVerifier[C any] interface {
	// Parses and verifies the provided token, returning the parsed claims.
	Verify(token string) (*C, error)
}

// A TokenHandler can both sign and verify identity tokens.
type TokenHandler[C any] interface {
	TokenSigner[C]
	TokenVerifier[C]
}
