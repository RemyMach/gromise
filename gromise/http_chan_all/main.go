package main

import (
	"context"
	"fmt"
	"net/http"
)

func doRequest(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response code")
	}

	return nil
}

func main() {
	// Créez un contexte avec annulation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls := []string{
		"http://httpbin.org/status/404",
		"http://httpbin.org/delay/5",
		"http://httpbin.org/delay/10",
	}

	errCh := make(chan error, len(urls))

	for _, url := range urls {
		go func(url string) {
			errCh <- doRequest(ctx, url)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		if err := <-errCh; err != nil {
			// En cas d'erreur, annulez le contexte
			fmt.Printf("request failed: %v\n", err)
			return
		}
	}

	fmt.Println("all requests succeeded")
}
