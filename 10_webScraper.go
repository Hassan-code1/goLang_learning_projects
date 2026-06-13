package main

import(
    "fmt"
    "io"
    "log"
    "net/http"
	"os"
	"sync"
)

type Result struct {
	Url string
	StatusCode int
	ContentLen int
	Err error
}

func fetchUrl(url string, ch chan<- Result){
	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{
			Url : url,
			Err :  err,
		}
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		ch<- Result{
			Url : url,
			Err :  err,
		}
		return
	}

	ch<- Result{
		Url : url,
		StatusCode : resp.StatusCode,
		ContentLen : len(body),
	}
}

func main(){
	urls := []string{
		"https://golang.org",
		"https://go.dev",
		"https://example.com",
	}

	resultCh := make(chan Result)

	var wg sync.WaitGroup
	var results []Result
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			fetchUrl(url, resultCh)
		}(url)
	}
	go func(){
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh{
		mu.Lock()
		results = append(results, result)
		mu.Unlock()
	}

	file, err := os.Create("results.txt")

	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	for _, r := range results {
		if r.Err != nil {
			fmt.Fprintf(
				file,
				"Url : %s\nError: %v\n\n",
				r.Url, r.Err,
			)
			continue
		}

		fmt.Fprintf(
			file,
			"Url: %s\nStatus: %d\nContent Length: %d\n\n",
			r.Url, r.StatusCode, r.ContentLen,
		)
	}
	fmt.Println("Results written to results.txt")
}