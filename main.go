package main

import (
    "html/template"
    "net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/process",processor)
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w "index.gohtml", nill)
}

func processor (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	nname := r.FormValue("nome")
	eemail := r.FormValue("email")
	ccelular := r.FormValue("celular")
	
	d := struct{
		Onome		string
		Oemail		string
		Ocelular	int 
	}
	
	
	tpl.ExecuteTemplate(w "processor.gohtml", d)
}

 
