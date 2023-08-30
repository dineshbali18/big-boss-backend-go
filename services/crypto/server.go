package crypto

// import (
// 	"big-boss-7/config"
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/x509"
// 	"encoding/base64"
// 	"errors"
// )

// // ServerEncrypt encrypts the given text using server-side encryption. This is used for text that will be encrypted and decrypted on the server.
// func ServerEncrypt(rawText []byte) (string, error) {

// 	// Parse the DER-encoded RSA public key
// 	publicKey, err := x509.ParsePKIXPublicKey(config.CryptoConfig.RSAPublicKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Assert the public key type as *rsa.PublicKey
// 	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
// 	if !ok {
// 		return "", errors.New("invalid RSA public key format")
// 	}

// 	// Encrypt the plaintext using RSA
// 	encryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, rawText)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Encode the encrypted text as a base64 string
// 	encodedText := base64.StdEncoding.EncodeToString(encryptedText)

// 	return encodedText, nil

// }

// // ServerDecrypt decrypts the given text that was encrypted using server-side encryption. This is used for text that will be encrypted and decrypted on the server.
// func ServerDecrypt(encryptedText string) (string, error) {

// 	// Decode the base64-encoded encrypted text
// 	decodedText, err := base64.StdEncoding.DecodeString(encryptedText)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Parse the DER-encoded RSA private key
// 	privateKey, err := x509.ParsePKCS1PrivateKey(config.CryptoConfig.RSAPrivateKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Decrypt the encrypted text using RSA
// 	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodedText)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(decryptedText), nil
// }
