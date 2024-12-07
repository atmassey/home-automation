package seam

import (
	"net/http"
	"time"

	goclient "github.com/seamapi/go/client"
)

func Authenticate(apiKey string) *goclient.Client {
	client := goclient.NewClient(goclient.WithApiKey(apiKey), goclient.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}))
	return client
}
