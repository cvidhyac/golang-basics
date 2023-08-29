# golang-basics

## Operator types

- Comparison Operator: Greater than `>`, Lesser than `<`, Equal to `==`, GreaterThanOrEqualTo `>=`,
  LesserThanOrEqualTo `<=`
- Arithmetic Operator: Add `+`, Subtract `-`, Multiply `*`, Divide `/`, Modulo `%`
- Assignment Operator: Right Shift `>>`, Left Shift `<<`
- Bitwise Operator: Bitwise AND `&`, Bitwise OR `|`, Bitwise XOR `^`
- Logical Operator: AND `&&`, OR `||`, NOT `!`

## Types of functions

### Single return type function

This type of function could accept any number of arguments but return only one arg.

```go
package main

func nameOfTheFunction(arg1 int) int {
	return arg1
}
```

### Multiple Return Type function

Go allows us to return multiple args in the return type. If a function returns multiple args, all
the args returned by the function should either be accessed or ignored using blank identifier (`_`)
where the function is being called.

```go
package main

import "fmt"

func nameOfTheFunction(arg1 int, arg2 string) (int, string) {
	return arg1, arg2
}

func main() {
	sum, str := nameOfTheFunction(10, "Hello")
	fmt.Printf("The params are %v %v \n", sum, str)

	_, str = nameOfTheFunction(10, "World")
	fmt.Printf("Print the string ignoring the first arg %v \n", str)
}

```

### Functions with named return values

In this function we name the return values in the method signature. With named return types, we can
simply `return` without defining the names of the variables

```go
package main

func namedReturnTypes(arg1 int, arg2 string) (myNumber int, myString string) {
	myNumber = arg1
	myString = arg2
	return
}

func main() {
	namedReturnTypes(10, "Hello")
}

```

### Variadic function

This function type uses an ellipsis(`...`) to define a number of args of same type as an Optional.
With this operator, the number of args passed to the function can be anywhere between `0..N`

```go
package main

import "fmt"

func withVariadicArgs(values ...int) bool {

	for _, value := range values {
		fmt.Println(value)
	}

	if len(values) == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	isExecuted := withVariadicArgs(10, 20, 30)
	fmt.Printf("The function executed with 3 variadic args %v", isExecuted)

	isExecuted = withVariadicArgs()
	fmt.Printf("The function executed with 0 variadig args %v", isExecuted)
}

```

### Anonymous functions

An anonymous function does not have a name. It can be assigned to variable and invoked as needed.

```go
package main

import "fmt"

func main() {
	x := func(a int, b int) int {
		return a + b
	}
	fmt.Printf("Sum of values %v \n", x(20, 30))
}
```

### Higher order functions

Sending function as a parameter of another function can help simplify complex if-else statements.
For an example of higher order functions, check out : [this example](main/higherorderfunctions.go)

## Defer keyword

Defers execution of the method until the subsequent methods finish execution.

```go
package main

import "fmt"

func printOne() {
	fmt.Println("Print 1")
}

func printTwo() {
	fmt.Println("Print 2")
}

func printThree() {
	fmt.Println("Print 3")
}

func main() {
	printOne()
	defer printTwo()
	printThree()
}

```

For this function, the output would be :

```text
printOne
printThree
printTwo
```

### Address and De-referencing operator

```go
package main

import "fmt"

func main() {

	i := 10
	fmt.Printf("Address of the datatype and accessing pointer %T %v \n", &i, &i)
	fmt.Printf("Dereferencing the pointer, %T, %v", *(&i), *(&i))
}

```

### Pointers

Pointers allow data manipulation and passing by value thereby allowing efficient memory utilization

```go
package main

import "fmt"

func addressAndDeReference() {
	i := 10
	ptr := &i
	fmt.Printf("Print dataType and address of pointer %T %v \n", ptr, ptr)
	fmt.Println("De-reference to access value of pointer ", *ptr)
}

func modify(value *string) {
	*value = "newValue"
}

func main() {
	addressAndDeReference()
	value := "currentValue"
	fmt.Println(value)
	modify(&value)
	fmt.Println(value)
}
```

Output of this program would be :

```text
Print dataType and address of pointer *int, 0xcd3vc
De-reference to access value of pointer 10
currentValue
newValue
```

### Pointer Datastructures

Maps and slices are pointers by default. Any modification to these datastructures modifies the
parent.

### Struct

Structs are equivalent to Objects in Java. However, in Golang we can pass an object by both pass by
value or pass by reference.

For example of how to declare and use structs, see - [this example](main/structs.go)

### Method Set

The set of methods defined in a go program or the interfaces is called a "Method Set"

## Interfaces

- No keywords are required to implement an interface
- Similar to pre-java8 interface, no constants allowed only declaration of methodSets without func
  keyword
- Interface methods cannot have a function body
- An implementation has to implement all the methods in an interface
