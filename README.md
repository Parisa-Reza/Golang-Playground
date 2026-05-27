# GOLANG PLAYGROUND
![GoLang Mascot](image-1.png)

## Go Language Characteristics

### General
- Developed by Google in 2009
- Open source, statically typed, compiled language
- Designed for simplicity and performance

### Key Characteristics

- **Fast** — compiles directly to machine code, very fast execution
- **Simple** — clean syntax, easy to read and write
- **Statically typed** — types are checked at compile time, not runtime
- **Strongly typed** — no implicit type conversion
- **Garbage collected** — memory is managed automatically
- **Compiled** — code compiles to a single binary, no runtime needed
- **Concurrent** — built-in support for concurrency via goroutines and channels
- **Cross-platform** — compile once, run on any OS
- **No classes** — uses structs and interfaces instead of traditional OOP
- **Unused imports/variables = error** — Go enforces clean code strictly

### What Go is Good For
- Backend web servers and APIs
- Command line tools
- Cloud and distributed systems
- DevOps and infrastructure tooling

## Concurrency vs Parallelism

### Concurrency
- Dealing with multiple tasks at once — not necessarily simultaneously
- Tasks take turns using the same resource
- About **structure** — how code is designed
- Example: one cook switching between chopping, stirring, and boiling

### Parallelism
- Actually running multiple tasks at the exact same time
- Requires multiple CPU cores
- About **execution** — how code actually runs
- Example: multiple cooks each doing their own task simultaneously

### Key Difference

| | Concurrency | Parallelism |
|---|---|---|
| Tasks run at same time | Not necessarily | Yes |
| Needs multiple cores | No | Yes |
| About | Structure | Execution |
| Goal | Handle more tasks | Go faster |

### In Go
- **Goroutines** enable concurrency — lightweight threads managed by Go
- **Channels** let goroutines communicate safely
- Go can achieve parallelism too — but concurrency is its core strength

### Famous Quote by Rob Pike (Go creator)
> Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.


## Starting a New Go Project Step by Step

### Step 1: Install Go
- Download and install Go from [golang.org](https://golang.org)
- Verify installation by checking the version in your terminal

### Step 2: Create a Project Folder
- Create a new folder with your project name
- Navigate into that folder in your terminal

### Step 3: Initialize the Module
- Run `go mod init <your-project-name>` inside the folder
- This creates a `go.mod` file — your project is now a Go module
- The project name you choose becomes the base import path

### Step 4: Create the Entry Point
- Create a file called `main.go` in the root of your project
- Declare `package main` at the top
- Add a `func main()` — this is where your program will start

### Step 5: Write Your First Code
- Write your logic inside `func main()`
- Use `fmt.Println()` to print output (import `"fmt"` at the top)

### Step 6: Run the Program
- Run `go run main.go` in your terminal to execute the program



---

### Quick Command Reference

| Action | Command |
|---|---|
| Initialize module | `go mod init <name>` |
| Run program | `go run main.go` |
| Build binary | `go build` |
| Add dependencies | `go get <package>` |
| Tidy dependencies | `go mod tidy` |










## Go: Modules, Packages & Imports

### Module
- The entire project
- Defined by a `go.mod` file at the project root
- Gives the project a name used as the base import path
- Every project has exactly one module

### Package
- A folder containing one or more Go files
- All files in the same folder must share the same package name
- Package name should match the folder name
- Used to group related functionality together

### `package main`
- A special package that produces a runnable program
- Must contain a `func main()` — where execution starts
- Every Go app has exactly one `package main`

## `package util`
- A regular reusable package
- Contains helper logic to be used by other packages
- Cannot run on its own

### Exported vs Unexported
- Uppercase name → exported → visible to other packages
- Lowercase name → unexported → only visible inside the same package

### Importing
- Used to bring another package's code into the current file
- Standard library packages are imported by their short name
- Local packages are imported using the full module path + folder name

### How It All Fits Together
- Module = the whole project
- Package = a folder of related code inside the module
- `package main` = the entry point
- `package util` = reusable helper logic
- Importing = the link between packages


## `fmt.Println` vs `fmt.Printf`

### Println
- Adds a new line automatically
- Adds spaces between multiple values
- No formatting needed
- Best for quick, simple output

### Printf
- No automatic new line — use `\n` manually
- Requires format verbs like `%s`, `%d`, `%v`
- Full control over output format
- Best for structured or mixed output

### Common Printf Verbs

| Verb | Use for |
|---|---|
| `%s` | string |
| `%d` | integer |
| `%f` | float |
| `%v` | any value |
| `%T` | type of value |
| `\n` | new line |

### Rule of Thumb
- Quick print → use `Println`
- Formatted print → use `Printf`

## `fmt.Scan` vs `bufio.Reader`

### fmt.Scan
- Reads **one word** at a time
- Stops at a space or newline
- Simple and short
- Good for single values like a number or one word

### bufio.Reader
- Reads the **entire line** including spaces
- Stops at a delimiter you choose (e.g. `\n`)
- Needs more setup — requires `os.Stdin`
- Good for full sentences or multi-word input

### Rule of Thumb

| Situation | Use |
|---|---|
| Single word or number | `fmt.Scan` |
| Full sentence with spaces | `bufio.Reader` |


## Unused Variables, Functions, Packages & Underscore in Go

### Unused Variable
- Go throws a **compile error** if a variable is declared but never used
- Every variable you declare must be used at least once

### Unused Package
- Go throws a **compile error** if a package is imported but never used
- Only import what you actually need

### Unused Function
- Go does **not** error on unused functions
- You can declare a function and never call it — no error

### Underscore `_`
In Go, functions often return two values — a result and an error.
If you don't care about the error (quick scripts, learning, etc.) you use _ to ignore it

* name → gets the value
* _ → error is ignored


## Array vs Slice in Go

### Array

- **Fixed size**, part of the type: `[3]int ≠ [4]int`
- **Value type** — assignment copies all elements
- Size must be a compile-time constant
- Rarely used directly in Go

---

### Slice

- **Dynamic size**, backed by an underlying array
- **Reference type** — assignment shares the same data
- Has `len` (elements in use) and `cap` (total capacity)
- Supports `append` to grow

### Creating a slice with `make`

`make([]T, len, cap)` allocates a new underlying array and returns a slice over it. Use it when you know the size upfront — it avoids repeated allocations as the slice grows.

- `make([]int, 5)` — len 5, cap 5, all zeros
- `make([]int, 0, 10)` — len 0, cap 10 (pre-allocated, append-ready)

---

### Key Differences

| | Array | Slice |
|---|---|---|
| Size | Fixed | Dynamic |
| Type includes size? | ✅ `[3]int` | ❌ `[]int` |
| Assignment | Copies | Shares reference |
| `append` | ❌ | ✅ |
| `make` | ❌ | ✅ |
| Nil zero value | ❌ | ✅ |

---

## Map in Go

A **map** is an unordered collection of key-value pairs. Keys must be comparable; values can be any type.

---

### Basics

- Declared as `map[KeyType]ValueType`
- Use `make(map[K]V)` or a literal to initialize before writing
- Use `make(map[K]V, hint)` with a size hint when the number of entries is known upfront

---

### Common Operations

| Operation | Notes |
|---|---|
| Set a key | Overwrites if key already exists |
| Get a value | Returns zero value if key is missing |
| Check existence | Two-value form: `v, ok := m[key]` — `ok` is false if missing |
| Delete a key | Built-in `delete(m, key)` — safe even if key doesn't exist |
| Get length | `len(m)` returns number of key-value pairs |
| Iterate | `for k, v := range m` — order is **not** guaranteed |

---
## Defer in Go

A **deferred** call is scheduled to run at the end of the surrounding function — no matter how it exits (normally, early return, or panic).

---

### How It Works

- Runs **after** the function returns, before it fully exits
- Multiple defers run in **LIFO order** (last in, first out)
- Arguments are evaluated **immediately** when defer is called, not when it runs

---

### Common Use Cases

| Use Case | Why |
|---|---|
| Close a file | Guarantees cleanup right next to where the file is opened |


---

## Go File Handling — Simple Notes

### What Each Function Does

| Function | What It Does | What It Returns |
|---|---|---|
| os.Create | Makes a new file. Wipes it clean if it already exists | File pointer, Error |
| io.WriteString | Writes text into the file | Bytes written (int), Error |
| os.Open | Opens an existing file for reading only | File pointer, Error |
| file.Read | Reads a small chunk of the file at a time into a buffer | Bytes read (int), Error |
| ioutil.ReadFile | Reads the entire file at once. Deprecated since Go 1.16 | Byte slice, Error |
| os.ReadFile | Same as ioutil.ReadFile but modern and recommended | Byte slice, Error |
| file.Close | Closes the file and frees up the resource | Error |

---

### Buffer vs No Buffer

**Using a Buffer (file.Read)**
- Reads the file piece by piece
- Uses less memory
- Good for large files

**Not Using a Buffer (ioutil.ReadFile / os.ReadFile)**
- Reads everything at once
- Simple and quick
- Not great for very large files

---
## Getting data from APIs

### http.Get in API Fetching
 
http.Get sends a GET request to a URL and returns the server response. The response contains a Body which is a stream — you need to read it with io.ReadAll and close it with defer resp.Body.Close() after you are done. To use it, you must import net/http — this package provides everything related to HTTP requests and responses in Go.

---

### ioutil.ReadAll in API Fetching
 
When you fetch data from an API, the response body is a stream — ioutil.ReadAll reads that entire stream and gives you the raw bytes at once. You then convert it to a string or parse it as JSON to use the data. Just like ReadFile, it is deprecated — use io.ReadAll instead.

---
## URL Breakdown in Go

### Example URL

https://example.com:8080/products/search?q=laptop&page=2

### Parts of a URL

| Part | Value | What It Means |
|---|---|---|
| Scheme | https | The protocol used to communicate |
| Host | example.com | The server you are talking to |
| Port | 8080 | The door on the server to connect through |
| Path | /products/search | The specific resource or page on the server |
| Query Parameters | q=laptop&page=2 | Extra data sent with the request (key=value pairs) |

### Query Parameters

Query parameters come after the ? sign and are separated by &. In this example there are two parameters — q which is the search keyword and page which is the page number.

### net/url Package

To work with URLs in Go — like parsing, building, or extracting query parameters — you need to import net/url. This package lets you break a URL into its parts and read or modify them easily.


## Marshal vs Unmarshal

### Marshal — Struct to JSON

You have data in your code and you want to send it to an API. Marshal converts your Go struct into JSON text so the API can understand it.

Your struct goes in, JSON text comes out.

---

### Unmarshal — JSON to Struct

You got a response from an API and it is raw JSON text. Unmarshal converts that JSON text into a Go struct so your code can work with it.

JSON text goes in, your struct comes out.

---

### Functions
 
| Function | Takes | Returns |
|---|---|---|
| json.Marshal | A Go struct or value | Byte slice (JSON text), Error |
| json.Unmarshal | Byte slice (JSON text) + a pointer to a struct | Error |
 
---

### One Line Summary

| | Direction | Think of it as |
|---|---|---|
| Marshal | Struct → JSON | Packing your data to send |
| Unmarshal | JSON → Struct | Unpacking data you received |

---

## Go Concurrency — Simple Notes

---

### Concurrency vs Parallelism

**Concurrency** means dealing with multiple tasks at the same time — but not necessarily running them at the exact same moment. Think of one cook switching between chopping and boiling.

**Parallelism** means actually running multiple tasks at the exact same moment — multiple cores doing work simultaneously. Think of two cooks each doing their own task at the same time.

| | Concurrency | Parallelism |
|---|---|---|
| Tasks | Switch between tasks | Run at the same time |
| Cores needed | One is enough | Needs multiple cores |
| Think of it as | Juggling | Two people juggling |

---

### Goroutine

A goroutine is a lightweight thread managed by Go. You use the go keyword before a function call to run it concurrently. It is extremely cheap — you can run thousands of them without issues.

- You just write go before any function call
- It runs in the background while the rest of your code continues
- If main finishes, all goroutines are killed immediately

---

### The Problem with Goroutines

Since goroutines run in the background, your main function can finish before they are done. You need a way to wait for them — that is where WaitGroup comes in.

---

### sync.WaitGroup

A WaitGroup makes your main function wait until all goroutines are finished. Think of it as a counter.

| Method | What It Does |
|---|---|
| Add(n) | Says "I am starting n goroutines, wait for them" |
| Done() | Says "one goroutine is finished" — call this inside each goroutine |
| Wait() | Blocks and waits until the counter reaches zero |

Always call defer wg.Done() at the start of your goroutine so it always runs even if something goes wrong.

---

### Channel

A channel is a pipe that lets goroutines talk to each other and pass data safely. One goroutine sends a value in, another receives it out.

- Use make(chan type) to create a channel
- Send a value with ch <- value
- Receive a value with value := <- ch
- Channels are blocking — the sender waits until someone receives and vice versa

| | What It Does |
|---|---|
| ch <- value | Send value into the channel |
| value := <-ch | Receive value from the channel |
| close(ch) | Close the channel when done sending |

---

## Buffered vs Unbuffered Channel

**Unbuffered** — sender blocks until receiver is ready. They must meet at the same time.

**Buffered** — has a capacity. Sender can drop values in without waiting as long as the buffer is not full.

---

### When to Use What

| Tool | Use When |
|---|---|
| Goroutine | You want to run something in the background |
| WaitGroup | You want to wait for multiple goroutines to finish |
| Channel | Goroutines need to pass data or signal each other |

---

### The Simple Flow

```
main starts
  → launches goroutines with go
  → WaitGroup tracks them
  → goroutines do their work
  → pass results through channels if needed
  → WaitGroup.Wait() holds main until all are done
main exits
```