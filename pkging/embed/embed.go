package embed

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/markbates/pkger/here"
	"github.com/ulikunitz/xz/lzma"
)

func Decode(src []byte) ([]byte, error) {
	dst := make([]byte, base64.RawStdEncoding.DecodedLen(len(src)))
	_, err := base64.RawStdEncoding.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	r, err := lzma.NewReader2(bytes.NewReader(dst))
	if err != nil {
		return nil, err
	}

	bb := &bytes.Buffer{}
	if _, err := io.Copy(bb, r); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func Encode(b []byte) ([]byte, error) {
	bb := &bytes.Buffer{}
	gz, err := lzma.NewWriter2(bb)

	if err != nil {
		return nil, err
	}

	if _, err := gz.Write(b); err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	s := base64.RawStdEncoding.EncodeToString(bb.Bytes())
	return []byte(s), nil
}

type Data struct {
	Infos map[string]here.Info `json:"infos"`
	Files map[string]File      `json:"files"`
	Here  here.Info            `json:"here"`
}
