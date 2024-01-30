# Welcome to Go! 

Welcome to the Go Techbier!

February 13th, 2024

Presented by Zak Cook & Selim Kälin

---

# Agenda

- Go Basics
- Standard Types and Syntax
- Structs 
- Functions and Methods
- Interfaces
- References and Pointers
- Control Structures
- Imports
- Standard Library
- Go Management Tools

---

## Get Your Fingers Dirty

- You get the chance to get your fingers dirty with your first Go project
- After each theory block, we will give you time to mess around in some Go code  

- Instructions to get the code skeleton:
```bash
# Install Go: https://go.dev/doc/install

cd ~/Downloads

sudo rm -rf /usr/local/go && tar -C /usr/local -xzf <go-version.tar.gz>
export PATH=$PATH:/usr/local/go/bin

# check installation
go version

# get skeleton code
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

Syntax

```go
// Comments start with two backslashes like so
// Variable declaration follows snakeCase syntax
var x int
var isTrue, isFalse bool
var (
    unsignedInteger uint8
    someFloat       float64
    myFirstString   string 
)

// Variable definition
x = 5
var (
    isTrue  bool = true
    isFalse bool = false
)

// Short syntax, type is inferred 
hello := "World"
```

---

Syntax

```go
// Everything in go belongs to a package
package main

// Objects starting with lowercase letters are NOT exported to other packages
numberInMainPackage := 42

// Uppercase names are exported
AvailableInOtherPackages := "Interpackagenal string"

// Constants are declared like so
const Pi float64 = 3.1415926
```

---

## Structs 

Structs are a sequence of named elements, called fields:

```go
// Empty struct
struct {}

// More interesting struct
struct {
    FieldOne   int
    FieldTwo   float64 
    FieldThree *[]uint16 // Pointer to an uint16 slice
}

// Define a new type
type Consultant struct {
    Name      string
    Age       int
    Project   string
    ahvNumber ahvNumber  // unexported field of type ahvNumber not declared here
}
```

---

## Functions and Methods

To distinguish between functions and methods in Go, we have to look at the context in which they are defined:

- Functions
  - standalone procedure, not associated with any object, i.e. a struct 

- Methods:
  - like a function but contains a receiver, which specifies what type the method belongs to
  - receiver can be any type, but in most cases it is a struct or pointer to a struct

Syntax:
```go
// Exported function with no return value
func SayHello() {
    fmt.Println("Hello!")
}

// Unexported method with a return value
func (c Consultant) getAhvNumber() ahvNumber {
    return c.ahvNumber
}
```

---

## Interfaces 

Interfaces specify a list of methods. A type set defined by an interface is the type set that implements all of those methods.

> **IMPORTANT**
> In go, interfaces are implemented **implicitly**! There is no explicit declaration of intent, such as the keyword `implements`.

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

## Pointers and References

Pointers are declared using the `*<variable>` syntax. Similarly, to pass a reference to a variable, we use the syntax `&<variable>`.

```go
package main

import "fmt"

func incrementByValue(x int) {
    x = x + 1 
}

func incrementByReference(x *int) {
    *x = *x + 1
}

func main() {
    myValue := 5
    incrementByValue(a)
    fmt.Println(a)            // Output: 5

    incrementByReference(&a)
    fmt.Println(a)            // Output: 6
}
```

---

## Control Structures

Go offers the following control structures:
- if / else if / else
- switch / case
- for / range / break / continue
- (select, defer, panic, go to)

---

```go
// If / else if / else structure
if condition {
    doSomething()
} else if someOtherCondition {
    doSomethingElse()
} else {
    doNothing()
}

// Switch-case structure
switch switchValue {
case caseOneValue:
    // Code for case 1
case caseTwoValue:
    // Code for case 2
default:
    // Code for default case
}
```

---

```go
// For loop
for i := 0; i < 10; i++ {
    fmt.Sprintf("Current number: %d", i)
}

// To imitate a while loop
for condition {
    // Code to execute while condition is true
}

// Range structure 
for index, value := range someCollection {
    fmt.Sprintf("Value at index %d: %d", index, value)
}
```

---

## Import Statements

- As stated previously, everything in Go belongs to a package, declared by the keyword `package`
- Packages are imported using the `import` statement at the beginning of a file
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

---

## Standard Library

- Go features a powerful and extensive standard library
- It covers areas such as I/O operations, text and image processing, cryptography, network programming, etc.
- You can find an overview here: https://pkg.go.dev/std

---
## Go Management Tools

- Just like its package management, Go offers very capable management tools
  - `go fmt` for code formatting
  - `go mod`, `go get`, and `go install` for module and dependency management
  - `go test` for testing

---

## Useful Resources

- Go official documentation: https://go.dev/doc/
- Effective Go (must-read): https://go.dev/doc/effective_go
- awesome-go: https://github.com/avelino/awesome-go






