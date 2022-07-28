package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

// Encoder 加密
type Encoder interface {
	Encode([]byte) ([]byte, error)
}

// Decoder 解密
type Decoder interface {
	Decode([]byte) ([]byte, error)
}

// Des3Encoding 加解密
type Des3Encoding struct {
	des3key []byte
}

var (
	_ Encoder = (*Des3Encoding)(nil)
	_ Decoder = (*Des3Encoding)(nil)
)

// NewDes3Encoding 新建加密解密对象
func NewDes3Encoding(des3key string) *Des3Encoding {
	return &Des3Encoding{des3key: []byte(des3key)}
}

// Encode 实现加密功能
func (h *Des3Encoding) Encode(b []byte) ([]byte, error) {
	b, err := tripleDesEncrypt(b, h.des3key)
	if err != nil {
		return nil, err
	}
	return base64Encode(b)
}

// Decode 实现解密接口
func (h *Des3Encoding) Decode(b []byte) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return nil, err
	}
	return tripleDesDecrypt(b, h.des3key)
}

var (
	_ Encoder = (*RsaEncoder)(nil)
	_ Decoder = (*RsaDecoder)(nil)
)

// RsaEncoder RSA 加密器
type RsaEncoder struct {
	pubKey *rsa.PublicKey
}

// NewRsaEncoder 新建 RSA 加密器
func NewRsaEncoder(publicKey []byte) (*RsaEncoder, error) {
	pubKey, err := loadPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return &RsaEncoder{pubKey: pubKey}, nil
}

// Encode 加密
func (r *RsaEncoder) Encode(data []byte) ([]byte, error) {
	b, err := rsa.EncryptPKCS1v15(rand.Reader, r.pubKey, data)
	if err != nil {
		return nil, err
	}
	return base64Encode(b)
}

// RsaDecoder RSA 解密器
type RsaDecoder struct {
	privKey *rsa.PrivateKey
}

// NewRsaDecoder 新建 RSA 解密器
func NewRsaDecoder(privKey []byte) (*RsaDecoder, error) {
	pri, err := loadPrivateKey(privKey)
	if err != nil {
		return nil, err
	}
	return &RsaDecoder{privKey: pri}, nil
}

// Decode 解密
func (r *RsaDecoder) Decode(b []byte) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, r.privKey, b)
}
