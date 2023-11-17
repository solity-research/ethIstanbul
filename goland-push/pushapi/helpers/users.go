package pushapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func get(options AccountEnvOptionsType) (*IUser, error) {
	if !isValidETHAddress(options.Account) {
		return nil, fmt.Errorf("Invalid address")
	}

	caip10 := walletToPCAIP10(options.Account)
	API_BASE_URL := getAPIBaseUrls(options.Env) // Replace with actual logic to get the base URL
	requestUrl := fmt.Sprintf("%s/v2/users/?caip10=%s", API_BASE_URL, caip10)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("[Push SDK] - API %s: %v", requestUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[Push SDK] - API %s: received status code %d", requestUrl, resp.StatusCode)
	}

	var user IUser
	bodyBytes, err := ioutil.ReadAll(resp.Body)
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
		data := ProfileData{
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

		signedData := generateHash(data) // Assuming this function is defined

		wallet := pCAIP10ToWallet(did) // Assuming this function is defined
		if isValidCAIP10NFTAddress(did) {
			wallet = pCAIP10ToWallet(keyMap["owner"])
		}

		isValidSig, err := verifyProfileSignature(verificationProof, signedData, wallet) // Assuming this function is defined and synchronous
		if err != nil {
			return "", err
		}
		if isValidSig {
			return parsedPublicKey, nil
		} else {
			return "", fmt.Errorf("Invalid Signature")
		}
	}

	return parsedPublicKey, nil
}

func getUserDID(address string) (string, error) {
	if isValidCAIP10NFTAddress(address) {
		addressParts := strings.Split(address, ":")
		if len(addressParts) == 6 {
			return address, nil
		}
		user, err := get(address)
		if err != nil {
			return "", err
		}
		if user != nil && user.Did != "" {
			return user.Did, nil
		}
		epoch := time.Now().Unix()
		address = fmt.Sprintf("%s:%d", address, epoch)
	}

	if isValidETHAddress(address) {
		return walletToPCAIP10(address), nil
	}
	return address, nil
}
