package main

import "fmt"
import "net/http"

func main() {
    //resp, _ := http.Get("http://6play.fr/")

    //wanna control header
    client:=&http.Client{}
    req, err := http.NewRequest("GET", "http://jsonapi.org/", nil)
    req.Header.Add("Accept","application/json")
    req.Header.Add("Content-Type","application/vnd.api+json")
    resp, err := client.Do(req)

    fmt.Println(resp,err)
}
