package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
    sum=0
    for i := 0; i < 1000000; i++ {
        sum += i
    }
    fmt.Fprint(w, "Hello, world!", sum)
}