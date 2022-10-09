package function

import (
	"fmt"
	"log"
	"net/http"
	"os"

	nats "github.com/nats-io/nats.go"
	handler "github.com/openfaas/templates-sdk/go-http"
)

var (
	subject        = "nats-test"
	defaultMessage = "Relay test"
)

// Handle a serverless request
func Handle(req handler.Request) (handler.Response, error) {
	log.Printf("Received: %q", string(req.Body))

	// if val, ok := os.LookupEnv("wait"); ok && len(val) > 0 {
	// 	parsedVal, _ := time.ParseDuration(val)
	// 	log.Printf("Waiting for %s before returning", parsedVal.String())
	// 	time.Sleep(parsedVal)
	// }

	msg := defaultMessage

	natsURL := nats.DefaultURL
	val, ok := os.LookupEnv("nats_url")
	if ok {
		natsURL = val
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		errMsg := fmt.Sprintf("can not connect to nats: %s", err)
		log.Printf(errMsg)
		r := handler.Response{
			Body:       []byte(errMsg),
			StatusCode: http.StatusInternalServerError,
		}
		return r, err
	}
	defer nc.Close()

	log.Printf("Publishing %d bytes to: %q\n", len(msg), subject)

	err = nc.Publish(subject, []byte(msg))
	if err != nil {
		log.Println(err)

		r := handler.Response{
			Body:       []byte(fmt.Sprintf("can not publish to nats: %s", err)),
			StatusCode: http.StatusInternalServerError,
		}
		return r, err
	}

	return handler.Response{
		Body:       []byte(fmt.Sprintf("Received: %q", string(req.Body))),
		StatusCode: http.StatusOK,
	}, nil
}
