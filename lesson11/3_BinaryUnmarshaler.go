type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}
