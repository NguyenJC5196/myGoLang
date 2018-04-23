package main
import ("fmt"
		"net/http"
		"html/template")
type NewsEggPage struct{
	Title string
	News string
}
func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.HandleFunc("/egg", news_egg_handlder)
	http.ListenAndServe(":8000",nil)
}
func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Wow, what now?")
}
func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "A beautiful design from Water_Ng")
}
func news_egg_handlder(w http.ResponseWriter, r *http.Request){
	p := NewsEggPage{Title:"Amazing News", News: "some news"}
	t, err := template.ParseFiles("basictemplate.html")
	// fmt.Println(err)

	// this will print out the error of executtion
	fmt.Println(t.Execute(w, p))
}