package bytes

import "encoding/base64"

type B64Encoder interface {
	Decode(src string) []byte
	EncodeToString(src []byte) string
	EncodeToUrl(src []byte) string
}

type B64 struct{}

func NewB64Encoder() *B64 {
	return &B64{}
}

var Std = NewB64Encoder()

func (b64 *B64) Decode(src string) (encoded []byte) {
	return []byte(b64.EncodeToString([]byte(src)))
}

func (b64 *B64) EncodeToString(src []byte) (encoded string) {
	return base64.StdEncoding.EncodeToString(src)
}

func (b64 *B64) EncodeToUrl(src []byte) (encoded string) {
	return base64.URLEncoding.EncodeToString(src)
}
