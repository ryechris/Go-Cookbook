# Riyan's Go Cookbook

A set of small programs that each performs a specific task, which demonstrates certain feature(s) of Go.

Riyan is responsible for all the Go code, but CS Professor Ian Harris of UCI supplies the inspiration.

## Table of Contents
- [0. Requisites](#0-requisites-go-and-git)
- [Data Types](#data-types)
    + [1. Truncate Decimals](#1-decimal-truncator)
    + [2. Find IBM](#2-ibm-finder)
    + [3. Sort Numbers](#3-number-sorter)
- [Formats](#formats)
    + [4. Make JSON](#4-json-maker)
    + [5. Read File](#5-file-reader)
- [Functions](#functions)
    + [6. Sort Bubble](#6-bubble-sorter)
    + [7. Compute Physics](#7-physics-calculator)
- Object Orientation (in Go)
    + Farm Animals
    +  
- Concurrency
    +  

#### 0. Requisites: Go and Git
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


### Data Type

#### 1. Decimal Truncator
```
go run 01-truncate-decimals.go
```
This program prompts the user to type a floating number (for example, 3.1415).
It then removes the decimal point and everything to the right of it (e.g. input: 3.14; result: 3).

Though small and clean, it demonstrates clearly the following concepts:
- User input
- Type handling & conversion


#### 2. IBM Finder
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



#### 3. Number Sorter
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


### Format

#### 4. JSON Maker
```
go run 04-make-json.go
```
A popular standard, JSON facilitates information exchange. Programs often need to communicate with other systems. Since systems vary, standards are important.

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

Hence, this program contains the following concepts:
- Map
- JSON



#### 5. File Reader
Programs of course must be able to talk with files. So here we illustrate Go's ability to do that.

When you did git pull, included in the folder are two example files:
```
05-file1.txt
05-file2.txt
```
Each of these files has a list of names: on each line are two name values, separated by a comma.
```
George,Washington
Abraham,Lincoln
```
We will use these files in the program.  When the user runs the following line:
```
go run 05-read-file.go
```
the program presents two options: file1 or file2. If the user types 1, the program processes file 1.
If the user types 2, the program processes file 2. If the user types something else, the program will output an error.

The program processes the file by
1. opening the file
2. reading the file
3. closing the file
In closing, it returns the data to the user in this format: 
```
First Name: George
Last Name: Washington

First Name: Abraham
Last Name: Washington 
```
Of course we can do anything with that data, such as printing it in JSON.
But we have already done a JSON example in file 04, so here we show another capability of Go programs.

In so doing, we are able to focus more on fewer ideas -- since it is a cookbook -- instead of having too many ingredients in the program:
- File Access
- Slice: type, iteration
- Flow control: loop, if, switch
- Functions


#### 6. Bubble Sorter
```
go run 06-sort-bubble.go
```
We continue our theme of code organization from program 05.

However, program 06 is not a continuation of prog 05.

This program asks input from user, and places them into an array (slice, actually).
The program then sorts the numbers in ascending order.
When it concludes, it displays the result to the user. 

For example, in organizing the code, lines that repeat often is placed into its own function for a higher-quality code writing.

Note: The algorithm used here, the Bubble Sort algorithm, is a poor choice for production.
Its inefficiency is pointed out by the comments in the code.
But, the algorithm is useful to demonstrate the following topics:
- Organization
- Control Flow: nested loops
- Functions


#### 7. Physics Calculator
Having covered functions in the previous recipe, we are now going to have fun to dive deeper into functions!

We assume the user of this .go file
```
go run 07-compute-displacement.go
```
has some observed an object moving with a constant acceleration, such as a pen falling from a table.
Hence, the program requests three pieces of data:
- the constant acceleration (e.g. gravity, 9.8 m/s2)
- the velocity of the object at the start of the observation (e.g. object starts at rest, 0 m/s)
- and the starting displacement (e.g. 0 m)

With these values reported (inputted) by the user, we will compute displacement.

In the process, we learn about functions in Go being First Class values:
- Functions as variables
- Functions as return values
- Function's environment & lexical scoping


### Object Orientation in Go

#### 8. 
Go does not have a class the way other Object Oriented languages have, such as Python.

However, Go has an equivalent to class. We can "create a class" by associating data and methods.

If there are various data types -- like a class often has -- you can place them in a struct:
```
type Point struct {
    x float64
    y float64
}
```
Of course, you can include as many pieces and types of data as you want in a struct.
After you have a struct, you make methods, and you ensure those methods are available only for the struct you have in mind.

In this project, we illustrate how to ...













