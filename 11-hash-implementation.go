/***** ORIGINAL *****

package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func convSHA256(data []byte) []byte {
	hashish := sha256.Sum256(data)
	return hashish[:]
}

func main() {
	data, err := ioutil.ReadFile("d://DOCUMENTS [EXTHD]/tester.txt")
	// data, err := ioutil.ReadFile("/run/media/EXTHD/DOCUMENTS [EXTHD]/tester.txt")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	hasher := convSHA256(data)
	fmt.Printf("RESULT: %x", hasher[:])
}

/*
func TestNewSHA256(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out string
	}{
		{[]byte(""), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{[]byte("abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{[]byte("hello"), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := NewSHA256(tt.in)
			if hex.EncodeToString(result) != tt.out {
				t.Errorf("want %v; got %v", tt.out, hex.EncodeToString(result))
			}
		})
	}
}

**********/

/********** COPILOT'S CONTRIBUTION BASED OFF ORIGINAL CODE **********/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

// Hashes data using SHA-256 or SHA-512 based on user selection
func hashData(data []byte, algorithm string) []byte {
	switch algorithm {
	case "sha256":
		hash := sha256.Sum256(data)
		return hash[:]
	case "sha512":
		hash := sha512.Sum512(data)
		return hash[:]
	default:
		fmt.Println("Unsupported hashing algorithm. Defaulting to SHA-256.")
		hash := sha256.Sum256(data)
		return hash[:]
	}
}

// Encrypts data using AES
func encryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Generate random IV
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt data
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

// Decrypts data using AES
func decryptData(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Extract IV
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Decrypt data
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

// Generates a 32-byte key from a passphrase using SHA-256
func generateKey(passphrase string) []byte {
	return hashData([]byte(passphrase), "sha256")[:32]
}

func main() {
	// Prompt user for input file
	fmt.Print("Enter the file path: ")
	var filePath string
	fmt.Scanln(&filePath)

	// Read file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}

	// Hashing
	fmt.Print("Choose hashing algorithm (sha256/sha512): ")
	var algorithm string
	fmt.Scanln(&algorithm)

	hashedData := hashData(data, algorithm)
	fmt.Printf("Hashed Data (%s): %x\n", algorithm, hashedData)

	// Encryption
	fmt.Print("Enter passphrase for encryption: ")
	var passphrase string
	fmt.Scanln(&passphrase)

	key := generateKey(passphrase)
	encryptedData, err := encryptData(data, key)
	if err != nil {
		fmt.Println("Encryption ERROR:", err)
		return
	}
	fmt.Printf("Encrypted Data (hex): %s\n", hex.EncodeToString(encryptedData))

	// Optional decryption for verification
	fmt.Print("Decrypt data? (yes/no): ")
	var decryptChoice string
	fmt.Scanln(&decryptChoice)

	if decryptChoice == "yes" {
		decryptedData, err := decryptData(encryptedData, key)
		if err != nil {
			fmt.Println("Decryption ERROR:", err)
			return
		}
		fmt.Printf("Decrypted Data: %s\n", string(decryptedData))
	}

	// Write encrypted data to file
	fmt.Print("Enter output file path for encrypted data: ")
	var outputPath string
	fmt.Scanln(&outputPath)

	err = ioutil.WriteFile(outputPath, encryptedData, 0644)
	if err != nil {
		fmt.Println("File write ERROR:", err)
		return
	}

	fmt.Println("Encrypted data saved successfully.")
}
