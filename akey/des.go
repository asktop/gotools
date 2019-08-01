package akey

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
)

// =================== ECB模式 ======================
// DES加密, 使用ECB模式，注意key必须为8位长度
func DesEncryptECB(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	src = PKCS5Padding(src, blockSize) // PKCS5补位
	dst := make([]byte, len(src))      // 创建数组
	for i, count := 0, len(src)/blockSize; i < count; i++ { // 加密
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(dst[begin:end], src[begin:end])
	}
	return dst, nil
}

// DES解密, 使用ECB模式，注意key必须为8位长度
func DesDecryptECB(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	dst := make([]byte, len(src)) // 创建数组
	for i, count := 0, len(dst)/blockSize; i < count; i++ { // 解密
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(dst[begin:end], src[begin:end])
	}
	dst = PKCS5UnPadding(dst) // 去除PKCS5补位
	return dst, nil
}

// 3DES加密, 使用ECB模式，注意key必须为24位长度
func DesEncryptECBTriple(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	src = PKCS5Padding(src, blockSize) // PKCS5补位
	dst := make([]byte, len(src))      // 创建数组
	for i, count := 0, len(src)/blockSize; i < count; i++ { // 加密
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(dst[begin:end], src[begin:end])
	}
	return dst, nil
}

// 3DES解密, 使用ECB模式，注意key必须为24位长度
func DesDecryptECBTriple(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	dst := make([]byte, len(src)) // 创建数组
	for i, count := 0, len(dst)/blockSize; i < count; i++ { // 解密
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(dst[begin:end], src[begin:end])
	}
	dst = PKCS5UnPadding(dst) // 去除PKCS5补位
	return dst, nil
}

// =================== CBC模式 ======================
// DES加密, 使用CBC模式，注意key必须为8位长度，iv初始化向量为非必需参数（长度为8位）
func DesEncryptCBC(src []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ivValue := ([]byte)(nil) // 获取初始化向量
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = key
	}

	src = PKCS5Padding(src, block.BlockSize())          // PKCS5补位
	dst := make([]byte, len(src))                       // 创建数组
	blockMode := cipher.NewCBCEncrypter(block, ivValue) // 加密模式
	blockMode.CryptBlocks(dst, src)                     // 加密
	return dst, nil
}

// DES解密, 使用CBC模式，注意key必须为8位长度，iv初始化向量为非必需参数（长度为8位）
func DesDecryptCBC(src []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := des.NewCipher(key)
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
		ivValue = key
	}
	if len(src)%blockSize != 0 {
		return nil, errors.New("src is not a multiple of the block size")
	}

	dst := make([]byte, len(src))                       // 创建数组
	blockMode := cipher.NewCBCDecrypter(block, ivValue) // 加密模式
	blockMode.CryptBlocks(dst, src)                     // 解密
	dst = PKCS5UnPadding(dst)                           // 去除PKCS5补位
	return dst, nil
}

// 3DES加密, 使用CBC模式，注意key必须为24位长度，iv初始化向量为非必需参数（长度为8位）
func DesEncryptCBCTriple(src []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	ivValue := ([]byte)(nil) // 获取初始化向量
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = key[:blockSize]
	}

	src = PKCS5Padding(src, block.BlockSize())          // PKCS5补位
	dst := make([]byte, len(src))                       // 创建数组
	blockMode := cipher.NewCBCEncrypter(block, ivValue) // 加密模式
	blockMode.CryptBlocks(dst, src)                     // 加密
	return dst, nil
}

// 3DES解密, 使用CBC模式，注意key必须为24位长度，iv初始化向量为非必需参数（长度为8位）
func DesDecryptCBCTriple(src []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
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

	dst := make([]byte, len(src))                       // 创建数组
	blockMode := cipher.NewCBCDecrypter(block, ivValue) // 加密模式
	blockMode.CryptBlocks(dst, src)                     // 解密
	dst = PKCS5UnPadding(dst)                           // 去除PKCS5补位
	return dst, nil
}
