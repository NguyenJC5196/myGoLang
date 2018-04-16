package main
import ("fmt"
		"net/http")

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000",nil)
}
func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Wow, what now?")
}
func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "A beautiful design from Water_Ng")
}