package crypto

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Signer 签名
type Signer interface {
	Sign(mess, timestamp, data string) (string, error)
}

// SignVerifier 验签
type SignVerifier interface {
	Verify(mess, timestamp, data, sign string) (bool, error)
}

// NewHmacSigner 新建 HMAC 签名
func NewHmacSigner(appKey string) (Signer, error) {
	return &hmacSigner{
		appKey: appKey,
	}, nil
}

// NewHmacSignVerifier 新建 HMAC 验签
func NewHmacSignVerifier(appKey string) (SignVerifier, error) {
	return &hmacSigner{
		appKey: appKey,
	}, nil
}

type hmacSigner struct {
	appKey string
}

// Sign 实现签名方法
func (h *hmacSigner) Sign(mess, timestamp, data string) (string, error) {
	plaintext := fmt.Sprintf("data=%s&mess=%s&timestamp=%s&key=%s", data, mess, timestamp, h.appKey)
	sign := hmac.New(sha256.New, []byte(h.appKey))
	_, err := sign.Write([]byte(plaintext))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sign.Sum(nil)), nil
}

// Verify 实现验证方法
func (r *hmacSigner) Verify(mess, timestamp, data string, sign string) (bool, error) {
	dest, err := r.Sign(mess, timestamp, data)
	if err != nil {
		return false, err
	}
	return dest == sign, nil
}

type rsaSigner struct {
	priv   *rsa.PrivateKey
	appKey string
}

// NewRsaSigner 新建 RSA 签名
func NewRsaSigner(privateKey, appkey string) (Signer, error) {
	priv, err := loadPrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	return &rsaSigner{
		appKey: appkey,
		priv:   priv,
	}, nil
}

func (r *rsaSigner) Sign(mess, timestamp, data string) (string, error) {
	b := []byte(fmt.Sprintf("data=%s&mess=%s&timestamp=%s&key=%s", data, mess, timestamp, string(r.appKey)))
	hash := sha256.New()
	_, err := hash.Write(b)
	if err != nil {
		return "", err
	}
	rsaSign, err := rsa.SignPKCS1v15(rand.Reader, r.priv, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(rsaSign), nil

}

// NewRsaSignVerifier 新建 RSA 签名验证
func NewRsaSignVerifier(publicKey, appKey string) (SignVerifier, error) {
	pubKey, err := loadPublicKey([]byte(publicKey))
	if err != nil {
		return nil, err
	}

	return &rsaSignVerifier{
		pubKey: pubKey,
		appKey: appKey,
	}, nil
}

type rsaSignVerifier struct {
	pubKey *rsa.PublicKey
	appKey string
}

// Verify 实现验证方法
func (r *rsaSignVerifier) Verify(mess, timestamp, data string, sign string) (bool, error) {
	b, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	bt := []byte(fmt.Sprintf("data=%s&mess=%s&timestamp=%s&key=%s", data, mess, timestamp, string(r.appKey)))
	hash := sha256.New()
	_, err = hash.Write(bt)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(r.pubKey, crypto.SHA256, hash.Sum(nil), b)
	if err != nil && err != rsa.ErrVerification {
		return false, err
	}
	return err == nil, nil
}
