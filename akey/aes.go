package akey

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// =================== ECB模式 ======================
// AES加密, 使用ECB模式
func AesEncryptECB(plainText []byte, key []byte) (ciphertext []byte, err error) {
	cipher, err := aes.NewCipher(aesGenerateKey(key))
	if err != nil {
		return nil, err
	}
	length := (len(plainText) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, plainText)
	pad := byte(len(plain) - len(plainText))
	for i := len(plainText); i < len(plain); i++ {
		plain[i] = pad
	}
	ciphertext = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(plainText); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(ciphertext[bs:be], plain[bs:be])
	}
	return ciphertext, nil
}

// AES解密, 使用ECB模式
func AesDecryptECB(ciphertext []byte, key []byte) (plainText []byte, err error) {
	cipher, err := aes.NewCipher(aesGenerateKey(key))
	if err != nil {
		return nil, err
	}
	plainText = make([]byte, len(ciphertext))
	for bs, be := 0, cipher.BlockSize(); bs < len(ciphertext); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(plainText[bs:be], ciphertext[bs:be])
	}
	trim := 0
	if len(plainText) > 0 {
		trim = len(plainText) - int(plainText[len(plainText)-1])
	}
	return plainText[:trim], nil
}

func aesGenerateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CBC模式 ======================
// AES加密, 使用CBC模式，注意key必须为16/24/32位长度，iv初始化向量为非必需参数
func AesEncryptCBC(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()                    // 获取秘钥块的长度
	plainText = aesPKCS5Padding(plainText, blockSize) // 补全码
	ivValue := ([]byte)(nil)                          // 获取初始化向量
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = key[:blockSize]
	}
	blockMode := cipher.NewCBCEncrypter(block, ivValue) // 加密模式
	cipherText := make([]byte, len(plainText))          // 创建数组
	blockMode.CryptBlocks(cipherText, plainText)        // 加密
	return cipherText, nil
}

// AES解密, 使用CBC模式，注意key必须为16/24/32位长度，iv初始化向量为非必需参数
func AesDecryptCBC(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // 分组秘钥
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 获取秘钥块的长度
	if len(cipherText) < blockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil) // 获取初始化向量
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = key[:blockSize]
	}
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	blockModel := cipher.NewCBCDecrypter(block, ivValue)    // 加密模式
	plainText := make([]byte, len(cipherText))              // 创建数组
	blockModel.CryptBlocks(plainText, cipherText)           // 解密
	plainText, e := aesPKCS5UnPadding(plainText, blockSize) // 去除补全码
	if e != nil {
		return nil, e
	}
	return plainText, nil
}

// 补全码
func aesPKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// 去除补全码
func aesPKCS5UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if blockSize <= 0 {
		return nil, errors.New("invalid blocklen")
	}

	if length%blockSize != 0 || length == 0 {
		return nil, errors.New("invalid data len")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, errors.New("invalid padding")
	}

	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, errors.New("invalid padding")
		}
	}

	return src[:(length - unpadding)], nil
}

// =================== CFB模式 ======================
// AES加密, 使用CFB模式
func AesEncryptCFB(plainText []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)
	return ciphertext, nil
}

// AES解密, 使用CFB模式
func AesDecryptCFB(ciphertext []byte, key []byte) (plainText []byte, err error) {
	block, _ := aes.NewCipher(key)
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}
