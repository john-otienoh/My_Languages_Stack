package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {
	// Making a GET Request
	response, err := http.Get("https://ifconfig.co/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	// Making a POST Request
	postData := strings.NewReader(`{"some": "json"}`)
	response1, err := http.Post("https://httpbin.org/post", "application/json", postData)
	if err != nil {
		log.Fatal(err)
	}
	defer response1.Body.Close()
	body1, err := io.ReadAll(response1.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body1)

	// Gaining Further Control over HTTP Requests
	client := &http.Client{Timeout: 50 * time.Millisecond}
	request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	if err != nil {
		log.Fatal(err)
	}
	response2, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response2.Body.Close()
	body2, err := io.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body2)

	// Debugging HTTP Requests
	debug := os.Getenv("DEBUG")
	if debug == "]" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", debugRequest)
	}
	response3, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response3.Body.Close()
	if debug == "1" {
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugResponse)
	}
	body3, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body3)
	// Dealing with timeouts
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	cl := &http.Client{
		Transport: tr,
	}
	request1, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	if err != nil {
		log.Fatal(err)
	}
	response5, err := cl.Do(request1)
	if err != nil {
		log.Fatal(err)
	}
	defer response5.Body.Close()
	body4, err := io.ReadAll(response5.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body4)
}
