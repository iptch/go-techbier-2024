---
author: Zak Cook & Selim Kälin
---

# Welcome to Go!

Welcome to the Go Techbier!

---

# Agenda

- Go Basics
- Standard Types and Syntax
- Structs
- Functions and Pointers
- Error Handling
- Dealing with JSON

**You are up!**

- Loops and Slices
- Packages, Exports, Constants
- Arrays, Slices, Maps
- Methods
- Interfaces

**You are up!**

- Imports
- Go Management Tools

**You are up!**

---

## Get Your Fingers Dirty

- You get the chance to get your fingers dirty with your first Go project
- After each theory block, we will give you time to mess around in some Go code
- It makes sense if more experienced Gophers sit with less experienced ones
  - Collaboration is very welcome!

- Instructions to get the code skeleton:

```bash
# Install Go: https://go.dev/doc/install

cd ~/Downloads

sudo rm -rf /usr/local/go && tar -C /usr/local -xzf <go-version.tar.gz>
export PATH=$PATH:/usr/local/go/bin

# Check installation
go version

# Get skeleton code
git clone https://github.com/iptch/go-techbier-2024.git
```

---

## Go Basics

About Go ...

- created in 2009 by R. Griesemer, R. Pike, and K. Thompson at Google
- statically typed and compiled, including to standalone binaries
- features memory safety, garbage collection, structural typing
- built for simplicity and efficiency, i.e. no classes or inheritance
- built-in support for concurrency through `goroutines` and `channels`
- powerful standard library
- great and supportive tooling, e.g. `go test`
- backbone of cloud technology like Kubernetes

---

## Standard Types and Syntax

- the full language specification can be found at https://go.dev/ref/spec
- types include:
  - boolean
  - numeric
  - string
  - array, slice, and map
  - struct
  - function and interface
  - pointer
  - channel

---

## Declaration and Definition Syntax Basics

```go
package main

import "fmt"

func main() {
    var x int  // Comments start with two slashes, like so
    var myBoolean bool = true  // Variable declaration follows snakeCase syntax
    var (  // Declaration blocks are delimited by parantheses
        unsignedInteger uint8
        someFloat       float64
        myFirstString   string
    )
    // Variables take default value when uninitialized
    fmt.Println(x, myBoolean, unsignedInteger, someFloat, myFirstString)

    x = 5  // Variable assignment, x has to be declared previously

    var (
        isTrue  bool = true
        isFalse bool = false
    )
    fmt.Println(isTrue, isFalse)

    hello := "World"  // Short syntax for declaration and assignment, type is inferred
    fmt.Println(hello)
}
```

---

## Pop Quiz

```go
package main

import "fmt"

func main() {
    x := "who"
    {
        x := "can guess"
        x = "this variable?"
        fmt.Println(x)
    }
    fmt.Println(x)
}

```

---

## Structs and Visibility

```go
package main

import "fmt"

// Define a new type
type Consultant struct {
    Name      string
    Age       int
    Project   string
    ahvNumber string // Unexported (private) field due to lowercase
}

func main() {
    // Initialize fields with order
    host1 := Consultant{"Zak Cook", 27, "BIT CBCD", "756.0001.0002.03"}
    fmt.Println(host1)

    // Initialize fields by name
    host2 := Consultant{
        Name: "Selim Kaelin",
        // Forgot how old selim was
        Project:   host1.Project,      // Access fields with dot syntax
        ahvNumber: "756.0001.0002.04", // This would not work from another package
    }
    fmt.Println(host2)

    var host3 Consultant // No initialization
    host3.Name = "Jakob"
    fmt.Println(host3)

    // host1.ahvNumber from another package would not work!
}

```

---

## Functions and Pointers


To distinguish between functions and methods in Go, we have to look at the
context in which they are defined:

- Functions: standalone procedure, not associated with any object, i.e. a struct

Pointers are defined using the `*` notation and referenced using `&`.

```go
package main

import "fmt"

func incrementByValue(x int) int {
    return x + 1
}

// Void return type
func incrementByReference(x *int) {
    *x = *x + 1
}

func main() {
    a := 5
    incrementByReference(&a)
    fmt.Println(a)
    fmt.Println(incrementByValue(a))
}
```

---

## Error Handling With if Statements

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    bytesWritten, err := createFile("/tmp/defer.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("wrote %d bytes.\n", bytesWritten)
}

func createFile(p string) (int, error) {
    f, err := os.Create(p)
    // This is how error handling is done in go
    if err != nil {
        return 0, err
    }
    // This is only executed after the function returns
    defer f.Close()

    return fmt.Fprintln(f, "what up")
}
```

---

## All Together Now: Parsing JSON

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Consultant struct {
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
    Project  string `json:"project"`
}

func main() {
    f, err := os.Open("slides/big_p.json")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    var consultant Consultant
    err = json.NewDecoder(f).Decode(&consultant)
    if err != nil {
        panic(err)
    }

    fmt.Println(consultant)
}
```

---

## Task 1

**Now you are up!**

Open our git repository and check out the branch `tasks/1`.

Look around the project and check out the file `pokeapi/api.go`.

You will find instructions in the code.

We will continue in about _20 minutes_. The next slide contains some details
about for loops and slices, which you need to task 1b.

---

## For Loops and Slices

```go
package main

import "fmt"

func main() {
    // Create a slice
    var numbers = make([]int, 0)

    // For loop
    for i := 0; i < 3; i++ {
        numbers = append(numbers, i)
    }

    // To imitate a while loop
    i := 3
    for i != 5 {
        numbers = append(numbers, i)
        i += 1
    }

    // Add many elements
    moreNumbers := []int{5, 6, 7}
    numbers = append(numbers, moreNumbers...)

    // Range structure
    for index, value := range numbers {
        fmt.Printf("Value at index %d: %d\n", index, value)
    }
}
```

---

## Packages, Exports, and Constants Syntax Basics

```go
// Everything in go belongs to a package
package main
// Java: package ch.ipt.ch;

// Lowercase letter objects are NOT exported to other packages
var numberInMainPackage = 42
// Java: private static int numberInMainPackage = 42;

// Uppercase names are exported
var ExportedString = "Interpackagenal string"
// Java: public static String publicString = "You get the point";

// Constants are declared like so
const Pi float64 = 3.1415926
// Java: public static final float PI = 3.1415926

func main() {
    // empty
}
```

---

## Methods and Their Syntax

To distinguish between functions and methods in Go, we have to look at the
context in which they are defined:

- Methods:
  - like a function but contains a receiver, which specifies what type the
    method belongs to
  - receiver can be any type, but in most cases it is a struct or pointer to a
    struct

```go
// Unexported method with a return value, c Consultant is the receiver object
func (c Consultant) getAhvNumber() ahvNumber {
    return c.ahvNumber
}
```

---

## Interfaces

Interfaces specify a list of methods. A type set defined by an interface is the
type set that implements all of those methods.

> **IMPORTANT** In go, interfaces are implemented **implicitly**! There is no
> explicit declaration of intent, such as the keyword `implements`.

Syntax:

```go
// If it quacks like a duck it is a duck
type Duck interface {
    Quack()
}

type Goose struct {}

// Oops, I guess a goose is a duck
func (g Goose) Quack() {
    fmt.Println("Quack!")
}
```

---

## Task 2

**Now you are up!**

Open our git repository and check out the branch `tasks/2`.

Look around the project and check out the file `ui/item.go`.

You will find instructions in the code.

We will continue in about _20 minutes_.

---

## Import Statements

- As stated previously, everything in Go belongs to a package, declared by the
  keyword `package`
- Packages are imported using the `import` statement at the beginning of a file
- Imports apply to the entire package, all exported identifiers will become
  available
- Package management is awesome! Look at the following example:

```go
package main

// Let's import multiple packages at once
import (
    "fmt"                                   // Standard library
    "math"                                  // Standard library
    http "net/http"                         // Create an alias called http
    "github.com/charmbracelet/bubbles/list" // External package we will need
)
```

```java
// Java
import java.util.*;
import java.util.ArrayList;
```

---

## Maps and "comma ok" notation

```go
package main

import (
    "fmt"
)

func main() {

    m1 := make(map[string]int)

    m1["k1"] = 7
    m1["k2"] = 13

    m2 := map[string]int{
        "k1": 7,
    }

    for k2, v2 := range m2 {
        // can be written without the ok, will panic on failure
        v1, ok := m1[k2]
        if ok && v1 == v2 {
            fmt.Printf("%s is present and equal in both maps\n", k2)
        }
    }
}
```

---

## Type assertions

```go
package main

import (
    "fmt"
)

func main() {
    var canBeAnything interface{}
    canBeAnything = "a string"

    // type **assertion**. we are telling Go "this is definitely a string, convert to one"
    ofTypeString := canBeAnything.(string)
    fmt.Println(ofTypeString)

    // comma ok notation possible
    _, ok := canBeAnything.(int)
    if !ok {
        fmt.Println("wasn't an int")
    }
}
```

---

## Go Management Tools

- Just like its package management, Go offers very capable management tools
  - `go fmt` for code formatting
  - `go mod`, `go get`, and `go install` for module and dependency management
  - `go test` for testing

---

## Task 3

**Now you are up!**

Open our git repository and check out the branch `tasks/3`.

Look around the project and check out the file `pokeapi/api.go`.

You will find instructions in the code. We have added a test in `pokeapi/api_test.go`,
check it out and run it with `go test ./...`.

We will continue in about _20 minutes_.

---

## Bonus Tasks

Wow! You have come a long ways.

If you are still wanting to play around more, have a look at the branch
`tasks/bonus`.

You will want to start in `pokeapi/api.go`.

---

## Useful Resources

- Go official documentation: https://go.dev/doc/
- Effective Go (must-read): https://go.dev/doc/effective_go
- awesome-go: https://github.com/avelino/awesome-go
