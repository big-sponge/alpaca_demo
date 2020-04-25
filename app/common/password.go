package common

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	cRand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh"
	"math/rand"
	"strings"
	"time"
)

/**
 * 密码生成
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func PasswordHash(pwd string) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(hashBytes), err
}

/**
 * 密码验证
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func PasswordVerify(hashedPwd string, plainPwd string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if check != nil {
		return false
	}
	return true
}

/**
 *  token
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func InitToken(str string) (result string) {
	base := InitRandomWord(32) + str
	h := md5.New()
	h.Write([]byte(base))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * md5
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func InitMd5(str string) (result string) {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/*
 * 生成随机字符串
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func InitRandomWord(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	cLen := len(chars)
	res := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		rfi := rand.Intn(cLen - 1)
		res += fmt.Sprintf("%c", chars[rfi])
	}
	return res
}

/*
 * 生成随机密码
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
func DesEncrypt(origData, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil

	}
	origData = PKCS5Padding(origData, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func InitRandomPassword(length int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cLen := len(chars)
	res := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		rfi := rand.Intn(cLen - 1)
		res += fmt.Sprintf("%c", chars[rfi])
	}
	return res
}

/*
 * 生成SSH_KEY
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}
func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}
func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(cRand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}
func MakeSSHKeyPair() map[string]string {
	var keyPairData = make(map[string]string)
	pkey, pubkey, err := GenerateKey(2048)
	if err != nil {
		return keyPairData
	}

	pub, err := EncodeSSHKey(pubkey)
	if err != nil {
		return keyPairData
	}
	keyPairData["PublicKey"] = strings.Replace(string(pub), "\n", "", -1) + " Generated-by-Nugget"
	keyPairData["PrivateKey"] = string(EncodePrivateKey(pkey))
	return keyPairData
}
func GenerateSSHKey(key string) []byte {
	keyPairData := MakeSSHKeyPair()
	mjson, _ := json.Marshal(keyPairData)
	cypt := DesEncrypt([]byte(string(mjson)), []byte(key))
	return cypt
}
