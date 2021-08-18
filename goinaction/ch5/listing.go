package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.email, u.email)
}

type admain struct {
	user
	level string
}

func main() {
	ad := admain{
		user: user{
			name:  "mikasa",
			email: "mikasa@email.com",
		},
		level: "super",
	}
	ad.user.notify()
	ad.notify()

}
func sendNotification(n notifier) {
	n.notify()
}
