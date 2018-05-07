package main

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

// function to be provided to the embedded lua script
func host_print(L *lua.LState) int {
	str := L.ToString(1) // get first (1) function argument and convert it to String
	fmt.Println(str)
	return 0 // Notify that we pushed 0 values to the stack (no return)
}

func main() {
	fmt.Println("Testing Go <--> Lua integration...")

	l := lua.NewState()
	defer l.Close()
	// Load lua scripts into the Lua State
	if err := l.DoFile("lua/hello.lua"); err != nil {
		panic(err)
	}
	// Register our go-provided host_print function into the Lua Global Context
	l.SetGlobal("host_print", l.NewFunction(host_print))

	// Preparing the invocation metadata for the the lua-provided "sayhello" function
	sayhello := lua.P{
		Fn:      l.GetGlobal("sayhello"), // name of Lua function
		NRet:    1,                       // number of returned values
		Protect: true,                    // return err or panic
	}
	// Equivalent to call sayhello('my friend') from Lua
	if err := l.CallByParam(sayhello, lua.LString("my friend")); err != nil {
		panic(err)
	}

	// Get the returned value from the stack and cast it to a lua.LNumber
	if size, ok := l.Get(-1).(lua.LNumber); ok {
		fmt.Println("the length of the argument is ", size)
	}

}
