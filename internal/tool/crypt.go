package tool

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ulyssesorz/douyin/pkg/viper"
)


const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var (
	coder = base64.NewEncoding(base64Table)
	config = viper.Init("crypt")
	PublicKeyFilePath = config.Viper.GetString("rsa.douyin_message_encrypt_public_key")
	PrivateKeyFilePath = config.Viper.GetString("rsa.douyin_message_decrypt_private_key")
)

func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func RsaEncrypt(originData []byte, publicKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originData)
}

func Sha256Encrypt(data string, salt string) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(data + salt))
	cipherStr := sha256.Sum256(nil)
	return fmt.Sprintf("%x", cipherStr)
}

func Md5Encrypt(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	encryptedData := hex.EncodeToString(cipherStr)
	return encryptedData
}

func ReadKeyFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}