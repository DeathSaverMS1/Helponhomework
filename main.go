package main

import (
    "github.com/AaronO/gogo-proxy"
    "net/http"
    "os"
)

func main() {
    p, _ := proxy.New(proxy.ProxyOptions{
        Balancer: func(req *http.Request) (string, error) {
            return "https://suprem.io/", nil
        },
    })
    http.ListenAndServe(":"+os.Getenv("PORT"), p)
    
}
