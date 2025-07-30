## 1️⃣ Setup & Tooling

- [x] Install Go (via `gvm`, `goenv`, or direct installer)
- [x] Understand `GOROOT`, `GOPATH`, and modules
- [x] Configure editor (Neovim + `gopls`, formatter, linter)
- [x] Learn basic Go CLI commands:
  - [x] `go version`, `go env`
  - [x] `go mod init`, `go run`, `go build`, `go test`

---

## 2️⃣ Language Fundamentals

- [ ] Packages & `package main`
- [ ] Variables (`var`, `:=`, `const`) and zero values
- [ ] Functions (multiple return values, named returns)
- [ ] Control flow: `if`, `for` (no `while`), `switch`
- [ ] Slices, arrays, maps (differences & iteration)
- [ ] Structs (composition > inheritance)
- [ ] Error handling (`if err != nil`)
- [ ] Go's explicit philosophy ("no magic")

### 🏆 Challenge Stage (Step 2)

- [ ] Reverse string (rune-safe)
- [ ] WordCount (maps + string parsing)
- [ ] FizzBuzz (control flow)
- [ ] Unique & Sorted copy (slices + maps)
- [ ] Vec2 struct operations (Add, Scale, Len)
- [ ] ParseAge with error handling

---

## 3️⃣ Data Structures & Methods

- [ ] Deep dive: slices, maps, pointers
- [ ] Methods vs functions
- [ ] Interfaces (duck typing) & embedding
- [ ] Custom types and receivers
- [ ] Idiomatic patterns (Option structs, Reader/Writer, decoupling via interfaces)

### 🏆 Challenge Stage (Step 3)

- [ ] Implement Stack (push, pop, peek)
- [ ] Implement Queue (enqueue, dequeue)
- [ ] Implement a type with a custom `String()` method
- [ ] Build a reusable `Set` type with `Add`, `Remove`, `Has`

---

## 4️⃣ Error Handling & Testing

- [ ] Idiomatic errors (`error` interface, `fmt.Errorf`)
- [ ] Sentinel errors vs `errors.Is` / `errors.As`
- [ ] Testing in Go: `testing` pkg, table-driven tests
- [ ] Temporary dirs/files in tests
- [ ] Coverage (`go test -cover`)

### 🏆 Challenge Stage (Step 4)

- [ ] Write tests for Step 2 challenges (table-driven)
- [ ] Add benchmarks for FizzBuzz and WordCount
- [ ] Create a mock interface and test against it

---

## 5️⃣ Concurrency

- [ ] Goroutines
- [ ] Channels (sync & async)
- [ ] `select` & timeouts
- [ ] `sync` pkg: WaitGroups, Mutexes
- [ ] Context cancellation (`context.Context`)

### 🏆 Challenge Stage (Step 5)

- [ ] Build a worker pool
- [ ] Write a concurrent `MapReduce`
- [ ] Implement a safe counter using Mutexes

---

## 6️⃣ Modules & Packages

- [ ] Multi-package project layout
- [ ] Go modules: `go mod tidy`, versioning, `replace`
- [ ] Import paths (internal, external)
- [ ] Dependency management best practices

### 🏆 Challenge Stage (Step 6)

- [ ] Split challenges into multiple packages
- [ ] Create a utility package and import it in `main`

---

## 7️⃣ Standard Library Essentials

- [ ] File I/O (`os`, `io`, `bufio`)
- [ ] HTTP servers & clients (`net/http`)
- [ ] JSON & YAML (`encoding/json`)
- [ ] Time handling (`time`)
- [ ] Logging (`log`), strings (`strings`), numbers (`strconv`)

### 🏆 Challenge Stage (Step 7)

- [ ] Write a JSON-to-CSV converter
- [ ] Build a tiny HTTP server serving static files
- [ ] Make a simple HTTP client with timeout support

---

## 8️⃣ Tooling & Best Practices

- [ ] Formatting (`goimports`, `gofumpt`)
- [ ] Linting (`golangci-lint`)
- [ ] Profiling & benchmarking (`pprof`)
- [ ] File embedding (`//go:embed`)
- [ ] Release binaries (`go build` cross-compilation)

### 🏆 Challenge Stage (Step 8)

- [ ] Add linters and fix all warnings
- [ ] Profile one of your challenges and optimize it

---

## 9️⃣ Advanced Topics

- [ ] Generics (Go 1.18+)
- [ ] Reflection (`reflect`)
- [ ] Unsafe & performance tuning
- [ ] Custom linters & code generation
- [ ] Private modules & vendoring

### 🏆 Challenge Stage (Step 9)

- [ ] Build a generic `Set[T]` type using generics
- [ ] Explore a reflection-based deep copy utility

---

## 🔟 Real-World Projects

- [ ] Build a REST API (`net/http` or `chi`)
- [ ] Build a CLI tool (`cobra`)
- [ ] Connect to a database (`pgx` for Postgres)
- [ ] Dockerize and deploy (CI/CD)
- [ ] Logging & monitoring (metrics, tracing)

### 🏆 Challenge Stage (Step 10)

- [ ] Complete a full project from scratch and deploy it

---
