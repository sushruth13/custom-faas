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
	defaultMessage = "ERROR_FN"
)

// Handle a serverless request
func Handle(req handler.Request) (handler.Response, error) {
	log.Printf("Processing input for %q", string(req.Body))

	if val, ok := os.LookupEnv("fn_name"); ok && len(val) > 0 {
		if string(req.Body) == val {
			// Correct function is being called
			log.Printf("Running function code for %q", string(req.Body))

			// Function code here

			msg := defaultMessage

			if val2, ok := os.LookupEnv("fn_to_call"); ok && len(val) > 0 {
				msg = val2
			}

			// Lookup nats URL
			natsURL := nats.DefaultURL
			val, ok := os.LookupEnv("nats_url")
			if ok {
				natsURL = val
			}

			// Connect to nats
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

			// Publish message to nats
			log.Printf("Publishing %d bytes to: %q\n", len(msg), subject)

			// Respond error if needed
			err = nc.Publish(subject, []byte(msg))
			if err != nil {
				log.Println(err)

				r := handler.Response{
					Body:       []byte(fmt.Sprintf("can not publish to nats: %s", err)),
					StatusCode: http.StatusInternalServerError,
				}
				return r, err
			}
		}
	}

	return handler.Response{
		Body:       []byte(fmt.Sprintf("Successfully processed %q", string(req.Body))),
		StatusCode: http.StatusOK,
	}, nil
}
