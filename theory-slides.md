# Welcome to Go! 

Welcome to the Go Techbier!

February 13th, 2024

Presented by Zak Cook & Selim Kälin

---

# Agenda

- Go Basics
- Standard Types and Syntax
- Structs and Interfaces
- References and Pointers
- Functions and Methods
- Control Structures
- Standard Library
- Imports
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

## Structs and Interfaces

Structs are a sequence of named elements, called fields

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
