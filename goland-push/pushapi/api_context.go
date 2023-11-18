package pushapi

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type PushConstants struct {
	exampleVar  string
	env         string
	anotherVar  string
	debugMode   string
	databaseURL string
	apiKey      string
	Config      Configuration
	// Add other fields as needed
}

type Configuration struct {
	// Configuration fields
}

var instance *PushConstants
var once sync.Once

// LoadEnv reads the .env file and populates the struct
func loadEnv() *PushConstants {
	err := godotenv.Load(".env") // Specify the correct .env file name
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	envStruct := &PushConstants{}

	envStruct.exampleVar = os.Getenv("EXAMPLE_VAR")
	envStruct.anotherVar = os.Getenv("ANOTHER_VAR")
	envStruct.debugMode = os.Getenv("DEBUG_MODE")
	envStruct.databaseURL = os.Getenv("DATABASE_URL")
	envStruct.apiKey = os.Getenv("API_KEY")
	return envStruct
}

// GetInstance returns the singleton instance of APIContext
func InitializePushConstants(config Configuration) *PushConstants {
	once.Do(func() {
		instance = loadEnv()
	})
	return instance
}
