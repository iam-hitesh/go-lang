package main

import (
  "fmt"
  "net/http"
)

func main(){
  client := &http.Client{}
  resp, err := client.Get("http://hiteshyadav.me")
  fmt.Println(resp)
  fmt.Println(err)
}

/**
Ouput

html of hiteshyadav.me website
**/
