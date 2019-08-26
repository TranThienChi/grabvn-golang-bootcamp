package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/avast/retry-go"
	metricCollector "github.com/myteksi/hystrix-go/hystrix/metric_collector"

	"github.com/myteksi/hystrix-go/hystrix"
	"github.com/myteksi/hystrix-go/plugins"
)

const (
	url = "http://127.0.0.1:8080/test"
)

func init() {
	//  Set hystrix circuit breaker
	hystrix.ConfigureCommand("callEchoServerCommand", hystrix.CommandConfig{
		ErrorPercentThreshold: 30,
	})
}

func retries() (body []byte, err error) {
	err = retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			if resp.StatusCode != 200 {
				return errors.New("retries failed")
			}
			defer resp.Body.Close()

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return nil
		},
		retry.Attempts(2),
		retry.DelayType(retry.FixedDelay),
	)
	return
}

func main() {
	// Respond body
	var body []byte

	// Initialize statsd client
	c, err := plugins.InitializeStatsdCollector(&plugins.StatsdCollectorConfig{
		StatsdAddr: "localhost:8080",
		Prefix:     "myserver.hystrix",
	})
	if err != nil {
		log.Fatalf("could not initialize statsd client: %v", err)
	}

	metricCollector.Registry.Register(c.NewStatsdCollector)

	for i := 0; i < 100; i++ {
		j := i
		// Circuit breaker
		hystrix.Do("callEchoServerCommand", func() error {
			fmt.Printf("running job: %d\n", j)
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			if resp.StatusCode != 200 {
				// Retries inside circuit breaker
				body, err = retries()
				if err != nil {
					return errors.New("a chaos monkey broke your server")
				}
				fmt.Printf("retries running job: %d\n", j)
			}
			defer resp.Body.Close()

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return nil
		}, func(err error) error {
			fmt.Printf("fallback running job: %d\n", j)
			return err
		})
		fmt.Println(string(body))
	}
}
