package hash

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

// simple hash an input of string
func HashSHA512(input string) (string, error) {
	hash := sha512.New()

	// Write the input string to the hash
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", err
	}

	// Calculate the hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}

// check hash is same or not
func CheckSHA512(input, hashToCheck string) bool {
	// Hash the input using the same method
	hashedInput, _ := HashSHA512(input)

	// Compare the hashed input with the provided hash
	return hashedInput == hashToCheck
}

func HashMD5(input string) (string, error) {
	hash := md5.New()

	// Write the input string to the hash
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", err
	}

	// Calculate the hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}

// check hash is same or not
func CheckMD5(input, hashToCheck string) bool {
	// Hash the input using the same method
	hashedInput, _ := HashMD5(input)

	// Compare the hashed input with the provided hash
	return hashedInput == hashToCheck
}
