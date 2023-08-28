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
