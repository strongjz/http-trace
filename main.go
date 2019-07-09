package main

import (
	"flag"
	"fmt"
	"github.com/lixiangzhong/dnsutil"
	"github.com/tcnksm/go-httpstat"
	"io"
	"io/ioutil"
	"net/http/httputil"
	"net/url"

	"log"
	"net/http"
	"time"

)

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}



func main() {
	start := time.Now()

	nameServer := flag.String("nameserver", "8.8.8.8", "nameserver for lookup")

	testURL := flag.String("url", "https://contino.io/resources", "URL to test")
	method := flag.String("method", "GET", "Method to test, GET, POST, etc ")

	fs := flag.NewFlagSet("url", flag.ExitOnError)

	fs.Var(&URLValue{u}, "url", "URL to parse")

	fs.Parse([]string{"-url", *testURL})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

	flag.Parse()


	// Create a new HTTP request
	req, err := http.NewRequest(*method, *testURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}

	err = res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var dig dnsutil.Dig
	err = dig.SetDNS(*nameServer) //or ns.xxx.com
	if err != nil {
		log.Fatalln(err)
	}
	a, err := dig.A(*testURL)  // dig google.com @8.8.8.8
	if err != nil {
		log.Fatalln(err)
	}

	for _, r := range a {

		log.Printf("DNS result %v",r)
	}

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}

	// Show the results
	log.Printf("req: %v", string(requestDump))

	result.End(time.Now())

	log.Printf("%+v\n", result)


	log.Printf("Entire timing: %v", time.Since(start))

}