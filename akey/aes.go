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
// AES加密, 使用ECB模式，注意key必须为16/24/32位长度
func AesEncryptECB(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src = PKCS5Padding(src, block.BlockSize()) // PKCS5补位
	dst := make([]byte, len(src))              // 创建数组
	blockMode := NewECBEncrypter(block)        // 加密模式
	blockMode.CryptBlocks(dst, src)            // 加密
	return dst, nil
}

// AES解密, 使用ECB模式，注意key必须为16/24/32位长度
func AesDecryptECB(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(src))       // 创建数组
	blockMode := NewECBDecrypter(block) // 加密模式
	blockMode.CryptBlocks(dst, src)     // 解密
	dst = PKCS5UnPadding(dst)           // 去除PKCS5补位
	return dst, nil
}

// ------------------- ECB处理工具 ----------------------
type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// =================== CBC模式 ======================
// AES加密, 使用CBC模式，注意key必须为16/24/32位长度，iv初始化向量为非必需参数（长度为16位）
func AesEncryptCBC(src []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 获取秘钥块的长度
	ivValue := ([]byte)(nil)       // 获取初始化向量
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = key[:blockSize]
	}

	src = PKCS5Padding(src, blockSize)                  // PKCS5补位
	dst := make([]byte, len(src))                       // 创建数组
	blockMode := cipher.NewCBCEncrypter(block, ivValue) // 加密模式
	blockMode.CryptBlocks(dst, src)                     // 加密
	return dst, nil
}

// AES解密, 使用CBC模式，注意key必须为16/24/32位长度，iv初始化向量为非必需参数（长度为16位）
func AesDecryptCBC(src []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // 分组秘钥
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 获取秘钥块的长度
	if len(src) < blockSize {
		return nil, errors.New("src is too short, less than block size")
	}
	ivValue := ([]byte)(nil) // 获取初始化向量
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = key[:blockSize]
	}
	if len(src)%blockSize != 0 {
		return nil, errors.New("src is not a multiple of the block size")
	}

	dst := make([]byte, len(src))                        // 创建数组
	blockModel := cipher.NewCBCDecrypter(block, ivValue) // 加密模式
	blockModel.CryptBlocks(dst, src)                     // 解密
	dst = PKCS5UnPadding(dst)                            // 去除PKCS5补位
	return dst, nil
}

// =================== CFB模式 ======================
// AES加密, 使用CFB模式
func AesEncryptCFB(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, aes.BlockSize+len(src))
	iv := dst[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(dst[aes.BlockSize:], src)
	return dst, nil
}

// AES解密, 使用CFB模式
func AesDecryptCFB(src []byte, key []byte) (dst []byte, err error) {
	block, _ := aes.NewCipher(key)
	if len(src) < aes.BlockSize {
		return nil, errors.New("src is too short, less than block size")
	}
	iv := src[:aes.BlockSize]
	src = src[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(src, src)
	return src, nil
}

// =================== PKCS5 ======================
// PKCS5补位
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// 去除PKCS5补位
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	padtext := int(src[length-1])
	return src[:(length - padtext)]
}
