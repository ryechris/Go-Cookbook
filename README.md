# Go: Cookbook

A set of small programs that each performs a specific task, which demonstrates certain feature(s) of Go.

Table of Contents:

## Table of Contents
- [0. Requisites}(#1-decimal-truncator)
- [1. Decimal Truncator](#1-decimal-truncator)
    + [Windows and macOS](#windows-and-macos)
    + [Linux](#linux)
- [Quick Start](#quick-start)
- [Contributing](#contributing)



### 0. Requisites: Go and Git Pull
To test these programs, please ensure you have Go installed on your machine.

You can do that by navigating to your Terminal and running this command:
```
go version
git version
```

If they return without error, you have them installed and you can proceed with 
an empty folder, and executing this line in that empty folder:
```
git pull https://github.com/ryechris/Go-Cookbook.git
```

And, of course, you can run the files with this format: 
```
go run <filename>.go
```

### 1. Decimal Truncator
```
go run 01-truncator.go
```
This program prompts the user to type a floating number (for example, 3.1415).
It then removes the decimal point and everything to the right of it (e.g. input: 3.14; result: 3).

Though small and clean, it demonstrates clearly the following concecpts:
- User Input
- Type Handling & Conversion 
