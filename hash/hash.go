package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func MD5(original string) string {
	hash := md5.New()
	hash.Write([]byte(original))
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA224(original string) string {
	hash := sha256.Sum224([]byte(original))
	return hex.EncodeToString(hash[:])
}

func SHA256(original string) string {
	hash := sha256.Sum256([]byte(original))
	return hex.EncodeToString(hash[:])
}

func SHA384(original string) string {
	hash := sha512.Sum384([]byte(original))
	return hex.EncodeToString(hash[:])
}

func SHA512(original string) string {
	hash := sha512.Sum512([]byte(original))
	return hex.EncodeToString(hash[:])
}
