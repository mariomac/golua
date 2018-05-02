package main

import (
  "github.com/yuin/gopher-lua"
	"fmt"
)

func go_print(L *lua.LState) int {  //*
	str := L.ToString(1)          // get first (1) function argument and convert to int
	fmt.Println(str)
	return 0                 // Notify that we pushed 0 values to the stack
}

func main() {
	l := lua.NewState()
	defer l.Close()

	if err := l.DoFile("lua/hello.lua"); err != nil {
		panic(err)
	}

	l.SetGlobal("go_print", l.NewFunction(go_print))

	sayhello := lua.P{
		Fn:      l.GetGlobal("sayhello"), // name of Lua function
		NRet:    1,                     // number of returned values
		Protect: true,                  // return err or panic
	}
	if err := l.CallByParam(sayhello, lua.LString("world")) ; err != nil {
		panic(err)
	}

	// Get the returned value from the stack and cast it to a lua.LString
	if size, ok := l.Get(-1).(lua.LNumber); ok {
		fmt.Println("the returned value is", size)
	}

}

