package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func parseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("URL parse error %s\n", err.Error())
	}
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
		log.Fatalf("http request error %s\n", err.Error())
	}

	if resp.StatusCode != 200 {
		log.Fatalf("return code is not 200: %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}
	parsedURL := parseURL(os.Args[1])

	body := makeRequest(parsedURL.String())
	fmt.Println(body)
	os.Exit(0)
}
