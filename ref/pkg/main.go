// must run go run ., not go run main.go
// $ go run main_0.go
//# command-line-arguments
//./main_0.go:5:2: undefined: doThings
//./main_0.go:7:2: undefined: DoTHings
//
// $ go run .
//do things one!
//Do things two!!
package main

func main()  {

	doThings()

	DoTHings()
}
