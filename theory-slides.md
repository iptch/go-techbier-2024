# Welcome to Go! 

Welcome to the Go Techbier!

February 13th, 2024

Presented by Zak Cook & Selim Kälin

---

# Agenda

- Go Basics
- Standard Library
- Standard Types and Syntax
- Arrays, Slices, Maps
- Structs 
- References and Pointers
- Functions and Methods
- Control Structures
- Dealing with JSON

**You are up!**

- Interfaces

**You are up!**

- Imports
- Go Management Tools

**You are up!**


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

## Standard Library

- Go features a powerful and extensive standard library
- It covers areas such as I/O operations, text and image processing, cryptography, network programming, etc.
- You can find an overview here: https://pkg.go.dev/std

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
var x int  // Comments start with two slashes, like so
var isTrue, isFalse bool  // Variable declaration follows snakeCase syntax
var (  // Declaration blocks are delimited by parantheses
    unsignedInteger uint8
    someFloat       float64
    myFirstString   string 
)

x = 5  // Variable definition, x has to be declared previously
var (
    isTrue  bool = true
    isFalse bool = false
)

hello := "World"  // Short syntax, type is inferred 
```
```java
// Compare this to Java
int myInt = 10;
```

---

## Packages, Exports, and Constants Syntax Basics

```go
// Everything in go belongs to a package
package main  
// Java: package ch.ipt.ch;

// Lowercase letter objects are NOT exported to other packages
numberInMainPackage := 42  
// Java: private int numberInMainPackage = 42;  

// Uppercase names are exported
ExportedString := "Interpackagenal string"  
// Java: public String publicString = "You get the point";

// Constants are declared like so
const Pi float64 = 3.1415926  
// Java: static final float PI = 3.1415926
```

---

## Arrays, Slices, Maps

```go
// Arrays have a fixed size
var myFirstArray [10]int
myFirstArray[4] = 7

mySecondArray := [3]int{1, 2, 3}

// Slices are dynamic
var myFirstSlice = []float64
myFirstSlice = append(myFirstSlice, 2.9)

mySecondSlice := []float64{1.0, 2.0, 3.0}

// Maps are similar to hashes or dictionaries
var myFirstMap map[string]string
myFirstMap["one"] = "two"

mySecondMap := map[int]string{1: "one", 2: "two"}
```

---

## Structs 

Structs are a sequence of named elements, called fields (similar to classes in Java):

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

## Functions and Methods

To distinguish between functions and methods in Go, we have to look at the context in which they are defined:

- Functions
  - standalone procedure, not associated with any object, i.e. a struct 

- Methods:
  - like a function but contains a receiver, which specifies what type the method belongs to
  - receiver can be any type, but in most cases it is a struct or pointer to a struct

---

## Function Syntax

```go
// Exported function with no return value
func SayHello() {
    fmt.Println("Hello!")
}
```

```java
// Apparently in Java, everything must be a class ... 
public class Hello {
  void sayHello() {
    System.out.println("Hello!"); 
  }

  public static void main(String[] args) {
    sayHello();
  }
}
```

---

## Method Syntax

```go 
// Unexported method with a return value, c Consultant is the receiver object
func (c Consultant) getAhvNumber() ahvNumber {
    return c.ahvNumber
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

## Dealing with JSON

- Go is widely used in web and cloud technology, where formats such as JSON and YAML are omnipresent
- The standard library has some helpful tools in `encoding/json` 

```go
package main

import "encoding/json"

type Consultant struct {
    Name      string  `json:"name" yaml:"name"`
    Age       int     `json:"age" yaml:"age"`
    Project   string  `json:"project" yaml:"project"`
}

func main() {
    // Struct -> JSON
    writeConsultant := Consultant{Name: "Felix Grüne", Age: 42, Project: "Unknown"}
    jsonData, err := json.Marshal(writeConsultant)
    // Handle error

    // JSON -> Struct
    var readConsultant Consultant
    err = json.Unmarshal(jsonData, &readConsultant)
    // Handle error
}
```

---

## Task 1

**Now you are up!**

Open our git repository and check out the branch `tasks/1`.  
Look around the project and check out the file `tbd`.  
You will find instructions in the code.

---

## Interfaces 

Interfaces specify a list of methods. A type set defined by an interface is the type set that implements all of those methods.

> **IMPORTANT**
> In go, interfaces are implemented **implicitly**! 
> There is no explicit declaration of intent, such as the keyword `implements`.

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
Look around the project and check out the file `tbd`.  
You will find instructions in the code.

---

## Import Statements

- As stated previously, everything in Go belongs to a package, declared by the keyword `package`
- Packages are imported using the `import` statement at the beginning of a file
- Imports apply to the entire package, all exported identifiers will become available
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

## Go Management Tools

- Just like its package management, Go offers very capable management tools
  - `go fmt` for code formatting
  - `go mod`, `go get`, and `go install` for module and dependency management
  - `go test` for testing

---

## Task 3

**Now you are up!**

Open our git repository and check out the branch `tasks/3`.  
Look around the project and check out the file `tbd`.  
You will find instructions in the code.

---

## Bonus Tasks

Wow! You have come a long ways. 

If you are still wanting to play around more, have a look at the branch `tasks/bonus`.

---

## Useful Resources

- Go official documentation: https://go.dev/doc/
- Effective Go (must-read): https://go.dev/doc/effective_go
- awesome-go: https://github.com/avelino/awesome-go





