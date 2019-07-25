package util

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math/rand"
	"time"
)

type BaiduEncrypt struct {
	Key        string
	ClientId   string
	AesEncrypt *AesEncrypt
}

func NewBaiduEncrypt(base64Key, clientId string) (baiduEncrypt *BaiduEncrypt) {
	baiduEncrypt = &BaiduEncrypt{}
	key, _ := base64.StdEncoding.DecodeString(base64Key + "=")
	baiduEncrypt.Key = string(key)
	baiduEncrypt.ClientId = clientId
	baiduEncrypt.AesEncrypt = &AesEncrypt{}
	return
}

func (baiduEncrypt *BaiduEncrypt) Encode(src []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	buffer.Write(baiduEncrypt.randomString())
	err := binary.Write(buffer, binary.BigEndian, uint32(len(src)))
	if err != nil {
		return []byte{}, errors.New("baidu_encrypt/encode: binary write error")
	}
	buffer.Write(src)
	buffer.Write([]byte(baiduEncrypt.ClientId))
	encryptStr := buffer.Bytes()
	encryptedStr, err := baiduEncrypt.AesEncrypt.CBCEncrypt(encryptStr, []byte(baiduEncrypt.Key))
	if err != nil {
		return []byte{}, errors.New("baidu_encrypt/encode: encrypt error")
	}

	return encryptedStr, nil
}

func (baiduEncrypt *BaiduEncrypt) Decode(src []byte) (data string, err error) {
	decryptedBytes, err := baiduEncrypt.AesEncrypt.CBCDecrypt(src, []byte(baiduEncrypt.Key))
	if err != nil {
		return "", err
	}
	if len(decryptedBytes) < 16 {
		return "", errors.New("baidu_encrypt/decode: decrypted data length less then 16")
	}
	dataWithoutRandom := decryptedBytes[16:]
	buffer := bytes.NewBuffer(dataWithoutRandom)
	var dataLen uint32
	err = binary.Read(buffer, binary.BigEndian, &dataLen)
	if err != nil {
		return "", errors.New("baidu_encrypt/decode: read data length error")
	}
	dataAndClientId := dataWithoutRandom[4:]
	if dataLen != uint32(len(dataAndClientId) - len(baiduEncrypt.ClientId)) {
		return "", errors.New("baidu_encrypt/decode: data length error")
	}
	data = string(dataAndClientId[:dataLen])
	clientId := string(dataAndClientId[dataLen:])

	if clientId != baiduEncrypt.ClientId {
		return "", errors.New("baidu_encrypt/decode: client id not match")
	}

	return data, nil
}

func (baiduEncrypt *BaiduEncrypt) randomString() []byte {
	randomPol := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz")
	polLen := len(randomPol)
	var randomBytes []byte
	for i := 0; i < 16; i++ {
		pos := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(polLen)
		randomBytes = append(randomBytes, randomPol[pos])
	}

	return randomBytes
}
