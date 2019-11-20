package main

import (
    "html/template"
    "net/http"
	"database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    	int
    Name  	string
    Email 	string
	Celular string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "A@a2019!"
    dbName := "goblog"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id int
        var name, email, cell string
        err = selDB.Scan(&id, &name, &email, &cell)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
		emp.Email = email
        emp.Celular = cell
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, email, cell string
        err = selDB.Scan(&id, &name, &email, &cell)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
		emp.Email = email
        emp.Celular = cell
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, email, cell string
        err = selDB.Scan(&id, &name, &email, &cell)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
		emp.Email = email
        emp.Celular = cell
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        email := r.FormValue("email")
		cell := r.FormValue("celular")
		
        insForm, err := db.Prepare("INSERT INTO Employee(name, email, cell) VALUES(?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, email, cell)
        log.Println("UPDATE: Name: " + name + " | Email: " + email + " | Celular: " + cell)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
		email := r.FormValue("email")
        cell := r.FormValue("celular")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Employee SET name=?, email=?, cell=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, email, cell, id)
        log.Println("UPDATE: Name: " + name + " | Email: " + email + " | Celular: " + cell)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}