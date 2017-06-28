# What is this?

This repo contains an implementation of the monkey language specified by this book : https://interpreterbook.com/

It is the result of my progress with said book.

Build and Test status : [![Build & Test Status](https://api.travis-ci.org/gideondsouza/monkey.svg?branch=master)](https://travis-ci.org/gideondsouza/monkey)

Progress: 
- [x] Implemented a Lexer. (With tests)
- [ ] Implemented Parser (With Tests)
- [ ] Implemented Evaluation (With Tests)
- [ ] Additions to the Original

Monkey looks like the following:

    // Bind values to names with let-statements
    let version = 1;
    let name = "Monkey programming language";
    let myArray = [1, 2, 3, 4, 5];
    let coolBooleanLiteral = true;

    // Use expressions to produce values
    let awesomeValue = (10 / 2) * 5 + 30;
    let arrayWithValues = [1 + 1, 2 * 2, 3];

Monkey also supports function literals and we can use them to bind a function to a name

    // Define a `fibonacci` function
    let fibonacci = fn(x) {
    if (x == 0) {
        0            // Monkey supports implicit returning of values
    } else {
        if (x == 1) {
            return 1;// ... and explicit return statements
        } else {
            fibonacci(x - 1) + fibonacci(x - 2); // Recursion! Yay!
        }
    }
    };

... the data types it supports are : booleans, strings, hashes, integers and arrays. We can combine them!

    // Here is an array containing two hashes, that use strings as keys and integers
    // and strings as values
    let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

    // Getting elements out of the data types is also supported.
    // Here is how we can access array elements by using index expressions:
    fibonacci(myArray[4]);
    // => 5

    // We can also access hash elements with index expressions:
    let getName = fn(person) { person["name"]; };

    // And here we access array elements and call a function with the element as
    // argument:
    getName(people[0]); // => "Anna"
    getName(people[1]); // => "Bob"
