package http_client

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req/v3"
)

// Make http request from fiber context
func MakeRequestFromFiber(c *fiber.Ctx, path string) error {
	client := GetHttpClient()
	request := client.R().
		SetPathParams(c.AllParams()).
		SetBody(c.Body()).
		SetHeaders(c.GetReqHeaders()).
		EnableDump()

	// Get response from api
	var (
		resp *req.Response
		err  error
	)
	// Mapping fiber method to http client method
	switch c.Method() {
	case fiber.MethodGet:
		resp, err = request.Get(path)
	case fiber.MethodPost:
		resp, err = request.Post(path)
	case fiber.MethodPut:
		resp, err = request.Put(path)
	case fiber.MethodDelete:
		resp, err = request.Delete(path)
	case fiber.MethodPatch:
		resp, err = request.Patch(path)
	default:
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"success": false,
			"message": "Method not allowed",
		})
	}

	// Logging http request
	log.Printf("[%s] Requesting to %s", request.Method, request.URL)

	// Force to return error if request failed
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gateway error",
		})
	}

	// Return response from api
	return c.Status(resp.StatusCode).Send([]byte(resp.String()))
}

// Make http request from fiber context with same path
func MakeRequestFromFiberWithSamePath(c *fiber.Ctx) error {
	// Get current path of api gateway
	path := c.Path()
	return MakeRequestFromFiber(c, path)
}
