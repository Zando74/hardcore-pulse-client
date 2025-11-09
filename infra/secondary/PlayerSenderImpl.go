package secondary

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"deathlog-tracker/domain/entity"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


const (
	BACKEND_URL = "https://hardcore-pulse.com/api/players"
	API_KEY = "DfLHO9cnozarqVgFBASiFAhson6quObK"
	AES_KEY = "bxuXHNFuZMXgW9A3WDwk8XIbhvvvquRi"
)

func encrypt(plaintext []byte) string {
    aes, err := aes.NewCipher([]byte(AES_KEY))
    if err != nil {
        panic(err)
    }

	
	gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }

	nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(fmt.Errorf("read nonce: %w", err))
	}

    ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	out := append(nonce, ciphertext...)

	hex := hex.EncodeToString(out)

    return hex
}

type PlayerEncrypted struct {
	Encrypted string `json:"encrypted"`
}

type PlayerSenderImpl struct {}


func (p *PlayerSenderImpl) SendBatch(players []entity.Player) bool {
	data, err := json.Marshal(players)
	if err != nil {
		log.Fatalf("Failed to marshal players: %v", err)
	}

	payload, err := json.Marshal(PlayerEncrypted{Encrypted: encrypt(data)})
	if err != nil {
		log.Fatalf("Failed to marshal players: %v", err)
	}

	req, err := http.NewRequest(
		"POST", 
		BACKEND_URL, 
		bytes.NewBuffer(
			[]byte(payload),
		),
	)
	if err != nil {
		log.Fatalf("[DEATHLOG-TRACKER]  Failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[DEATHLOG-TRACKER]  Request failed")
		return false
	}
	defer resp.Body.Close()

	message, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[DEATHLOG-TRACKER] Failed to read response body")
		return false
	}

	if resp.StatusCode != http.StatusCreated {
		log.Println("[DEATHLOG-TRACKER] Bad response status: " + resp.Status)
		log.Printf("[DEATHLOG-TRACKER] Response body: %s", message)
	} else {
		log.Println("[DEATHLOG-TRACKER] Players sent successfully")
		return true
	}

	return false
}