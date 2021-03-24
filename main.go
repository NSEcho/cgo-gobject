package main

import (
	"fmt"

	"github.com/lateralusd/cgo-gobject/api"
)

func main() {

	tobj := api.CreateNewObject()
	fmt.Printf("Current value is: %d\n", tobj.GetValue())
	tobj.SetValue(15)
	fmt.Printf("Now value is: %d\n", tobj.GetValue())

	//tobj.AddCallback(func() {
	//	fmt.Println("You know me now bro?")
	//})

	tobj.ConnectSignal()

	tobj.Emit()
	/*
		tobj := CreateNewObject()
		fmt.Printf("[*] GObject @%p\n", tobj)

		fmt.Printf("[*] Current value is: %d\n", tobj.GetValue())

		fmt.Println("[*] Changing value to 15")
		tobj.SetValue(15)
		fmt.Printf("[*] Value is now: %d\n", tobj.GetValue())

		init_callback(func() {
			fmt.Println("Fuck you bro")
		})

		tobj.ConnectCallback()

		tobj.Emit()*/
}

/*
var defaultFn func()

func init_callback(fn func()) {
	defaultFn = fn
}

func defaultCallback() {
	fmt.Println("I am default callback")
}

func customCallback() {
	defaultFn()
}
*/
