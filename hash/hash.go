package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// Md5
func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

// 计算文件MD5值
func Md5File(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Sha1
func Sha1(value string) string {
	sha := sha1.New()
	sha.Write([]byte(value))
	b := sha.Sum(nil)
	return fmt.Sprintf("%x", b)
}

// Sha256
func Sha256(value string) string {
	sha := sha256.New()
	sha.Write([]byte(value))
	b := sha.Sum(nil)
	return fmt.Sprintf("%x", b)
}

// Sha512
func Sha512(value string) string {
	sha := sha512.New()
	sha.Write([]byte(value))
	b := sha.Sum(nil)

	return fmt.Sprintf("%x", b)
}
