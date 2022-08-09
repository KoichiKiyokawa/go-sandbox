package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	mu := sync.Mutex{}

	results := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {

			result, err := fetchTodo(i)
			if err != nil {
				return err
			}

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("results:", results)
}

func fetchTodo(id int) (string, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}
