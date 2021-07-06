package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"math/rand"
)

var backendHosts []Host
type Host struct {
	url *url.URL
}

func init() {
	backendHosts = append(backendHosts, Host{
		url: &url.URL{
			Scheme: "http",
			Host: "localhost:8000",
		},
	})
	backendHosts = append(backendHosts, Host{
		url: &url.URL{
			Scheme: "http",
			Host: "localhost:8001",
		},
	})
	backendHosts = append(backendHosts, Host{
		url: &url.URL{
			Scheme: "http",
			Host: "localhost:8002",
		},
	})
}

func LbHandler(w http.ResponseWriter, r *http.Request){
	index := rand.Intn(len(backendHosts))
	fmt.Println(index)
	randomHost := backendHosts[index]
	httputil.NewSingleHostReverseProxy(randomHost.url).ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", LbHandler)
	http.ListenAndServe(":8005", mux)
}