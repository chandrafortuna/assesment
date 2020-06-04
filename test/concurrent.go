package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()

	// product test concurrent
	var jsonStr = []byte(`
	{
		"customerId": "user01",
		"productVariantId": 1,
		"qty": 1
	}
	`)
	r := bytes.NewBuffer(jsonStr)

	// ball test concurrent
	// r := strings.NewReader("")

	resp, _ := http.Post(url, "application/json", r)
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go MakeRequest(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// time ./concurrent http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer http://localhost:8000/balls/addToContainer

// time ./concurrent http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve http://localhost:8000/productVariant/reserve
