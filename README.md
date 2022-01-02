## go-ticket-booking-app

Simple CLI application which books ticket for a Go conference made to learn the fundamentals of Go programming language.

### Goroutine vs OS Threads
- Go doesnt directly use OS level threads, it uses `Green Threads` which are an `abstraction` over an actual thread.
- `Goroutines` are managed by the go runtime, we are only interacting with these high level go routines. While OS thrads are managed by kernel.
- These Goroutines are `cheaper` and `lightweight` while OS threads are costly and take some time to start up. Thus we can run thousands of go routines without affecting the performance much.

- Also, the have built-in functionality called `Channels` for internal communication i.e. receiving and sending data. Comparitively it is difficult to communicate in OS Threads.