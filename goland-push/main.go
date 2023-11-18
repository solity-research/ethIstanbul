package main

import "pushapi-sdk/pushapi"

func main() {

	config := pushapi.Configuration{
		// Configuration settings
	}
	ctx := pushapi.InitializePushConstants(config)

	ctx.Subscribe()
}
