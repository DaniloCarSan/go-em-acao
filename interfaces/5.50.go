package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	Name  string
	Email string
}

type admin struct {
	user
}

func (u *user) notify() {
	fmt.Printf("Send email for user: %s \n", u.Email)
}

func (a *admin) notify() {
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
