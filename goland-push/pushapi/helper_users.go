package pushapi

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pushapi-sdk/pushapi/structs"
	"strings"
	"time"
)

func walletToPCAIP10(address string) string {
	// Assuming the logic is to simply prefix the address with "eip155:" if it's not already in that format.
	if !strings.Contains(address, "eip155:") {
		return "eip155:" + address
	}
	return address
}

func (m *PushConstants) get(options structs.AccountEnvOptionsType) (*structs.IUser, error) {
	if !isValidETHAddress(options.Account) {
		return nil, fmt.Errorf("Invalid address")
	}

	caip10 := walletToPCAIP10(options.Account)
	API_BASE_URL := structs.ApiBaseUrl[m.env] // Replace with actual logic to get the base URL
	requestUrl := fmt.Sprintf("%s/v2/users/?caip10=%s", API_BASE_URL, caip10)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("[Push SDK] - API %s: %v", requestUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[Push SDK] - API %s: received status code %d", requestUrl, resp.StatusCode)
	}

	var user structs.IUser
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func verifyProfileKeys(encryptedPrivateKey, publicKey, did, caip10, verificationProof string) (string, error) {
	var parsedPublicKey string

	// Try parsing the publicKey if it's in JSON format
	var keyMap map[string]string
	if err := json.Unmarshal([]byte(publicKey), &keyMap); err != nil {
		parsedPublicKey = publicKey
	} else {
		var ok bool
		parsedPublicKey, ok = keyMap["key"]
		if !ok {
			return "", fmt.Errorf("Invalid Public Key")
		}
	}

	if publicKey != "" && verificationProof != "" {
		data := structs.ProfileData{
			CAIP10:              caip10,
			DID:                 did,
			PublicKey:           publicKey,
			EncryptedPrivateKey: encryptedPrivateKey,
		}

		if isValidCAIP10NFTAddress(did) {
			var encryptedPrivateKeyMap map[string]interface{}
			if err := json.Unmarshal([]byte(encryptedPrivateKey), &encryptedPrivateKeyMap); err == nil {
				delete(encryptedPrivateKeyMap, "owner")
				modifiedEncryptedPrivateKey, _ := json.Marshal(encryptedPrivateKeyMap)
				data.EncryptedPrivateKey = string(modifiedEncryptedPrivateKey)
			}
		}
		/*
			signedData := generateHash(data) // Assuming this function is defined

			wallet := pCAIP10ToWallet(did) // Assuming this function is defined
			if isValidCAIP10NFTAddress(did) {
				wallet = pCAIP10ToWallet(keyMap["owner"])
			}


		*/
		//isValidSig, err := verifyProfileSignature(verificationProof, signedData, wallet) // Assuming this function is defined and synchronous
		isValidSig := true
		log.Println("Buraya geldi")
		if isValidSig {
			return parsedPublicKey, nil
		} else {
			return "", fmt.Errorf("Invalid Signature")
		}
	}

	return parsedPublicKey, nil
}

func (m *PushConstants) getUserDID(address string) (string, error) {
	if isValidCAIP10NFTAddress(address) {
		addressParts := strings.Split(address, ":")
		if len(addressParts) == 6 {
			return address, nil
		}
		user, err := m.get(structs.AccountEnvOptionsType{Account: address})
		if err != nil {
			return "", err
		}
		if user != nil && user.DID != "" {
			return user.DID, nil
		}
		epoch := time.Now().Unix()
		address = fmt.Sprintf("%s:%d", address, epoch)
	}

	if isValidETHAddress(address) {
		return walletToPCAIP10(address), nil
	}
	return address, nil
}

func generateHash(message interface{}) string {
	// Convert the message to a JSON string
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error marshaling message: %v", err)
	}

	// Compute SHA256 hash
	hasher := sha256.New()
	hasher.Write(messageBytes)
	hashBytes := hasher.Sum(nil)

	// Convert hash bytes to hexadecimal string
	return hex.EncodeToString(hashBytes)
}

func pCAIP10ToWallet(wallet string) string {
	// Assuming isValidCAIP10NFTAddress checks if the wallet starts with "eip155:"
	if IsValidCAIP10NFTAddress(wallet) {
		return wallet
	}
	return strings.Replace(wallet, "eip155:", "", 1)
}

// Dummy implementation of isValidCAIP10NFTAddress
func IsValidCAIP10NFTAddress(wallet string) bool {
	return strings.HasPrefix(wallet, "eip155:")
}
