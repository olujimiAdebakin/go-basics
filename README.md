# Go Core Concepts: A Practical Guide 🐘

A comprehensive repository demonstrating the fundamental concepts of the Go programming language. This project serves as a hands-on guide, with each file dedicated to a specific feature, from basic data structures to advanced concurrency patterns. It's designed to be a clear and practical resource for both beginners and experienced developers looking to solidify their understanding of Go.

## ✨ Features

-   **Core Data Structures**: In-depth examples for Arrays, Slices, and Maps.
-   **Control Flow**: Clear demonstrations of Loops, Conditions, and Switch statements.
-   **Functions & Methods**: Covers everything from basic functions and variadic parameters to closures and methods on structs.
-   **Types & Interfaces**: Practical examples of Go's type system, including structs, pointers, and the power of interfaces for polymorphism.
-   **Concurrency**: A deep dive into Goroutines and Channels, showcasing worker pools, select statements, and common concurrency patterns.
-   **Memory Management**: Explanations of `defer`, pointers, and the immutability of certain types.

## 🛠️ Technologies Used

| Technology | Description                            |
| :--------- | :------------------------------------- |
| **Go**     | The core programming language used.    |
| **Go CLI** | Used for running and managing the code. |

## 🚀 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Installation

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/olujimiAdebakin/go-basics.git
    ```
2.  **Navigate to the Directory**:
    ```bash
    cd go-basics
    ```
3.  **Ensure you have Go installed**:
    You can download it from the [official Go website](https://go.dev/doc/install).

## 💡 Usage

Each `.go` file in this repository is a self-contained program that can be run individually to explore a specific concept.

To run any of the files, use the `go run` command followed by the filename:

```bash
go run <filename>.go
```

For example, to run the demonstration on slices:
```bash
go run slice.go
```

### Available Modules

Here is a guide to the concepts covered in each file:

| File              | Concept Demonstrated                                                                  |
| ----------------- | ------------------------------------------------------------------------------------- |
| `var.go`          | Declaration and usage of various data types (int, string, bool, etc.).                |
| `type.go`         | Type conversion and handling user input of different types.                           |
| `operator.go`     | Arithmetic, comparison, and logical operators.                                        |
| `conditions.go`   | `if`, `else if`, and `else` statements for conditional logic.                         |
| `switch.go`       | Basic, fallthrough, and type-switch statements.                                       |
| `loop.go`         | `for` loops, while-style loops, and iterating over collections.                       |
| `array.go`        | Fixed-size arrays, declaration, iteration, and multi-dimensional arrays.              |
| `slice.go`        | Dynamic arrays (slices), including creation, appending, slicing, and capacity.        |
| `maps.go`         | Key-value pairs, including creation, modification, deletion, and iteration.           |
| `range.go`        | Using the `range` keyword to iterate over slices, maps, and strings.                  |
| `function.go`     | Defining functions, parameters, multiple return values, and variadic functions.       |
| `struct.go`       | Creating custom data types using structs and defining methods on them.                |
| `interface.go`    | Defining and implementing interfaces to achieve polymorphism.                         |
| `pointer.go`      | Using pointers to reference and modify data in memory.                                |
| `deref.go`        | The concept of dereferencing a pointer to access its underlying value.                |
| `defer.go`        | Postponing function execution for cleanup tasks like closing files or connections.    |
| `closure.go`      | Functions that capture variables from their surrounding scope.                        |
| `channel.go`      | Communicating between goroutines using channels (buffered and unbuffered).            |
| `make.go`         | Using the `make` function to initialize slices, maps, and channels.                   |
| `immutable.go`    | Demonstrates immutable types like strings and numbers.                                |
| `mutable.go`      | Demonstrates mutable types like slices and maps.                                      |

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  **🍴 Fork the Project**
2.  **🌿 Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3.  **📝 Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4.  **🚀 Push to the Branch** (`git push origin feature/AmazingFeature`)
5.  **🎉 Open a Pull Request**

## 📄 License

This project is licensed under the MIT License.

## 👤 Author

**Olujimi Adebakin**


-   **Twitter**: [@olujimi_the_dev](https://x.com/olujimi_the_dev)

<br/>

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)