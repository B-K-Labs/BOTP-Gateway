package http_request

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
		SetUserAgent("my-custom-client").
		SetCommonContentType("application/json").
		// TODO: Set common bearer auth token
		SetCommonBearerAuthToken("my-token").
		SetBaseURL(baseUrl).
		SetTimeout(30 * time.Second)
}

func GetHttpClient() *req.Client {
	if client == nil {
		Initialize()
	}
	return client
}
