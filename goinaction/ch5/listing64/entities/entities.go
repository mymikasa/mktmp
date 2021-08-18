package entities

type user struct {
	Name  string
	Email string
}

type Admain struct {
	user
	Rights int
}
