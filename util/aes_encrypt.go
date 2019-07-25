package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type AesEncrypt struct {
	key string
}

func (aesEncrypt *AesEncrypt) PKCS5Padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

func (aesEncrypt *AesEncrypt) PKCS5UnPadding(src []byte, blockSize int) ([]byte, error) {
	n := len(src)
	unPadNum := int(src[n-1])
	if unPadNum > blockSize {
		return []byte{}, errors.New("aes_encrypt/PKCS5UnPadding: un-padding number error")
	}
	return src[:n-unPadNum], nil
}

func (aesEncrypt *AesEncrypt) CBCEncrypt(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = aesEncrypt.PKCS5Padding(src, block.BlockSize())
	iv := key[:block.BlockSize()]
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

func (aesEncrypt *AesEncrypt) CBCDecrypt(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src)%block.BlockSize() != 0 {
		return nil, errors.New("aes_encrypt/CBCDecrypt: input not full blocks")
	}
	iv := key[:block.BlockSize()]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(src, src)
	dest, err := aesEncrypt.PKCS5UnPadding(src, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return dest, nil
}
