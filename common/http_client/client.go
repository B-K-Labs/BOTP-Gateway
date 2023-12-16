package http_client

import (
	"botp-gateway/common/constants"
	"botp-gateway/config"
	"time"

	"github.com/imroc/req/v3"
)

var client *req.Client

func Initialize() {
	// Load base url from .env file
	baseUrl := config.Env(constants.ENV_KEY_BOTP_BACKEND_URL)
	// Initialize http client
	client = req.C().
		// TODO: Remove this line when not in development mode
		DevMode().
		// TODO: Set common user agent
		SetUserAgent("my-custom-client").
		SetCommonContentType("application/json").
		// TODO: Set common bearer auth token
		SetCommonBearerAuthToken("my-token").
		SetBaseURL(baseUrl).
		SetTimeout(30 * time.Second).
		OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			// TODO: Do something before request
			return nil
		}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			// TODO: Do something after response
			return nil
		})
}

func GetHttpClient() *req.Client {
	if client == nil {
		Initialize()
	}
	return client
}
