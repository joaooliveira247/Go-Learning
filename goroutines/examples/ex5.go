package examples

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL string
	Size int
}

func RunEx5() {
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://apple.com",
		"https://microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string)  {
			res, err := http.Get(url)

			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)

			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL: url,
				Size: len(bs),
			}
		}(url)
	}
	var biggest HomePageSize

	for range urls {
		result := <- results
		
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("The biggest HomePage is:", biggest.URL)
}