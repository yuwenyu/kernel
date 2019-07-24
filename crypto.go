package kernel

import (
"crypto/hmac"
"crypto/md5"
"crypto/sha1"
"encoding/hex"
)

type MCrypto interface {
	MCrypto(method string, data string) string
	MHmac(key string, data string) 		string
}

type mCrypto struct {}

var _ MCrypto = &mCrypto{}

func NewCrypto() *mCrypto {
	return &mCrypto{}
}

func (mc *mCrypto) MCrypto(method string, data string) string {
	switch method {
	case "md5":
		return mc.mMd5(data)
	case "sha":
		return mc.mSha(data)
	default:
		panic("Middleware Crypto Method Error")
	}
}

func (mc *mCrypto) mMd5(data string) string {
	cryptoMd5 := md5.New()
	cryptoMd5.Write([]byte(data))
	return hex.EncodeToString(cryptoMd5.Sum([]byte("")))
}

func (mc *mCrypto) mSha(data string) string {
	cryptoSha := sha1.New()
	cryptoSha.Write([]byte(data))
	return hex.EncodeToString(cryptoSha.Sum([]byte("")))
}

func (mc *mCrypto) MHmac(key string, data string) string {
	cryptoHmac := hmac.New(md5.New, []byte(key))
	cryptoHmac.Write([]byte(data))
	return hex.EncodeToString(cryptoHmac.Sum([]byte("")))
}


