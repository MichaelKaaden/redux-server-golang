# Redux Server written in Go

This is a tiny REST service managing counters. The counters
are kept in memory, so they are reset every time you restart
the service.

Each counter has
- a unique index (a number greater or equal 0) and
- a value.

You can either get or set a counter. But in any distributed
environment, the latter would be bad practice. Use this only
for setting values for presentation purposes. Usually, you
would use the increment and decrement operations instead.

You can either get or set a counter. Of course, you shouldn't
set any counter in a distributed environment. Instead, you
should get it and then use the increment or decrement operations
on it. For presentations, it is a reasonable choice to set
some counters before showing anything to your audience.

## Building and Running the Server

To build and run the app, you need to install the Gin Web Framework
and its CORS middleware first.

To do so:
```bash
$ go get github.com/gin-gonic/gin
$ go get github.com/gin-contrib/cors
$ go run github.com/MichaelKaaden/redux-server-golang
```

## Alternative and Corresponding Implementations

This is only one possible solution to this kind of problem.

There are some implementations of single-page applications using the services which are implemented in different
programming languages and frameworks.

Here's the full picture.

## Client-Side Implementations

- [https://github.com/MichaelKaaden/redux-client-ngrx](https://github.com/MichaelKaaden/redux-client-ngrx) (Angular with
  NgRx)
- [https://github.com/MichaelKaaden/redux-client-ng5](https://github.com/MichaelKaaden/redux-client-ng5) (Angular
  with `angular-redux`)
- [https://github.com/MichaelKaaden/redux-client-ng](https://github.com/MichaelKaaden/redux-client-ng) (AngularJS
  with `ng-redux`)

## Server-Side Implementations

- [https://github.com/MichaelKaaden/redux-server-rust](https://github.com/MichaelKaaden/redux-server-rust) (Rust
  with `actix-web`)
- [https://github.com/MichaelKaaden/redux-server-golang](https://github.com/MichaelKaaden/redux-server-golang) (Go
  with `Gin`)
- [https://github.com/MichaelKaaden/redux-server-nest](https://github.com/MichaelKaaden/redux-server-nest) (Node.js
  with `Nest`)
- [https://github.com/MichaelKaaden/redux-server](https://github.com/MichaelKaaden/redux-server) (Node.js with `Exprss`)
