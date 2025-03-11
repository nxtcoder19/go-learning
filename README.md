# Golang

Golang, often referred to as Go, is an open-source programming language 
developed by Google in 2007 and released to the public in 2009. It was created by 
Robert Griesemer, Rob Pike, and Ken Thompson to address performance, simplicity, 
and scalability issues present in other programming languages.

Type: Statically typed, compiled language.
Designed For: Concurrent, scalable, and highly efficient applications.
Syntax: Clean, concise, and similar to C, but with improved safety and performance features.

# Why Golang

Golang is widely used for modern software development because of its performance, 
simplicity, and ability to handle concurrent programming efficiently. 
Here are the key reasons to use Go:

- Compiled to Native Machine Code:
    Go is compiled to native machine code, which makes it fast and efficient. 
    This means that Go programs can run faster than equivalent C or C++ programs.

    When a program is compiled to native machine code, it means the source code is translated directly
    into the binary instructions that the CPU can execute without needing an intermediary 
    (like a virtual machine or interpreter).
    This approach ensures that the program runs directly on the hardware with 
    maximum efficiency, as the instructions are specifically optimized for the 
    target architecture (e.g., x86, ARM).

- Concurrent Programming:
    Go is designed to handle concurrent programming efficiently. 
    It provides built-in support for goroutines, channels, and synchronization primitives.

- Garbage Collection:
    Go uses a garbage collector to automatically manage memory allocation and deallocation. 
    This eliminates the need for manual memory management, making Go programs easier to write and maintain.

- Static Typing:
    Go is a statically typed language, which means that variables have a specific type and cannot be changed 
    during runtime. This helps catch errors early and improves code reliability.

- Concise Syntax:
    Go has a clean and concise syntax that is easy to read and write. 
    It uses a combination of keywords and symbols to express concepts, making it easy to learn and use.

- Standard Library:
    Go comes with a standard library that provides a wide range of functionality, including networking, 
    file I/O, and cryptography. This makes it easy to build robust and scalable applications.

- Cross-Platform Compatibility:
    Go is designed to be cross-platform compatible, which means that Go programs can run on different operating systems 
    and architectures without modification. This makes it easy to deploy and run Go applications on a wide range of devices.

# [Runes](https://exercism.org/tracks/go/concepts/runes)
#### [read in detail here](https://www.educative.io/answers/what-is-the-rune-type-in-golang)

    
The rune type in Go is an alias for int32.
Given this underlying int32 type, the rune type holds a signed 32-bit integer value.
However, unlike an int32 type,
the integer value stored in a rune type represents a single Unicode character.

In Go, the rune type represents a single Unicode code point.

rune in Go is a data type that stores codes that represent Unicode characters.
Unicode is actually the collection of all possible characters present in 
the whole world. In Unicode, each of these characters is assigned
a unique number called the Unicode code point. This code point is what
we store in a rune data type.

# Alias
  
Type aliasing refers to the technique of providing an alternate name for an
existing type. Type aliasing was introduced in Go version 1.9 and above. 
This helps to facilitate code refactor for large codebases.
    
# Unicode and Unicode Code Points

Unicode is a superset of ASCII that represents characters by 
assigning a unique number to every character. 
This unique number is called a Unicode code point. Unicode aims to represent
all the world's characters including various alphabets, numbers, symbols,
and even emoji as Unicode code points.

# [utf-8](https://hackthedeveloper.com/golang-utf8-package-text-encoding/)
    
UTF-8 is a variable-width character encoding that is used to encode 
every Unicode code point as 1, 2, 3, or 4 bytes.
Since a Unicode code point can be encoded as a maximum of 4 bytes,
the rune type needs to be able to hold up to 4 bytes of data.
That is why the rune type is an alias for int32 as an int32 
type is capable of holding up to 4 bytes of data.

Go source code files are encoded using UTF-8.

# Number Parsing

-   Syntax:
    `func ParseInt(s string, base int, bitSize int) (i int64, err error)`

-   parameter:

    s: String value to be converted in an integer number.

    base: The base value of the given value. It can range from 0, 2 to 36.

    bitSize: It specifies the integer type, such as, int(0), int8(8), int16(16), int32(32), and int64(64).

- base?

  In the context of number parsing, the "base" refers to the numerical base or radix that is used to interpret the characters in a string as a numeric value. Different bases represent different numbering systems. The most common bases are decimal (base 10), binary (base 2), octal (base 8), and hexadecimal (base 16).

When parsing a number from a string, you need to specify the base to correctly interpret the characters in the string as digits of the desired numeric base. For example:

Decimal (base 10) uses digits 0-9.
Binary (base 2) uses digits 0 and 1.
Octal (base 8) uses digits 0-7.
Hexadecimal (base 16) uses digits 0-9 and letters A-F (or a-f) to represent values 10-15.

In Go's strconv package, the base is specified using the second argument in functions like ParseInt and ParseUint. For instance:

go

i, err := strconv.ParseInt("101010", 2, 64) // Parses binary "101010" as base 2

In this example, the string "101010" is interpreted as a binary number (base 2) 
because the second argument is set to 2.

Similarly, if you want to parse a hexadecimal string:

go

h, err := strconv.ParseInt("1A3F", 16, 64) // Parses hexadecimal "1A3F" as base 16

In this example, the string "1A3F" is interpreted as a hexadecimal number 
(base 16) because the second argument is set to 16.

By specifying the correct base, you ensure that the characters in
the string are interpreted correctly to represent the desired numeric value.
  

# SHA256 hashes

SHA256 hashes are frequently used to compute short identities for binary
or text blobs. For example, TLS/SSL certificates use SHA256 to compute
a certificate’s signature. Here’s how to compute SHA256 hashes in Go.

# BASE64

you can perform Base64 encoding and decoding using the encoding/base64 package.
Base64 encoding is commonly used to encode binary data into a text-based format
that is safe for transport and storage, such as in emails or URLs. Here's how
you can use the encoding/base64 package to perform Base64 encoding and decoding.

# Lines filters

A line filter is a common type of program that reads input on stdin, 
processes it, and then prints some derived result to stdout.
grep and sed are common line filters.

echo "hello" > /tmp/data
cat /tmp/data | go run lines-filter.go

# command line arguments & flags

Command Line Flags (Options):
Command line flags, also known as options, are typically preceded by a hyphen 
(-) or a double hyphen (--). They provide additional information or modify 
the behavior of a command. Flags are often used to enable or disable certain
features,set configuration values, or control the output of a program. For example:

shell

$ ls -l      # Display detailed list of files
$ git commit -m "Initial commit"   # Commit with a message
$ node script.js --debug           # Run a Node.js script in debug mode

Command Line Arguments:
Command line arguments are the values provided to a command or program 
after the command itself and any flags. They are typically used to specify \
the input or data that a program should operate on. Arguments can be filenames,
paths, text strings, or any other data required by the program. For example:

shell

$ cat file.txt             # Display the contents of a file
$ python script.py arg1    # Run a Python script with an argument
$ node calculator.js 5 3  

# context

In Go, the context package provides a way to carry deadlines, 
cancellation signals, and other request-scoped values across 
API boundaries and between processes. The context package is 
particularly useful in scenarios such as handling concurrent or distributed 
systems, where you need to manage and propagate context information across 
various components.

The context package introduces the Context type, which represents
a context for carrying request-scoped values and cancellation signals.
A Context value can be used to control the behavior of operations, such
as timeouts, cancellations, and passing values between functions without
explicitly passing them as function parameters.

# describe below struct:

    type Manager struct {
    ManagerFirstName   string  `json: "manager_first_name"`
    ManagerLastName    string  `json: "manager_last_name"`
    ManagerEmployeeID  string  `json: "manager_employee_id"`
    ManagerSalary      float64 `json: "manager_salary,omitempty"`
    }

-   The Manager struct is a data structure in the Go programming language that is likely intended to represent information about a manager. This struct is designed to be compatible with JSON serialization and deserialization, as indicated by the use of struct tags. Here's a breakdown of the fields and their annotations:

-    ManagerFirstName: This field is of type string and is intended to store the first name of the manager. The struct tag json:"manager_first_name" specifies that this field should be encoded as "manager_first_name" in the JSON representation.

-    ManagerLastName: Similar to ManagerFirstName, this field is also of type string and is used to store the last name of the manager. The json:"manager_last_name" tag indicates that this field should be encoded as "manager_last_name" in the JSON representation.

-    ManagerEmployeeID: This field is of type string and is meant to hold the employee ID of the manager. The json:"manager_employee_id" tag ensures that this field is encoded as "manager_employee_id" in the JSON representation.

-    ManagerSalary: This field is of type float64 and appears to represent the salary of the manager. The json:"manager_salary,omitempty" tag has two parts:
        json:"manager_salary" specifies that this field should be encoded as "manager_salary" in the JSON representation.
        ,omitempty indicates that if the ManagerSalary field is zero or has a zero value (like 0.0 for a float64), it will be omitted from the JSON output. This can help reduce unnecessary data in the JSON representation.

# Go Routines
        
Go routines are lightweight threads of execution that allow concurrent execution of functions. 
They are useful for performing tasks concurrently, such as handling multiple requests or processing data in parallel.
To create a new goroutine, you use the go keyword followed by a function call.

Go routing, based on goroutines, is Golang’s approach to managing concurrency. 
It allows functions to run concurrently, making programs efficient, scalable, 
and capable of handling many tasks simultaneously without blocking the main thread.

Steps in Goroutine Execution
Goroutine Creation:
    
- When go is used, the Go runtime creates a new goroutine, registers it with the scheduler, and immediately returns control to the calling function.
        Goroutine Scheduling:
        
- The scheduler assigns the goroutine to an available OS thread.
        If all OS threads are busy, the scheduler queues the goroutine until a thread becomes available.
        Execution:
        
- The goroutine executes independently of other goroutines.
        The Go scheduler ensures that all active goroutines get time to execute.
        Termination:
        
- When the goroutine completes, the memory and resources it used are reclaimed.

# Channel

- A channel is a typed conduit through which you can send and receive values of a type.
- The channel type is defined as:

```go
type ChannelType chan ValueType
```

    Channels in Golang provide a way to communicate and synchronize between goroutines. 
    They act as a conduit to pass data from one goroutine to another, ensuring safe 
    and synchronized communication.

- Channels can be used for broadcasting or sending event notifications.
- Channels can be used to pass data between goroutines, allowing them to communicate 
    and synchronize their execution.
- Channels can be used to implement buffered or unbuffered communication between 
    goroutines.
- Channels can be used to implement message passing between goroutines.
- Timeouts can be used to control the behavior of goroutines that are blocked on 
    channel operations.

# Unbuffered Communication
An unbuffered channel has no capacity to store messages. Communication between the producer (sender) and consumer (receiver) 
happens synchronously, meaning:

The sender blocks until a receiver is ready to consume the message.
The receiver blocks until a sender is ready to send a message.

An unbuffered channel has no capacity to hold messages. 
Every send (ch <- value) operation will block until a corresponding receive (<-ch) 
operation is ready, regardless of whether the send is inside a loop.

# Buffered Communication
A buffered channel has a specified capacity, meaning it can store a limited number of messages. Communication between the sender and receiver happens asynchronously up to the channel's capacity.

The sender does not block until the channel is full.
The receiver does not block until the channel is empty.

A buffered channel has capacity to hold messages. In a loop, 
the send operation (ch <- value) will not block until the buffer is full. Similarly, 
the receive operation (<-ch) will not block until the buffer is empty.

# Variadic Functions

A variadic function in Go is a function that can accept a variable number of arguments.
Variadic functions provide flexibility by allowing you to pass zero, one, or multiple 
arguments of the same type to the function.

```
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))       // Output: 6
	fmt.Println(sum(4, 5, 6, 7, 8)) // Output: 30
	fmt.Println(sum())              // Output: 0 (no arguments passed)
}

```

# Generic Types

Generics, introduced in Go 1.18, allow you to write flexible and reusable 
functions and types that work with different data types while maintaining type safety. 
Generics eliminate the need for duplicating code for different types or using interface{} 
(which sacrifices type safety).

```cgo
package main

import "fmt"

// Generic function with type parameter T
func printValues[T any](values []T) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func main() {
	printValues([]int{1, 2, 3})       // Works with int
	printValues([]string{"a", "b"})  // Works with string
	printValues([]float64{1.1, 2.2}) // Works with float64
}

```

# Type Assertion

Type assertion in Go is a way to retrieve the underlying concrete value from an interface{} type.
It is used when working with interfaces to extract the actual type stored in the interface.

```bash
value, ok := interfaceValue.(TargetType)


package main

import "fmt"

func main() {
    var data interface{} = "Hello, Go!"

    str, ok := data.(string) // Type Assertion
    if ok {
        fmt.Println("Extracted String:", str)
    } else {
        fmt.Println("Type assertion failed")
    }
}

```

# Line Filter

A line filter in Go is a program that reads input line-by-line, processes it, 
and outputs the modified content. It is commonly used in command-line tools to filter, 
transform, or analyze text-based data.

✅ Example Use Cases:
- Removing empty lines from a file.
- Extracting lines containing a specific keyword.
- Converting text to uppercase/lowercase.
- Counting occurrences of a word in each line.

# Race Condition

A race condition in Golang occurs when two or more goroutines access the same shared
resource (such as a variable) concurrently, and at least one of them modifies it,
leading to unpredictable behavior.

# sync.WaitGroup

The sync.WaitGroup in Golang is used to wait for a group of goroutines to complete their execution.
It helps in synchronizing multiple goroutines by ensuring that the main goroutine waits until 
all spawned goroutines have finished their tasks.

# Memory Allocation

In Golang, memory allocation and utilization are managed by the runtime through 
a combination of stack allocation, heap allocation, and garbage collection. 

- stack allocation: 
    - The stack is used for function calls and local variables with a known size at compile time.
    - It follows a Last-In-First-Out (LIFO) approach, making it efficient for function execution.
    - Stack allocation is automatic and happens automatically when a function is called.Variables allocated on the stack are automatically deallocated when the function returns.

    ```bash
    func main() {
        c := a + b  // 'c' is stored in the stack
    return c

     }
  ```


- heap allocation:
    - The heap is used for dynamic memory allocation, where memory is allocated and deallocated dynamically at runtime.
    - If a variable escapes the function scope (e.g., returned from a function or assigned to a pointer), it gets allocated on the heap.


    ```b[PANCard.pdf](..%2F..%2F..%2F..%2FDownloads%2FPANCard.pdf)ash
    func newPerson(name string) *string {
         p := name  // Heap allocation because it escapes the function scope
    return &p

     }
    ```

- garbage collection:
    - The garbage collector is responsible for managing the memory allocated on the heap.
    - It periodically scans the heap and identifies objects that are no longer reachable.
    - When an unreachable object is found, it is marked as garbage and can be deallocated.
    - GC runs in the background and pauses the program briefly to reclaim memory.

# Status code

200	http.StatusOK	Success
201	http.StatusCreated	Resource created
400	http.StatusBadRequest	Bad request
401	http.StatusUnauthorized	Unauthorized
403	http.StatusForbidden	Forbidden
404	http.StatusNotFound	Not found
500	http.StatusInternalServerError	Internal server error

