package encryption

import (
	"encoding/base64"
	"log"
)

const(
	///////////12345678901234567890123456789012
	AES_KEY = "YmVyYXBhcHVuaXR1c2VtdWFoYXJnYWhh"
)

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func AES256_Encrypt(pText string) (string, error) {

	reversedKey := reverse(AES_KEY)
	encoded := base64.StdEncoding.EncodeToString( []byte(pText) )
	encoded = AES_KEY +  encoded + reversedKey[0:6] + "=="

	return encoded, nil

	//text := []byte(pText)
	//key := []byte(pKey)
	//c, err := aes.NewCipher(key)
	//if err != nil {
	//	log.Println(err)
	//	return "",err
	//}

	//gcm, err := cipher.NewGCM(c)
	//if err != nil {
	//	log.Println(err)
	//	return "",err
	//}

	//nonce := make([]byte, gcm.NonceSize())
	//if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
	//	log.Println(err)
	//	return "",err
	//}

	//result := gcm.Seal(nonce, nonce, text, nil)
	//encoded := base64.StdEncoding.EncodeToString(result)
	//return encoded, nil

}

func AES256_Decrypt(encrypted string) (string, error) {

	encrypted = encrypted[32:len(encrypted)-8]
	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		log.Println("decode error:", err)
		return "", err
	}

	if string(cipherText) == "[object Object]" {
		return "", nil
	}
	return string(cipherText), nil

	//// Create new cipher block
	//key := []byte(AES_KEY)
	//block, err := aes.NewCipher(key)
	//if err != nil {
	//	return "", err
	//}
	//
	//// The IV needs to be unique, but not secure. Therefore it's common to
	//// include it at the beginning of the ciphertext.
	//if len(cipherText) < aes.BlockSize {
	//	panic("ciphertext too short")
	//}
	//iv := cipherText[:aes.BlockSize]
	//cipherText = cipherText[aes.BlockSize:]
	//
	//// CBC mode always works in whole blocks.
	//if len(cipherText)%aes.BlockSize != 0 {
	//	panic("ciphertext is not a multiple of the block size")
	//}
	//mode := cipher.NewCBCDecrypter(block, iv)
	//
	//// CryptBlocks can work in-place if the two arguments are the same.
	//mode.CryptBlocks(cipherText, cipherText)
	//fmt.Println("ciphertext::", string(cipherText))
	//
	//// Output: exampleplaintext
	//return string(cipherText), nil

	//key := []byte(pKey)
	//c, err := aes.NewCipher(key)
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}
	//
	//gcm, err := cipher.NewGCM(c)
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}

	//nonceSize := gcm.NonceSize()
	//if len(cipherText) < nonceSize {
	//	log.Println(err)
	//	return "", err
	//}

	//nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	//plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}

	//return string(plaintext), nil

}
