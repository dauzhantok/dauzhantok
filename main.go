package main

import ("fmt"; "net/http"; "html/template")
type User struct {
  Name string
  Age uint16
  Money int16
  Avg_grades, Happiness float64
  Hobb []string
}

func (u User) getAllInfo() string{
  return fmt.Sprintf("Name: %s, age: %d, money: %d", u.Name, u.Age, u.Money)
}

func (u *User) setName(sname string){
  u.Name=sname
  }

func home_page(w http.ResponseWriter, r *http.Request){
  bob := User{"Dauzhan",20,21,0.3,0.9,[]string{"food", "dance"}}
  tmpl, _ := template.ParseFiles("templates/profile_page.html")
  tmpl.Execute(w, bob)
}

func cont_page(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Cont text")
}

func handleRequest(){
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", cont_page)
  http.ListenAndServe(":9001", nil)
}

func main(){
  //bob := User{name:"Dauzhan", age: 20, money: 100, avg_grades: 5.0, happiness:0.9}
  handleRequest()
}
