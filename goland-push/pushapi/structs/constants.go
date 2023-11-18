package structs

// ENV - Supported Environments

const (
	PROD    string = "prod"
	STAGING string = "staging"
	DEV     string = "dev"
	LOCAL   string = "local" // This is for local development only
)

// ENCRYPTION_TYPE - Supported Encryptions for Push Profile
type ENCRYPTION_TYPE string

const (
	PGP_V1    ENCRYPTION_TYPE = "x25519-xsalsa20-poly1305"
	PGP_V2    ENCRYPTION_TYPE = "aes256GcmHkdfSha256"
	PGP_V3    ENCRYPTION_TYPE = "eip191-aes256-gcm-hkdf-sha256"
	NFTPGP_V1 ENCRYPTION_TYPE = "pgpv1:nft"
)

// MessageType - Supported Message Types for Push Chat
type MessageType string

const (
	TEXT          MessageType = "Text"
	IMAGE         MessageType = "Image"
	VIDEO         MessageType = "Video"
	AUDIO         MessageType = "Audio"
	FILE          MessageType = "File"
	GIF           MessageType = "GIF" // Deprecated - Use `MediaEmbed` Instead
	MEDIA_EMBED   MessageType = "MediaEmbed"
	META          MessageType = "Meta"
	REACTION      MessageType = "Reaction"
	RECEIPT       MessageType = "Receipt"
	USER_ACTIVITY MessageType = "UserActivity"
	INTENT        MessageType = "Intent"
	REPLY         MessageType = "Reply"
	COMPOSITE     MessageType = "Composite"
	PAYMENT       MessageType = "Payment"
)

const (
	INITIAL_PAGE int = 1
	LIMIT        int = 10
	LIMIT_MIN    int = 1
	LIMIT_MAX    int = 50

	DEFAULT_CHAIN_ID int = 11155111
	DEV_CHAIN_ID     int = 99999

	ENC_TYPE_V1 string = "x25519-xsalsa20-poly1305"
	ENC_TYPE_V2 string = "aes256GcmHkdfSha256"
	ENC_TYPE_V3 string = "eip191-aes256-gcm-hkdf-sha256"
	ENC_TYPE_V4 string = "pgpv1:nft"
)

var NON_ETH_CHAINS = []int{137, 80001, 56, 97, 10, 420, 1442, 1101, 421613, 42161}
var ETH_CHAINS = []int{1, 11155111}

// Defining environment types as constants
const (
	EnvProd    = "prod"
	EnvStaging = "staging"
	EnvDev     = "dev"
	EnvLocal   = "local"
)

// Defining the API base URL values
var ApiBaseUrl = map[string]string{
	PROD:    "https://backend.epns.io/apis",
	STAGING: "https://backend-staging.epns.io/apis",
	DEV:     "https://backend-dev.epns.io/apis",
	LOCAL:   "http://localhost:4000/apis",
}

// Defining blockchain network constants
const (
	EthMainnet          = "eip155:1"
	EthSepolia          = "eip155:11155111"
	PolygonMainnet      = "eip155:137"
	PolygonMumbai       = "eip155:80001"
	BscMainnet          = "eip155:56"
	BscTestnet          = "eip155:97"
	OptimismTestnet     = "eip155:420"
	OptimismMainnet     = "eip155:10"
	PolygonZkEvmTestnet = "eip155:1442"
	PolygonZkEvmMainnet = "eip155:1101"
	ArbitrumTestnet     = "eip155:421613"
	ArbitrumOneMainnet  = "eip155:42161"
)
