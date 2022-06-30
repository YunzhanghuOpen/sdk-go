package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// tripleDesEncrypt 3DES加密
func tripleDesEncrypt(originData, des3key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(des3key)
	if err != nil {
		return nil, err
	}
	originData = pkcs5Padding(originData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, des3key[:8])
	crypt := make([]byte, len(originData)) // nolint
	blockMode.CryptBlocks(crypt, originData)
	return crypt, nil
}

// tripleDesDecrypt 3DES解密
func tripleDesDecrypt(crypt, des3key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(des3key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, des3key[:8])
	originData := make([]byte, len(crypt)) // nolint
	blockMode.CryptBlocks(originData, crypt)
	originData = pkcs5UnPadding(originData)
	return originData, nil
}

// pkcs5Padding 填充
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs5UnPadding 取消填充
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后⼀一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// loadPublicKey 加载公钥
func loadPublicKey(data []byte) (pub *rsa.PublicKey, err error) {
	block, _ := pem.Decode(data)
	if block == nil {
		err = fmt.Errorf("decode public key fail")
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		err = errors.New("load public key fail")
	}
	return
}

// loadPrivateKey 加载私钥
func loadPrivateKey(privateKey []byte) (priv *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = fmt.Errorf("decode private key fail")
		return
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	priv, ok := key.(*rsa.PrivateKey)
	if !ok {
		err = errors.New("load private key fail")
	}
	return
}

func base64Encode(b []byte) ([]byte, error) {
	buf := bytes.Buffer{}
	en := base64.NewEncoder(base64.StdEncoding, &buf)
	_, err := en.Write(b)
	if err != nil {
		return nil, err
	}

	err = en.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
