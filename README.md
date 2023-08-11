# Riyan's Go Cookbook

A set of small programs that each performs a specific task, which demonstrates certain feature(s) of Go.

Riyan is responsible for all the Go code, but CS Professor Ian Harris of UCI supplies the inspiration.

## Table of Contents
- [0. Requisites](#1-decimal-truncator)
- [1. Truncate Decimals](#1-decimal-truncator)
- [2. Find IBM](#2-ibm-finder)
- [3. Sort Numbers](#3-number-sorter)
- [4. Make JSON](#4-json-maker)
- [5. Read Data]()
    + [Windows and macOS](#windows-and-macos)
    + [Linux](#linux)

### 0. Requisites: Go and Git
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
go run 01-truncate-decimals.go
```
This program prompts the user to type a floating number (for example, 3.1415).
It then removes the decimal point and everything to the right of it (e.g. input: 3.14; result: 3).

Though small and clean, it demonstrates clearly the following concepts:
- User input
- Type handling & conversion


### 2. IBM Finder
```
go run 02-find-ibm.go
```
The above line prompts the user to enter a string. Any string.

The program then searches that string to see if the characters "i", "b", or "m" are present -- in that order.

For example, "icbm" returns Found, but "mbi" does not.

The letter-finding program succeeds regardles of cases (uppercase, lowercase, or a mix) in the string. 

This program plays with:
- Arrays/Slices
- Handling type string
- Using the strings package



### 3. Number Sorter
```
go run 03-sort-numbers.go
```
Go's array has a fixed length: it can only contain as many elements as what you set it to be when you create it.
However, in this program, you can add as many numbers as you want to the container.
The way we do that is with a slice, which is like a window to an array, instead of an array itself.
That way, we can keep increasing the size of the slice to contain every number you place in it.

So the concepts we cover in this program are:
- Slices (of varying length)
- Packages



### 4. JSON Maker
```
go run 04-make-json.go
```
A popular standard, JSON facilitates information exchange. Programs often need to communicate with other systems. Since systems may vary largely, standards are important.

Go just happens to provide a json package: its standardized implementation to encode JSON from its various structures.

This program demonstrates Go's ability to communicate in JSON.
It prompts the user to enter a name, and then an address.
It maps the address to the name, and the program returns the data's JSON object.

```
Go:
map[Name:Address]

JSON:
{"Name":"Address"}
```



