package decoder

type Decoder interface {
	Decode(input []byte)
}
