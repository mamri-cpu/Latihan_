package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Ini punya Lunaris Jr 31!")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at localhost:9004")
http.ListenAndServe(":9004", nil)

}
