package main

import ("fmt"
		"net/http"
		"io/ioutil")

func main(){
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(response.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	response.Body.Close()
}