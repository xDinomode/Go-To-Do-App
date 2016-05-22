package main

//Simple to do app in Golang (not finished)

//import postgres driver github.com/lib/pq 
//gorilla/mux handles routing
import (
  "database/sql"
  _ "github.com/lib/pq"
  "net/http"
  "github.com/gorilla/mux"
)

//open a connection to postgresql
  var db, err = sql.Open("postgres", "user=postgres password=root dbname=notes sslmode=disable")

//localhost:8080/post GET
func handleGet(res http.ResponseWriter, req *http.Request){

  //select * from posts table
  rows, err := db.Query("SELECT * FROM posts")
  checkErr(err)
  defer rows.Close()
 
  //loop over selected rows 
  for rows.Next(){

    //assign each column name returned from rows to a var 
    //this table only has two columns ID and POST
    var post string
    var id int
    //this method assigns the value returned to our var above
    err = rows.Scan(&id,&post)
    
    //write to the client our values
    res.Write([]byte(post))
  }
}

//localhost:8080/post POST
func handlePost(res http.ResponseWriter, req *http.Request){

  //grab the form value in the post request
  //<input type="text" name="msg">
  msg := req.FormValue("msg")

  //execute a query without returning rows
  db.Exec("INSERT INTO posts (post) VALUES($1)", msg)

  //this should redirect to / but havent' fixed it yet
  res.Write([]byte(msg))
}
func main(){

  defer db.Close()

  //mux.NewRouter() will be our router for handling get/post requests
  r := mux.NewRouter()
  r.HandleFunc("/post", handleGet).Methods("GET")
  r.HandleFunc("/post", handlePost).Methods("POST")
  r.Handle("/", http.FileServer(http.Dir("public")))
  
  http.ListenAndServe(":8080", r) 
}

//checks for errors but not the ideal way to do it
func checkErr(err error){
  if err != nil{
    panic(err)
  }
}
