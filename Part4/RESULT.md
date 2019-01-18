# Choosing a language
For implementing the project I would like to use the programming language GO.

### Think about how want to move data around (reading buttons, network, setting motor & lights, state machines, etc). Do you think in a shared-variable way or a message-passing way? Will you be using concurrency at all?
Every active or passive hardware can be represented by a goroutine. By communicating between the goroutines with messages the use of shared-variables can be avoided. If every hardware component is represented by a goroutine there should be no need for concurrency at all.

### How will you split into modules? Functions, objects, threads? Think about what modules you need, and how they need to interact. This is an iterative design process that will take you many tries to get "right" (if such a thing even exists!).
Modules:
- Button Up
- Button Down
- Motor
- Light
- Elevator Manager (logic)
- ...

### The networking part is often difficult. Can you find anything useful in the standard libraries, or other libraries?
The language GO already provides the channels.

### While working on new sections on the project you'll want to avoid introducing bugs to the parts that already work properly. Does the language have a framework for making and running tests, or can you create one? Testing multithreaded code is especially difficult.
Testing frameworks:
- [Native GO testing](https://golang.org/pkg/testing/)
- [Native GO testing Example](https://blog.alexellis.io/golang-writing-unit-tests/)
- [Native GO testing important facts](https://segment.com/blog/5-advanced-testing-techniques-in-go/)

### Code analysis/debugging/IDE support?
Code analysis:
- [Clean code style](https://github.com/mre/awesome-static-analysis#go)
- [Clean code style - Checkstyle](https://github.com/qiniu/checkstyle)

Debugging:
- IDE?

IDE:
- [GoLand by JetBrains](https://www.jetbrains.com/go/)
