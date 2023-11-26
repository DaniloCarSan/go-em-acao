package main

import "fmt"

type notifier interface {
	notify()
}

type User struct {
	Name  string
	Email string
}

func (u *User) notify() {
	fmt.Printf("Send email for user: %s \n", u.Email)
}

type Admin struct {
	Name  string
	Email string
}

func (a *Admin) notify() {
	fmt.Printf("Send email for user: %s \n", a.Email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {

	u := User{"User Name", "user@email.com"}
	a := Admin{Name: "Admin Name", Email: "admin@email.com"}

	sendNotification(&u)
	sendNotification(&a)

}
