package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	port := flag.String("l", "8080", "listen on port, e.g., 8080")
	directory := flag.String("d", ".", "directory to server HTML files from, e.g. /var/www")
	redirect := flag.String("r", "8545", "redirect port to redirect to, e.g., 8545. This will"+
		"redirect http://localhost:8080/8545 to http://localhost:8545")

	flag.Parse()
	fmt.Printf("Static file HTTP server and reverse proxy for localhost. This tool exists due to CORS.\n")
	flag.PrintDefaults()

	redirUrl := "http://localhost:" + *redirect
	origin, err := url.Parse(redirUrl)

	if err != nil {
		log.Fatalf("Cannot parse URL %v - %v", redirUrl, err)
	}

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host
		req.URL.Path = origin.Path
	}

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/"+*redirect, func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *directory+"/"+r.URL.Path[1:])
	})

	fmt.Printf("Serving: %v on http://localhost:%v, redirecting http://localhost:%v/%v to http://localhost:%v",
		*directory, *port, *port, *redirect, *redirect)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
