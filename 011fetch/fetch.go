package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		// if !strings.HasPrefix(url, "http://") {
		// 	url = "http://" + url
		// }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("STATUS: ", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body) //==> copying directly to standard output.
		// b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("Runtime: %.2fs \n", secs)
}
