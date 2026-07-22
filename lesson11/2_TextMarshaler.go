type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}
