package security

import "encoding"

func Decode[T ~[]byte](s string) (T, error) {
	v := make(T, 0)
	if unmarshaler, ok := any(v).(encoding.TextUnmarshaler); ok {
		if err := unmarshaler.UnmarshalText([]byte(s)); err != nil {
			return nil, err
		}
	}

	return v, nil
}

func MustDecode[T ~[]byte](s string) T {
	v, err := Decode[T](s)
	if err != nil {
		panic(err)
	}

	return v
}
