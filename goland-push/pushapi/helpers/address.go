package pushapi

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
)

func isValidETHAddress(address string) bool {
	if isValidCAIP10NFTAddress(address) {
		return true
	}
	if strings.Contains(address, "eip155:") {
		splittedAddress := strings.Split(address, ":")
		if len(splittedAddress) == 3 {
			return common.IsHexAddress(splittedAddress[2])
		}
		if len(splittedAddress) == 2 {
			return common.IsHexAddress(splittedAddress[1])
		}
	}
	return common.IsHexAddress(address)
}

func isValidCAIP10NFTAddress(wallet string) bool {
	walletComponents := strings.Split(wallet, ":")
	if len(walletComponents) != 5 && len(walletComponents) != 6 {
		return false
	}
	if walletComponents[0] != "nft" || walletComponents[1] != "eip155" {
		return false
	}
	// Add additional checks for numeric values and address validation
	return common.IsHexAddress(walletComponents[3])
}

func isValidNFTCAIP10Address(realCAIP10 string) bool {
	walletComponent := strings.Split(realCAIP10, ":")
	if len(walletComponent) != 3 || walletComponent[0] != "eip155" {
		return false
	}
	_, err := strconv.Atoi(walletComponent[1])
	if err != nil {
		return false
	}
	return common.IsHexAddress(walletComponent[2])
}

type AddressValidator func(address string) bool

var AddressValidators = map[string]AddressValidator{
	"eip155": func(address string) bool {
		return isValidETHAddress(address)
	},
	// Additional chain validators can be added here
}

func validateCAIP(addressInCAIP string) bool {
	parts := strings.Split(addressInCAIP, ":")
	if len(parts) < 3 {
		return false
	}
	blockchain, address := parts[0], parts[2]
	if !isValidCAIP10NFTAddress(addressInCAIP) {
		if validatorFn, ok := AddressValidators[blockchain]; ok {
			return validatorFn(address)
		}
	}
	return false
}

func getCAIPDetails(addressInCAIP string) (*structs.CAIPDetailsType, error) {
	if !validateCAIP(addressInCAIP) {
		return nil, errors.New("Invalid CAIP address")
	}

	parts := strings.Split(addressInCAIP, ":")
	blockchain, networkId, address := parts[0], parts[1], parts[2]

	return &structs.CAIPDetailsType{
		Blockchain: blockchain,
		NetworkId:  networkId,
		Address:    address,
	}, nil
}

func (m *PushConstants) getFallbackETHCAIPAddress(address string) string {
	chainId := "1" // default for PROD
	if env == DEV || env == STAGING || env == LOCAL {
		chainId = "11155111" // example chain ID for non-prod environments
	}
	return "eip155:" + chainId + ":" + address
}

func getCAIPAddress(address string) (string, error) {
	if isValidCAIP10NFTAddress(address) {
		// Assuming getUserDID function exists and returns a string or an error.
		return getUserDID(address)
	}
	if validateCAIP(address) {
		return address, nil
	}
	if isValidETHAddress(address) {
		return getFallbackETHCAIPAddress(address), nil
	}
	return "", errors.New("Invalid Address: " + address)
}
