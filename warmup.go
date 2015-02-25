package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

var (
	delay    = flag.Int("delay", 29000, "Delay (in ms) between items in the request list")
	filename = flag.String("filename", "urls.json", "List of URLs")
	timeFile = flag.Bool("time-file", true, "If you want warmup.txt beeing created with last run time")
)

type headers map[string]string
type urls map[string]headers

func run(data urls) {
	for url, heads := range data {
		doRequest(url, heads)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	file, err := ioutil.ReadFile(*filename)

	if err != nil {
		log.Fatal(err)
	}

	var data urls

	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Fatal(err)
	}

	for true {
		run(data)
		fmt.Println(fmt.Sprintf("waiting %ds...", *delay/1000))

		if *timeFile {
			t := time.Now().UTC()
			ioutil.WriteFile("warmup.txt", []byte(t.String()), 0644)
		}

		if *delay > 0 {
			time.Sleep(time.Duration(*delay) * time.Millisecond)
		}
	}
}

// Do request against url with headers
func doRequest(url string, heads headers) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", url))
		return
	}

	for key, value := range heads {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", url))
		return
	}

	defer resp.Body.Close()

	fmt.Println(fmt.Sprintf("done: %s", url))
}
