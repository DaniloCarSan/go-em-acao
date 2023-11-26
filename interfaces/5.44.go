package main

type notifier interface {
	notify()
}

type User struct {
	Name  string
	Email string
}

func (u *User) notify() {

}

type User1 struct {
	Name  string
	Email string
}

func (u User1) notify() {

}

func sendNotification(n notifier) {

}

func main() {

	u := User{"Danilo Santos", "danilocarsan@gmail.com"}
	u1 := User1{"Danilo Santos", "danilocarsan@gmail.com"}

	// erro pois a interface foi implementada com receptor do tipo ponteiro
	// logo só valores do tipo ponteitos são aceitos
	sendNotification(u)
	sendNotification(&u)

	// quando implementado com receptor do tipo valor
	// tanto  um ponteiro como um tipo valor pode ser passado
	sendNotification(u1)
	sendNotification(&u1)

}
