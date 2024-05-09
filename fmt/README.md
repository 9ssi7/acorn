## Go fmt Package

The `fmt` package is used to format and print values. It provides functions like `Print`, `Printf`, and `Println` to print values to the standard output. The `fmt` package also provides functions like `Sprint`, `Sprintf`, and `Sprintln` to format values and return the resulting string.

| **Verb** | **Meaning** | **Go Code Example** |
|---|---|---|
| %c | Prints a character | `fmt.Printf("%c", 'A')` |
| %d | Prints an integer in base ten | `fmt.Printf("%d", 10)` |
| %f | Prints a floating-point number with default width and precision | `fmt.Printf("%f", 3.14)` |
| %t | Prints a boolean value | `fmt.Printf("%t", true)` |
| %s | Prints a string value | `fmt.Printf("%s", "Hello, World!")` |
| %v | Prints the value in the default format | `fmt.Printf("%v", 10)` |
| %b | Prints the binary representation | `fmt.Printf("%b", 10)` |
| %x | Prints the hexadecimal representation | `fmt.Printf("%x", 10)` |
| %e | Prints the scientific notation | `fmt.Printf("%e", 3.14)` |

The given table shows some of the common verbs used in the `fmt` package. The `Printf` function is used to format and print values to the standard output. The first argument to the `Printf` function is a format string that contains verbs and the second argument is the value to be printed.

The `Sprintf` function is used to format values and return the resulting string. It takes a format string and a value as arguments and returns the formatted string.

The `Sprint` function is similar to the `Sprintf` function, but it does not take a format string. It takes a value as an argument and returns the formatted string.

The `Print`, `Println`, and `Println` functions are used to print values to the standard output. The `Print` function prints the value without a newline, the `Println` function prints the value with a newline, and the `Println` function prints the value with a newline and spaces between the values.

The `fmt` package also provides functions like `Fprint`, `Fprintf`, and `Fprintln` to print values to a specific `io.Writer` like a file or a network connection.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Printf("%c\n", 'A')
    fmt.Printf("%d\n", 10)
    fmt.Printf("%f\n", 3.14)
    fmt.Printf("%t\n", true)
    fmt.Printf("%s\n", "Hello, World!")
    fmt.Printf("%v\n", 10)
    fmt.Printf("%b\n", 10)
    fmt.Printf("%x\n", 10)
}
```

Output:

```
A
10
3.140000
true
Hello, World!
10
1010
a
```

In the above example, we have used the `Printf` function to print values to the standard output. We have used different verbs like `%c`, `%d`, `%f`, `%t`, `%s`, `%v`, `%b`, and `%x` to format and print values.