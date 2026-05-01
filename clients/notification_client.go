package clients

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

func SendNotification(orderID uint, status string) error {
	client := resty.New()

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		fmt.Println("[Resty] Requesting:", req.Method, req.URL)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		fmt.Println("[Resty] Response Code:", resp.StatusCode())
		return nil
	})

	notifyURL := os.Getenv("NOTIFICATION_SERVICE_URL")
	if notifyURL == "" {
		notifyURL = "http://localhost:8081/notify"
	}

	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"order_id": orderID,
			"status":   status,
		}).
		Post(notifyURL)

	if err != nil {
		return err
	}

	return nil
}
