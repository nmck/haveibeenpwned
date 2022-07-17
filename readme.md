# HaveIBeenPwned CLI

A HaveIBeenPwned CLI written in Go.

Currently it can query the [Pwned Passwords API]('https://haveibeenpwned.com/Passwords')

## Run our Docker Container

```
docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version
```

## Run with Go Run

Run the application with `go run`

```
go run main.go pw YourPasswordHere
```

## Building the Application

Build your application into a static binary: <br/>

```
go build
```

This will produce a compiled program called `app`
You can run this program easily:

```
./main pw YourPasswordHere
```
