# Toy test about Lua scripts embedding

**Goal**: to embed Lua code into different host languages (for example, Go and Java), and invoke the next function from the host language:

```lua
function sayhello(name)
    host_print('hello ' .. name)
    return string.len(name)
end
```

At the same time, Lua must invoke the `host_print` function, which is not provided as Lua code, but implemented in the Host language.

At the end, the Lua function returns a value that must be collected by the host language.

We have integrated the Lua code in Go and Java with the help of the next Lua implementations:

* Go: http://github.com/yuin/gopher-lua
* Java: http://luaj.org/luaj.html

## Running in Go

From the root directory of this project (where this `README.md` file is located):

```
golua $ go run go/main.go
Testing Go <--> Lua integration...
hello my friend
the length of the argument is  9
```

## Running in Java

Go to the `java` subfolder and execute `./gradlew run`:

```
golua $ cd java
java $ ./gradlew run

> Task :run
hello my friend
the length of the argument is 9

BUILD SUCCESSFUL in 1s
2 actionable tasks: 2 executed
```