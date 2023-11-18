package structs

type CAIPDetailsType struct {
	Blockchain string
	NetworkId  string
	Address    string
}

type IUser struct {
	MsgSent             int
	MaxMsgPersisted     int
	DID                 string
	Wallets             string
	Profile             UserProfile
	EncryptedPrivateKey string
	PublicKey           string
	VerificationProof   string
	Origin              *string // Pointer for nullable string
}

type UserProfile struct {
	Name                     *string
	Desc                     *string
	Picture                  *string
	ProfileVerificationProof *string
	BlockedUsersList         *[]string
}

type AccountEnvOptionsType struct {
	Account string
}

type ProfileData struct {
	CAIP10              string
	DID                 string
	PublicKey           string
	EncryptedPrivateKey string
}
