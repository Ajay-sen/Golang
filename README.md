# Golang
# Go Programming Language Fundamentals 🚀

A quick reference guide for **Go (Golang)** basics, commands, build processes, and memory management.

---

## 🔹 Basics

### Execution Model
- Go is **compiled** (can feel like interpreted during development).
- Similar workflow to **Java**, but simpler.

### Native Execution
- Compiles directly to **native machine code**.
- ❌ No JVM or virtual machine required.

### Industry Usage
- Widely used in **web apps**, **cloud platforms**, and **distributed systems**.
- Actively used by **big tech companies** in production.

### Object-Oriented Support
- **Yes and No**
  - ❌ No classes or traditional objects.
  - ✅ Uses **structs** to group data.
  - ❌ No method/function overloading.

### Syntax Rules
- **Case-sensitive**
- Very **strict syntax** (compiler catches errors early).

---

## 🛠 Go Commands

- `go help`  
  Displays all available Go commands and documentation.

- `go env GOPATH`  
  Shows the Go workspace and package path.

---

## 📦 Building Executables

Go supports **easy cross-compilation**.

### Build for Current OS
```bash
go build
```

### Build for specific OS
```bash
GOOS="windows" go build
```

## 🧠 Memory Management
- Automatic Garbage Collection (GC)
- No manual memory allocation or deallocation needed.

### Allocation Methods

| Feature  | `new()`               | `make()`                |
| -------- | --------------------- | ----------------------- |
| Purpose  | Allocates memory only | Allocates + initializes |
| Result   | Returns pointer       | Returns value           |
| Used For | All types             | Slices, maps, channels  |
| Storage  | Zeroed                | Initialized             |