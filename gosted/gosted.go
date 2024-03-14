package gosted

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// -0--0--0--0--0--0--0--0--0--0-
/* a survey into the Rusted Go */
// -0--0--0--0--0--0--0--0--0--0-

// in Rust since we don't have gc values are moved by default when we pass them to new scopes
// in Go since we have gc values will be copied when we pass them by value
// in Rust we can either pass by reference or clone of the type (clone make high ram overhead but Rust handles the heap by moving)
// in Go we can pass the type by pointer or reference, the default is mutable pointer and by dereferencing the pointer in other methods it affetcs the underlying data
// in Rust there are two kinds of pointers, immutable and mutable ones, can't have both of them at the same time and only one mutable pointer in each scope is allowed also we should pass &mut Type to methods to mutate the underlying data by its pointer
// in Go there is no restriction on having multiple pointers of a type in a same scope
// in Rust all the pointers of a type will get updated and validated after moving the type into a new scope and updaing its ownership address by the compielr but can't use none of them
// in Rust a mutable pointer type is of type &mut Type and can be defined using &mut like: let mut name = &mut String::from(""); its value can be taken by dereferencing it using *
// in Go a pointer is mutable by default and is of type *Type and can be declared using var pointer *Type;
// in Go a pointer can be initialized using & like: pointer = &Type; and its value can be taken by dereferencing it using *
// in Rust don't be afraid of allocating data on the heap cause Rust will clean the heap and drop types by itself
// in Rust tokio::spawn() is like goroutines both of them spawn a task inside a threadpool in the background, we can communicate between inside and outside of goroutines and tokio spawn threadpool using channels
// in Rust traits are dynamic sized types and must be on the heap behind a pointer, they can be used as method param like: para: impl Trait, return type like: -> impl Trait, bound generic to them like: F: FnMut() -> () or box them like: Box<dyn Trait> in case we don't know the implementor
// in Go interfaces are types and can be used directly in every where
// in Rust unlike Go when an interface or trait is defined for a type it supports both pointer and the type itself
// in Rust structs can be object with methods using impl Struct{} or extend their behaviour using traits in both we can have self.nethod() or Struct::method()
// in Go struct can be object but their methods must be implemented using interface{} or anonym func with absence of interface{}
// in Rust passing heap data into new scopes will move and drop their ownership out of the ram and transfer the value into a new ownership in that scope, we can pass by ref or clone them
// in Go everything is passed by value which makes a copy of data causes the allocated section in the heap get larger which reduces the speed of gc cacher
// in Rust once a fuction gets executed all its internal variables gets dropped out of the ram that's why we can't return a pointer to them from the function since Rust don't allow to have dnagled pointer
// in Rust passing a heap data into a function moves its ownership into the function scope and that's why we don't have access the type after moving

// there is no concept of immutable and mutable pointers and their restrictions in Go
// so every pointer type would be defined like `var p *User` and initialize like p = &User{Name: "erfan"}
// note that by default in Go every pointer is mutable and passing them into new scopes and functions
// first borrows the data second mutating them mutates the underlying data outside of the function

// generics with interface{} and go generics

type Function[T any, G any] func(*T) *G

type PlayingCard struct {
	Suit string
	Rank string
}

type Embed struct {
	Age int
}

type Player struct {
	Name string
	base Embed
}

type CustomError struct {
	Code int
	Msg  string
}

// / -------------------------
// / interfaces
type Deck struct {
	// slice of interfaces or traits objects the type of each card is interface{} can by any type of struct
	// this interface{} in this scope can have any methods like AddCard() and RandomCard() that we've added
	// for each deck, since each card is an interface or trait type we can define methods for that and use
	// them any where. we can't just define it as []*PlayingCard, though. we define it as []interface{} so
	// it can hold any type of card we may create in the future
	cards []interface{} // any type it can be
}

// camel cases are public and small cases are private
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []User `json:"friends"`
	Extra   []byte `json:"extrafield"`
	isAdmin bool   `json:"is_admin"`
}

type PointertoInt *int
type Num int
type PointertoNum *Num

type Number = uint8
type Fn = func(int) (Num int)

type String struct {
	Content string
}

func log() {

	fmt.Println("hello")
	var _ []byte = []byte("wildonion")
}

func main() {

	var u = User{
		Name:    "wildonion",
		Friends: []User{User{Name: "f1", Age: 20}}, // there is a cycle, in Rust this should be either pinned or behind a Rc, Arc, RefCell or Box
		Age:     28,
		Extra:   []byte("wildonion"),
		isAdmin: true,
	}

	// -----------------------------------
	// -------------- serding
	jsonDataBytes, _ := json.Marshal(&u)
	println(jsonDataBytes) // utf8 bytes
	println(string(jsonDataBytes))
	var new_user = new(User)
	var decoded_json = json.Unmarshal(jsonDataBytes, &new_user) // passing the pointer of new_user to fill it with utf8 bytes
	println(decoded_json)
	// ----------------------------------------
	/* Rust serding example using serde_json
	// ----------------------------------------
	// #[derive(Serialize, Deserialize)]
	// struct User{}
	// let user = User{};
	// let encoded = serde_json::to_vec(&user).unwrap();
	// let strigified = serde_json::to_string_pretty(&user).unwrap();
	// let decoded = serde_json::from_slice::<User>(&encoded).unwrap();
	// -----------------------------------
	*/

	/////////////// ----------------------------------------
	/////////////// pointers
	/////////////// ----------------------------------------
	// unlike Rust in Go we define pointer types using * but use it with & and also deref it with * for mutating it
	// in Rust we define pointer using & also the type of pointer is & and deref it using * for mutating it
	// pu is a pointer of type User, to create it we simply put the struct instance behind &
	var pu *User // in rust pu is defined using &User{} and shows the value of user not the pointer
	var num *int
	println("pointer to user:", pu) // 0x since we didn't initialize the User
	println("pointer to num:", num) // 0x since we didn't initialize the User
	pu = &u                         // pu contrains the address of u instance
	println("pu pointer", pu)

	// ---------------------------------------------
	// --------- compex map and slice of User objs
	// make is like Box in Rust it stores on slices, channel and map on the heap
	var user_map = make(map[string][]*User) // mapping between user names and slice of user objs
	user_map["onion"] = append(user_map["onion"], &User{Name: "onion", Age: 1})
	user_map["erfan"] = append(user_map["erfan"], &User{Name: "onion", Age: 1})
	println("user map", user_map)
	for key, val := range user_map {
		println("key, val", key, val)
		new_instance := User{Name: "new", Age: 1}
		user_instance := val[0]
		*user_instance = new_instance // * used for dereferencing to get the value to update the pointer accordingly the underlying data
	}
	println("user map", user_map)

	floatmap := map[string][]float32{
		"wildonion": []float32{12.34, 12.45},
		"erfan":     []float32{1.4, 5.6},
	}
	println(floatmap)

	// ---------------------------------------------

	type List[T any] struct {
		head *List[T] // self-ref struct, in rust this is cycle we should break it using Box, Pin, Rc, Arc
		tail *List[T]
	}
	_ = List[string]{} // we can init like this without naming

	userInstance := User{Name: "", Age: 0}
	println("&userInstance pointer >>> ", &userInstance)
	update_user(&userInstance)
	println("name after mutating >>> ", userInstance.Name) // won't get changed!!
	println("&userInstance pointer >>> ", &userInstance)

	update_user1(&userInstance)
	println("name after mutating >>> ", userInstance.Name) // will get changed!!

	update_user1(&User{Name: "wildonion", Age: 28, isAdmin: false})

	userinstance := User{Name: "nist", Age: 12}
	userpointer := &userinstance
	(*userpointer).isAdmin = false // dereferencing the pointer
	(*userpointer).Name = "hast"   // dereferencing the pointer
	println("user instance", userinstance.Name)
	println("user pointer", &userinstance)
	println("userpointer user instance value", userpointer)
	println("userpointer address itself", &userpointer)

	userinstance2 := User{Name: "nist", Age: 12}
	userpointer = &userinstance2
	(*userpointer).isAdmin = false // dereferencing the pointer
	(*userpointer).Name = "hast"   // dereferencing the pointer
	println("user instance", userinstance.Name)
	println("user pointer", &userinstance2)
	println("userpointer user instance value", userpointer)
	println("userpointer address itself", &userpointer)

	defer DeferMe1()

	res, err := RetRef([]int64{2, 3, 4, 5}...)
	println("", res, err)

	a := 100
	b := 102
	UpdateWithPointer(a, &b) // by default pointers in go are mutable
	println("a , b", a, b)

	// anonym function
	simpleFunc := func(num1 string, num2 string) (Rsult string, Response string) {
		Rsult = num1 + "<->" + num2
		Response = "json respone"
		return // knows should return Result and Response
	}

	println(simpleFunc("wild", "onion"))
	defer DeferMe0()

	jsons := ParseUsers([]User{u, *new_user}...)
	println("all users jsons : ", jsons)

	name := "0x4F"
	println([]byte(name))

	/////////////// ----------------------------------------
	/////////////// interfaces and traits
	/////////////// ----------------------------------------
	// extending struct behaviour can be done by implementing interface methods for its pointers or none pointers separately
	// ----------------------------------------
	/* 	  Rust trait and interface example
	// ----------------------------------------
		trait Interface{
			type This;
			fn getQuery(&mut self) -> &Self;
		}
		#[derive(Debug)]
		struct User{
			Name: String,
			Age: u8,
			IsAdmin: bool,
		}
		impl Interface for User{ // unlike Go Interface in Rust will be implemented for both pointer and none pointer instances
			type This = Self;
			fn getQuery(&mut self) -> &Self { // we can return ref since the pointer is valid as long as instance is valid
				if self.Name == "oniontori"{
					self.Name = String::from("wildonion");
				}
				self
			}
		}
		let mut user = User{Name: String::from("oniontori"), Age: 28, IsAdmin: true};
		let mut mutpuser = &mut user;
		mutpuser.getQuery(); // mutating the Name field of the user instance using the Interface trait and its mutable pointer
		// println!("user is changed {:?}", user); // the user is changed using its mutable pointer
		mutpuser.Name = String::from("change to anything else again");
		println!("user is changed {:?}", user);
		// println!("mutpuser is changed {:?}", mutpuser); // the mutpuser is changed also
	*/
	type QueryInterface = interface {
		GetQuery(queryId *string) string
	}
	NewUserInstace := User{Name: "oniontori", Age: 10, isAdmin: false}
	var upointer *User = &NewUserInstace
	getQuery := upointer.GetQuery() // the interface is implemented only for all User pointers
	println("the name after calling interface method is ", getQuery)
	println("NewUserInstace", NewUserInstace.Name)
	upointer.Name = "changed to anything else again"
	println("NewUserInstace", NewUserInstace.Name)
	defer DeferMe2() // the first defer is the last one it's a fifo poppig out process

	type Interface = interface{}
	type Interface1 interface {
		AddQuery(s *string)
	}

	player := Player{Name: "erfan"}
	pointer := &player
	pointer.GetName() // can only be called on pointers of Player

	// it's like haveing a trait object in Rust
	// trait Interface{}
	// struct Player{}
	// impl Interface for Player{}
	// let instance = Player{} // we can access all the traits methods on instance
	// let trait_object: Box<dyn Interface> = Box::new(instance)
	// interface can be any type so we're creating an interface object using Player
	// instance thus the interface methods must be implemented for the Player struct
	// it's like impl Trait for Struct in Rust which grants the all methods of traits access
	// ---------------------------
	// Rust traits
	// ---------------------------
	// trait Interface{}
	// struct Cat{}
	// struct Dog{}
	// impl Interface for Cat{}
	// impl Interface for Dog{}
	// fn get_param(param: impl Interface){}
	// let interfaces: Vec<&dyn Interface> = vec![&Cat{}, &Dog{}];

	/*
		1 -
			trait MyInterface{
				fn Describe(&self) -> String;
			}

		2 - impl MyInterface for Player{
				fn Describe(&self) -> String{
					self.Name
				}
			}
		3 - let player = Player::default();
		4 - let pname = player.Describe();
	*/
	type MyInterface = interface {
		Describe() string // adding trait methods, make sure you define it for either pointer of Player or its none pointer instance
	}
	var pinterface interface { // implementing trait for player (not pointer) instance
		Describe() string // adding trait methods, make sure you define it for either pointer of Player or its none pointer instance
	} = player
	// var emptyinterface interface{}
	println("interface >>", pinterface.Describe()) // method can be called on any Player instance
	// checking that the pinterface if of type Player or not, in Rust we can define a trait object
	// like: let myinterface_trait_object = Box::new(player); we are sure that the type of Box::new(player);
	// is a trait object of type MyInterface
	_, ok := pinterface.(Player)

	// ********************************************************
	// a trait casting in Rust also do the job
	// let trait_object = &user as &dyn Interface;
	// trait Interface{}
	// #[derive(Default)]
	// struct UserMe{}
	// impl Interface for UserMe{}
	// let casted = &UserMe::default() as &dyn Interface;
	// ********************************************************

	// a type can be an interface as long as it implements the interface or trait
	// methods so player or &player can be interface type if it implements the Describe method

	// for an empty interface{} it means _ (in the following) or the name of type which is _ right now
	// can be any type which is player now and it the future it can implements methods
	// inside the interface{}, to specify the exact type of interface{} we can use type assertion
	var _ interface{} = player

	names := []string{"wildonion", "erfan"} // slice of names
	// converting each name into interface
	// in Go, you cannot directly assign a pointer to one type to a pointer of another type, even if the underlying types are related.
	// since interface{} is a pointer to underlying type and also &name is a pointer to
	interfaces_names := make([]interface{}, len(names))
	for i, n := range names {
		// name := interface{}(n) // making n as interface
		interfaces_names[i] = n
	}

	// ***********************************************
	// ******** impl String{} to impl methods ********
	// note that interfaces can have any underlying type, later we can do type assertion to check the type
	// if we want to add methods for a type and call it by . on the instance
	// we should use interface{} in Rust it's like calling extra methods on
	// the object other than the ones that already supported by the struct
	// itself which can be done using traits
	var pointerString *String = &String{Content: "wildonion"}
	var stringinterface interface {
		getContent() *string // impl method for String struct only so the value can't be simple and builtin string
	} = pointerString // implementing the interface for the pointer instance of String it's like impl Interface for Struct{} in Rust

	get_cont := pointerString.getContent()
	println(get_cont)
	if cont, ok := stringinterface.(*String); ok { // if let Ok() = String{}
		println(cont.Content)
	}
	// another example
	// dead simple trait interface
	type AnotherInterface = interface {
		getContent() *string
		receievData([]*byte) *String // []*byte is &mut [u8] in Rust
	}
	var structpoiner *String
	structpoiner = &String{Content: "erfan"}
	stringinstance := String{Content: "harchi"}
	// When you call a method with a pointer receiver on a non-pointer instance, Go automatically takes the address of the non-pointer instance before calling the method.
	// Similarly, when you call a method with a pointer receiver on a pointer instance, Go automatically dereferences the pointer before calling the method.
	stringinstance.getContent()
	structpoiner.getContent()

	// in place initialization of interface
	// trait := interface {
	// }(structpoiner)

	var inter interface{} = "wildonion" // interfaces can have any underlying type in this case is string
	println(inter)                      // address of inter

	// functions that return interface{} values tend to be quite annoying,
	// and as a rule of thumb you can just remember that it’s typically better
	// to take in an interface{} value as a parameter than it is to return an interface{} value
	// fn return_interface<'valid>(interfaces: &'valid [&'valid dyn Interfaces]) -> &'valid dyn Interface{
	// 		trait Interface{}
	// 		let interfaces: &[&dyn Interface]; // since the size of trait is not known (is ?Sized) it must be behind &dyn
	// }
	// anonym function is like closure in Rust:
	// let anonym = (|pid: u8|{
	// 	String::from("")
	// })(0);
	_ = func(in []interface{}) (inter interface{}) {
		return in[0]
	}(interfaces_names)

	// slice of MyInterface types or those types that MyInterface methods
	// are implemented for them like Player struct it's like all structs
	// that implements MyInterface trait in Rust or in other words slice
	// of MyInterface trait objects, it's like dyn keyword which is used
	// to call methods of object safe traits on types dynamically using
	// its two pointers: one for vtable and the other to the struct instance
	// like what interface{} does exactly, interfaces and traits in Rust
	// are not exactly a type they're kinda dynamic types
	interfaces := []MyInterface{player}
	print(interfaces)
	// since all types implements interface{} type thus we can pass any type to
	// the method which has func Method(v interface{}) signature cause later we
	// can cast the type into the actual data that implements the interface method
	// valuing interface{} with any type be like implementing Interface trait for
	// every type in Rust
	// the following type assertion example:

	Num := 43
	// kinda like if let Some() = ...
	if code, err := f2(&Num); err != nil { // unpack the result then check if the error wasn't nill
		if ce, ok := err.(*CustomError); ok { // unpack the error trait assert into the CustomError pointer
			println(ce.Code)
			println(ce.Msg)
		}
		if code == 1 {
			println("no error")
		}
	}

	/////////////// ----------------------------------------
	/////////////// the deck and card
	/////////////// ----------------------------------------

	deck := NewPlayingCardDeck()
	// get a random card
	card := deck.RandomCard()
	fmt.Printf("selected card %s >>>> ", card)

	// asserting the interface{} card value into a *PlayingCard value cause cards
	// contains pointers to PlayingCard instances
	// asserts that the interface value i holds the concrete type
	// T and assigns the underlying T value to the variable t.
	playingCard, ok := card.(*PlayingCard) // get a reference to PlayingCard instance
	if !ok {
		println("false assertion card holds different type")
		os.Exit(1)
	}

	fmt.Printf("playing card %s >>> ", playingCard)

	/// ---------------
	// generics

	deck0 := NewPlayingCardDeck0()
	card0 := deck0.RandomCard0()
	fmt.Printf("selected card %s >>>> ", card0)

}

/////////////// ----------------------------------------
/////////////// methods
/////////////// ----------------------------------------

// seems we can't extend the methods of built in type string
func (s *String) getContent() *string {
	return &s.Content // return reference to the content if the type is already ref (*String) don't need to use &
}

// we would do normaly this in Rust:
// type Condition<T> = fn(T) -> bool;
// fn filter<T: IntoIterator<Item=T>>(slice: &[T], condition: Condition<T>) -> Vec<T>{ todo!() }
func filter[O any](arr []*O, condition func(*O) bool) (res []*O) {
	for _, v := range arr {
		if result := condition(v); result { // in Rust => if let Ok(res) = condition(&v){}
			res = append(res, v)
		}
	}
	return
}

func f2(arg *int) (int, error) { // it's like Box<dyn Error> the Error() method must be implemented for the CustomError
	if *arg == 43 { // arg is reference we should deref it to compare
		return 0, &CustomError{Code: 10, Msg: "arg is 43"}
	}

	return 1, nil
}

// Error() method is insie some builtin interface loaded in every app
func (err *CustomError) Error() string { // it's like implemeing Error trait for the instance
	return fmt.Sprint("%s : %s", err.Code, err.Msg)

}

// any interface method must be implemented in here so we can create an interface object with any instance of any type
// traits, implementing interface trait methods for Player struct and its pointer only
func (p *Player) GetName() string {
	return p.Name
}

func (p Player) Describe() string {
	return p.Name
}

// implementing the interface method for the User pointer only - anonym method
func (user *User) GetQuery() string {
	if user.Name == "oniontori" {
		user.Name = "wildonion"
	}
	return user.Name
}

func DeferMe0() {
	println("=========== defer me 0")
}

func DeferMe1() {
	println("=========== defer me 1")
}

func DeferMe2() {
	println("=========== defer me 2")
}

// Rust: type Func = fn(HttpRequest, String) -> String;
//
//	fn get_cls<F>(cls: F) where F: FnMut(String) -> String{
//			cls(String::from(""))
//	}
func Apply(n int, function func(int) int) int {
	return function(n)
}

func UpdateWithPointer(a int, b *int) {
	*b = 100 // * is dereferencing also we use it in defining pointer types instead of & which Rust uses it's a notation
	a = 10
}

var users_jsons [][]byte = make([][]byte, 0) // initializing using make or [][]string{}

func ParseUsers(users ...User) (JsonBytes [][]byte) {
	for _, user := range users {
		user_json_bytes, _ := json.Marshal(&user)
		users_jsons = append(users_jsons, user_json_bytes)
	}
	JsonBytes = users_jsons
	return // we have access it globaly outside of this method
}

func RetRef(values ...int64) (*int, error) { // Variadic function : ...int64 is []int64
	a := 3
	var value_arr = values
	for index, value := range value_arr {
		println("index and value", index, value)
	}
	return &a, nil
}

func update_user(user *User) (Response *User) {
	println("user instance address using user pointer", user)
	Response = user

	println("user pointer address itself", &user)

	// --------- changing address with new data but same data outside ------------

	// changing the address that is stored in user pointer to points to a new value address,
	// we're allocating new User instance which is located at another address into the user
	// pointer, doing so changes the address inside the user pointer which makes the pointer
	// points to a new User instance location when we return, that version of user inside method
	// is popped from the stack and the user pointer in main points to the very first location
	// of User instance which is unchanged, the reason is that we're putting a new User object
	// inside the user pointer inside the method which by executing the function the pointer
	// gets poped out of the stack and the outside instance remains the same, it's like binding
	// a new instance into an existing mutable pointer without dereferencing it like:
	// let mut user = User::default();
	// let mut mutpointer = &mut user;
	// let mut user1 = User::default();
	// mutpointer = &mut user1;
	// or
	// let name = String::from("");
	// let mut pname = &name;
	// let mut anotehr_pname = &String::from("new content");
	// println!("pname points to location of name : {:p}", pname);
	// println!("anotehr_pname points to location of name : {:p}", anotehr_pname);
	// pname = anotehr_pname;
	// println!("pname points to location of anotehr_pname : {:p}", pname);
	// println!("pname content : {:?}", pname);
	user = &User{Name: "erfan2", Age: 0, isAdmin: false} // outside of this method instance won't change

	println("user instance address using Response pointer", Response)
	println("Response pointer address itself", &Response)

	return
}

func update_user1(user *User) {

	// --------- changing data with same address updated data outside ------------

	// mutating user pointer mutates the actual User instance outside of the method
	// mutating the user pointer by dereferencing user and assigning a
	// new object to that location, overwriting the existing memory, when
	// we return the pointer in main is still pointing to the same location
	// but we have different data at that location in memory which made an
	// update to User instance too but remain the old address, it's like &mut _ in Rust
	// let mut user = User::default();
	// let mut mutpointer = &mut user;
	// let mut new_binding = User{name: String::from("wildonion")};
	// *mutpointer = &mut new_binding; // user has changed too
	// mutpointer.Name = String::from("only Name field"); // changing only name field
	println("user pointer address itself", &user)
	*user = User{Name: "erfan", Age: 0, isAdmin: false} // mutating the pointer with a new binding this will change the whole instance of underlying data
	println("user pointer address itself", &user)
	return
}

// ----------------------------------------
/* ---------- Rust version of pointers
// ----------------------------------------

fn but_the_point_is(){

    // type Ret = &'static str;
    // fn add(num1: Ret, num2: Ret) -> Ret where Ret: Send{
    //     for ch in num2.chars(){
    //         num1.to_string().push(ch);
    //     }
    //     let static_str = helpers::misc::string_to_static_str(num1.to_string());
    //     static_str
    // }

    // let addfunc: fn(&'static str, &'static str) -> &'static str = add;
    // let res = addfunc("wild", "onion");

    #[derive(Default, Debug)]
    struct User{
        name: String,
        age: u8,
    }

    let mut user = User::default(); // there is no null or zero pointer in rust thus the user must be initialized

    let mut mutpuser = &mut user; // mutating mutpuser mutates the user too
    println!("user address: {:p}", mutpuser); // contains user address
    println!("mutpuser address itself: {:p}", &mutpuser); // contains user address
    mut_user(mutpuser);

    fn mut_user(mut user: &mut User){ // passing by mutable pointer or ref to avoid moving

        // mutating the user pointer with new value which contains the user address
        // this makes an update to user instance too, can be viewable outside of the method
        println!("before mutating with pointer: {:#?}", user);
        user.name = "erfan".to_string();
        println!("after mutating with pointer: {:#?}", user);
        // or
        println!("before derefing: {:p}", user); // same as `contains user address`
        let mut binding = User{
            name: String::from("wildonion"),
            age: 0
        };
        // updating pointer which has the user instance value with a new binding by dereferencing pointer
        // note that we're not binding the new instance into the pointer completely cause by dereferencing
        // the underlying data will be changed
        *user = binding;
        println!("user after derefing: {:#?}", user);
        println!("user address after derefing: {:p}", user); // same as `contains user address`

    }

    // println!("out after mutating with pointer: {:#?}", user);
    let mut binding = User{
        name: String::from("wildonion"),
        age: 0
    };
    println!("mutpuser address itself: {:p}", &mutpuser); // contains user address
    println!("mutpuser contains address before binding: {:p}", mutpuser); // same as `contains user address`
    // binding a complete new instance to mutpuser, causes to point to new location
    mutpuser = &mut binding;
    // the address of mutpuser will be changed and points to new binding instance address
    println!("mutpuser contains address after binding: {:p}", mutpuser);
    println!("mutpuser address itself: {:p}", &mutpuser); // contains user address

    // we're getting a mutable pointer to an in place User instance
    // the in place instance however will be dropped after initialization
    // and its ownership transferred into mutpuser, Rust won't allow us to
    // do so cause a pointer remains after dropping the in place instance
    // which is and invalid pointer, we must use a binding to create a longer
    // lifetime of the User instance then borrow it mutably
    // mutpuser = &mut User{
    //     name: String::from(""),
    //     age: 0
    // }; // ERROR: temporary value is freed at the end of this statement

    // SOLUTION: using a `let` binding to create a longer lived value
    // let binding = User{
    //     name: String::from("wildonion"),
    //     age: 0
    // };
    // *mutpuser = binding;


    // let puser = &user;
    // println!("user address (puser): {:p} ", puser); // contains the address of user
    // let anotherpuser = puser;

    // println!("user address (anotherpointer): {:p} ", anotherpuser); // also contains the address of user

    // println!("pointer address: {:p} ", &puser); // the address of the puser pointer itself
    // println!("anotherpointer address: {:p} ", &anotherpuser); // the address of the puser pointer itself

    // user address (puser): 0x7ffea5896328
    // user address (anotherpointer): 0x7ffea5896328
    // pointer address: 0x7ffea5896348
    // anotherpointer address: 0x7ffea5896390

}
*/

func NewPlayingCard(suit string, rank string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: rank}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

func NewPlayingCardDeck() *Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			// AddCard appends an instance of type PlayingCard into the cards
			// slice of interfaces
			deck.AddCard(NewPlayingCard(suit, rank)) // cards contain PlayingCard instance our assertion will be ok later
			// deck.AddCard("new card") // another different types of adding using interface type
		}
	}

	return deck
}

// with interface we can define multiple methods each of
// with different signatutes that is callable on any instance
// of the type that has been passed to the interface method
// this way is kinda simulating the generics

// implements AddCard and RandomCard methods only for Deck pointer
// they can be called on a pointer of Deck to add different styles of card
func (d *Deck) AddCard(card interface{}) {
	d.cards = append(d.cards, card) // appending an interface
}

func (d *Deck) RandomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

/// -------------------------
/// generics

// ---------------------
// Rust approach
// ---------------------
/*
struct Deck<C>{
	pub cards: Vec<C>
}

struct Card{
	pub suits: Vec<String>,
	pub ranks: Vec<String>
}

struct Card1{
	pub info: String
}

let deck0 = Deck::<Card>{
	cards: vec![Card{ suits: vec![], ranks: vec![] }]
};
let deck1 = Deck::<Card1>{
	cards: vec![Card1{info: todo!()}]
};


impl<C> Deck<C>{
	pub fn NewPlayingCardDeck(&mut self) -> &mut[&Card]{
		todo!()
	}
}
*/
// any is an alias to interface{} type
type Deck0[C any] struct {
	cards []C
}

func (d *Deck0[C]) AddCard0(card C) {
	d.cards = append(d.cards, card)
}

func (d *Deck0[C]) RandomCard0() C {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx] // values inside cards are of type C
}

func NewPlayingCardDeck0() *Deck0[*PlayingCard] { // a pointer to deck which contains all PlayingCard pointers
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck0[*PlayingCard]{} // generic must be in [] right after the Type
	for _, suit := range suits {
		for _, rank := range ranks {
			// AddCard appends an instance of type PlayingCard into the cards
			// slice of interfaces
			deck.AddCard0(NewPlayingCard(suit, rank)) // cards contain PlayingCard instance our assertion will be ok later
			// deck.AddCard("new card") // another different types of adding using interface type
		}
	}

	return deck
}
