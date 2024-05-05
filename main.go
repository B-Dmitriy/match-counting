package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// defer resp.Body.Close()
// Вызовы отложенной строки. Close() внутри цикла for не выполняются до тех пор, пока функция не завершит свое выполнение.
// Не в конце каждого шага цикла for !!!
// Такая реализация может привести к переполнению стека функции и другим проблемам.

const Substring = "Go"

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Error reading stdin: ", err)
	}

	fileData := string(stdin)
	urls := strings.Split(fileData, "\n")

	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		go getAndCheckURLBody(&wg, url, Substring)
	}

	wg.Wait()
}

func getAndCheckURLBody(wg *sync.WaitGroup, url, substr string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Count for %s: unknown [request error]: %s\n", url, err)
		wg.Done()
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Count for %s: unknown [reading response body error]: %s\n", url, err)
		wg.Done()
		return
	}

	fmt.Printf("Count for %s: %d\n", url, strings.Count(string(body), substr))
	wg.Done()
}
