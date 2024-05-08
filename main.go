package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

const Substring = "Go"
const K = 5

func main() {
	var total int
	semaphore := make(chan struct{}, K)
	scanner := bufio.NewScanner(os.Stdin)

	wg := sync.WaitGroup{}

	for scanner.Scan() {
		wg.Add(1)
		semaphore <- struct{}{}
		go getAndCheckURLBody(&wg, &total, semaphore, scanner.Text(), Substring)
	}

	wg.Wait()
	fmt.Printf("Total: %d\n", total)
}

func getAndCheckURLBody(wg *sync.WaitGroup, total *int, ch chan struct{}, url, substr string) {
	defer wg.Done()
	defer func() { <-ch }()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Count for %s: unknown [request error]: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Count for %s: unknown [reading response body error]: %s\n", url, err)
		return
	}

	count := strings.Count(string(body), substr)
	*total = *total + count

	fmt.Printf("Count for %s: %d\n", url, count)
}
