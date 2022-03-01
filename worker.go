package main

import (
	"crypto/md5"
	"fmt"
)

type ClientInterface interface {
	DoRequest(url string) ([]byte, error)
}

type UrlModifierInterface interface {
	Modify(address string) (string, error)
}

type Worker struct {
	number      int
	client      ClientInterface
	urlModifier UrlModifierInterface
}

type Result struct {
	Url  string
	Body []byte
}

func NewWorker(number int, client ClientInterface, urlModifier UrlModifierInterface) *Worker {
	return &Worker{number: number, client: client, urlModifier: urlModifier}
}

func (w *Worker) Run(urls []string) {
	input := make(chan string)
	output := make(chan Result)

	for i := 0; i < w.number; i++ {
		go w.doRequests(input, output)
	}

	for _, url := range urls {
		input <- url
	}
	close(input)

	for i := 0; i < len(urls); i++ {
		r := <-output
		fmt.Printf("%s %x\n", r.Url, md5.Sum(r.Body))
	}

}

func (w *Worker) doRequests(input <-chan string, output chan<- Result) {
	for address := range input {
		url, err := w.urlModifier.Modify(address)
		if err != nil {
			//todo error handling
		}
		result, err := w.client.DoRequest(url)
		if err != nil {
			//todo error handling
		}
		output <- Result{
			Url:  url,
			Body: result,
		}
	}
}
