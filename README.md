smaug
=====

Package smaug implements a simple leveled logging package.

There are four logging levels: `DEBUG`, `INFO`, `WARN` and `ERROR`.
And four corresponding loggers: `Debug`, `Info`, `Warn` and `Error`.

Each logger provides four logging methods: `Dump`, `Print`, `Printf` and
`Println`.

`Dump` is a pretty printer.

`Print`, `Printf` and `Println` wrap the equivalent functions in the
standard library log package.

```go
// set our log level to INFO
smaug.Level = smaug.INFO

// logging through Debug will not create ouptut
smaug.Debug.Print("you won't see me")

// logging through Info, Warn and will create ouptut
smaug.Info.Print("you will see me")
smaug.Warn.Print("you will see me")
smaug.Error.Print("you will see me")
```

Package smaug also contains convenience methods for single error
handling, where you're okay with panicing or exiting your process.

```go
// this will call log.Fatal:
smaug.FatalIf(func () error {
	return errors.New("oh no!")
})

// this will call log.Panic:
smaug.PanicIf(func () error {
	return errors.New("oh no!")
})
```

Overall, no big shakes. Just a lighweight wrapper around other packages.
