package PushAPI

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type PushConstants struct {
	exampleVar  string
	anotherVar  string
	debugMode   string
	databaseURL string
	apiKey      string
	// Add other fields as needed
}

// InitializePushConstants is the constructor function for MyEnvStruct
func InitializePushConstants() *PushConstants {
	envStruct := &PushConstants{}
	envStruct.loadEnv()
	return envStruct
}

// LoadEnv reads the .env file and populates the struct
func (m *PushConstants) loadEnv() {
	err := godotenv.Load(".env") // Specify the correct .env file name
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	m.exampleVar = os.Getenv("EXAMPLE_VAR")
	m.anotherVar = os.Getenv("ANOTHER_VAR")
	m.debugMode = os.Getenv("DEBUG_MODE")
	m.databaseURL = os.Getenv("DATABASE_URL")
	m.apiKey = os.Getenv("API_KEY")
	// Load other environment variables here
}

func (m *PushConstants) Subscribe() string {
	return ""
}
