module pushapi-sdk

go 1.19

require (
	github.com/ethereum/go-ethereum v1.10.26
	github.com/joho/godotenv v1.5.1
)

require (
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.10.0
