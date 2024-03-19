package main

// networking: http, go, arc, mutex, channels, tcp, actor rpc, p2p
// ...
// impl.go -> all traits and models impls
// traits.go -> interfaces
// models.go -> structs
// errors.go -> custom errors
// events

// https://stackoverflow.com/questions/34464146/the-idiomatic-way-to-implement-generators-yield-in-golang-for-recursive-functi
// https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1
// https://levelup.gitconnected.com/use-go-channels-as-promises-and-async-await-ee62d93078ec#:~:text=To%20declare%20an%20%E2%80%9Casync%E2%80%9D%20function,logic%20inside%20that%20anonymous%20function.
// https://gobyexample.com/goroutines

var fullname string = "mohammaderfan arefimoghaddam"

type Person struct {
	Name string
}

type ByteSeq = []byte
type Trait = interface {
	// interface methods and types go here
	// ...
	// ~[]byte
	getCode()
	sayHello()
}
type Trait2 interface {
	~[]byte
}

func main() {
	person := Person{Name: "erfan"}

	// trait Trait{}
	// struct Struct{}
	// impl Trait for Struct{}

	var _ = ByteSeq("")

	var _ Trait = person // Trait interface is now of type person or person is bouded to Trait interface

	person.getName()

	// user := gosted.User{Name: "Erfan", Age: 28}
	// println("built user %s", user)

	// function must be defer to recover the panic since it must be the
	// last function that is being executed which recovers the last panic
	defer func() {
		if r := recover(); r != nil {
			println("recover")
		}
	}() // building and calling at the same time

	Anothermypanic()

	mypanic()

}

func mypanic() {
	panic("i panicked")
}

func Anothermypanic() {
	panic("i another panicked")
}

// impl for the struct itself
func (p Person) getName() string {
	return p.Name
}

// trait impls
func (p Person) getCode() {

}

func (p Person) sayHello() {

}
