package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"sync"
)

func main() {
	parallel := flag.Int("parallel", 10, "Parallelisation")
	flag.Parse()

    argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	websites := argsWithoutProg

	if argsWithoutProg[0] == "-parallel" {
		websites = argsWithoutProg[2:]
	}
	
	var wg sync.WaitGroup
	wg.Add(len(websites))
	guard := make(chan struct{}, *parallel)

	for _, website := range websites {
		guard <- struct{}{} 
		go func(site string) {
			defer wg.Done()
			url, md5String := MakeRequestAndGetBodyMD5Checksum(site)
			fmt.Println(url, md5String)
			<-guard
		}(website)
	}
	wg.Wait()
}

func MakeRequestAndGetBodyMD5Checksum(url string) (string, string){

	new_url := url
	if !strings.HasPrefix(url, "http") {
		new_url = "http://" + url
	}

	resp, err := http.Get(new_url)
	if err != nil {
		return new_url, "ERROR"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return new_url, "ERROR"
	}

	hasher := md5.New()
	hasher.Write(body)
	resultMd5Str := hex.EncodeToString(hasher.Sum(nil))
	
	return new_url, resultMd5Str
}
