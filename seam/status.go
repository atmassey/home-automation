package seam

import (
	goclient "github.com/seamapi/go/client"
)

func Authenticate(apiKey string) *goclient.Client {
	client := goclient.NewClient(goclient.WithApiKey(apiKey))
	return client
}
