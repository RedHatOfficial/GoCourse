type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}
