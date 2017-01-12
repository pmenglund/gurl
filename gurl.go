package main

import "fmt"
import "io/ioutil"
import "net/url"
import "os"
import "net/http"

func parseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		fmt.Printf("URL parse error %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("url:", u)
	return u
}

func makeRequest(parsedURL string) string {
	/* TODO: add an option to follow redirects
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get(u.String())
	*/

	resp, err := http.Get(parsedURL)
	if err != nil {
		fmt.Printf("http request error %s\n", err.Error())
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("return code is not 200: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s\n", os.Args[0])
		os.Exit(1)
	}
	// rawURL = os.Args[1]
	rawURL := "http://localhost:8080/api/v1/ping"
	parsedURL := parseURL(rawURL)

	body := makeRequest(parsedURL.String())
	fmt.Println(body)
	os.Exit(0)
}
