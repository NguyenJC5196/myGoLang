package main

import ("fmt"
		"net/http")

func main(){
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `<h1>Hello guys</h1>
		<p>Go is not quite diffcult as I thought</p>
		<p>You %s even add %s </p>`, "can","<strong>variables</strong>")

	/*
	The example above is short-hand write to write multi-line
	*/

	
	// fmt.Fprintf(w, "<h1>Hello guys</h1>")
	// fmt.Fprintf(w, "<p>Go is not quite diffcult as I thought</p>")
	// fmt.Fprintf(w, "<p>You %s even add %s </p>", "can","<strong>variables</strong>")
}