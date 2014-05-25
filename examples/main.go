package main

import (
	"fmt"

	"github.com/levicook/smaug"
)

func main() {
	fmt.Printf("smaug.Level: %q\n", smaug.Level)
	fmt.Println(``)

	fmt.Println(`smaug.Debug:`)
	smaug.Debug.Dump("Will Dump!")
	smaug.Debug.Print("Will Print\n")
	smaug.Debug.Printf("Will %q", "Printf")
	smaug.Debug.Println("Will Println")
	fmt.Println(``)

	fmt.Println(`smaug.Info:`)
	smaug.Info.Dump("Will Dump!")
	smaug.Info.Print("Will Print\n")
	smaug.Info.Printf("Will %q", "Printf")
	smaug.Info.Println("Will Println")
	fmt.Println(``)

	fmt.Println(`smaug.Warn:`)
	smaug.Warn.Dump("Will Dump!")
	smaug.Warn.Print("Will Print\n")
	smaug.Warn.Printf("Will %q", "Printf")
	smaug.Warn.Println("Will Println")
	fmt.Println(``)

	fmt.Println(`smaug.Error:`)
	smaug.Error.Dump("Will Dump!")
	smaug.Error.Print("Will Print\n")
	smaug.Error.Printf("Will %q", "Printf")
	smaug.Error.Println("Will Println")
	fmt.Println(``)

	fmt.Println("------------------------")
	smaug.Level = smaug.INFO
	fmt.Printf("smaug.Level: %q\n", smaug.Level)
	fmt.Println(``)

	fmt.Println(`smaug.Debug:`)
	smaug.Debug.Dump("Will not Dump!")
	smaug.Debug.Print("Will not Print\n")
	smaug.Debug.Printf("Will not %q", "Printf")
	smaug.Debug.Println("Will not Println")
	fmt.Println(``)

	fmt.Println(`smaug.Info:`)
	smaug.Info.Dump("Will Dump!")
	smaug.Info.Print("Will Print\n")
	smaug.Info.Printf("Will %q", "Printf")
	smaug.Info.Println("Will Println")
	fmt.Println(``)

	fmt.Println(`smaug.Warn:`)
	smaug.Warn.Dump("Will Dump!")
	smaug.Warn.Print("Will Print\n")
	smaug.Warn.Printf("Will %q", "Printf")
	smaug.Warn.Println("Will Println")
	fmt.Println(``)

	fmt.Println(`smaug.Error:`)
	smaug.Error.Dump("Will Dump!")
	smaug.Error.Print("Will Print\n")
	smaug.Error.Printf("Will %q", "Printf")
	smaug.Error.Println("Will Println")
	fmt.Println(``)

	fmt.Println("------------------------")
	smaug.Level = smaug.WARN
	fmt.Printf("smaug.Level: %q\n", smaug.Level)
	fmt.Println(``)

	fmt.Println(`smaug.Debug:`)
	smaug.Debug.Dump("Will not Dump!")
	smaug.Debug.Print("Will not Print\n")
	smaug.Debug.Printf("Will not %q", "Printf")
	smaug.Debug.Println("Will not Println")
	fmt.Println(``)

	fmt.Println(`smaug.Info:`)
	smaug.Info.Dump("Will not Dump!")
	smaug.Info.Print("Will not Print\n")
	smaug.Info.Printf("Will not %q", "Printf")
	smaug.Info.Println("Will not Println")
	fmt.Println(``)

	fmt.Println(`smaug.Warn:`)
	smaug.Warn.Dump("Will Dump!")
	smaug.Warn.Print("Will Print\n")
	smaug.Warn.Printf("Will %q", "Printf")
	smaug.Warn.Println("Will Println")
	fmt.Println(``)

	fmt.Println(`smaug.Error:`)
	smaug.Error.Dump("Will Dump!")
	smaug.Error.Print("Will Print\n")
	smaug.Error.Printf("Will %q", "Printf")
	smaug.Error.Println("Will Println")
	fmt.Println(``)

	fmt.Println("------------------------")
	smaug.Level = smaug.ERROR
	fmt.Printf("smaug.Level: %q\n", smaug.Level)
	fmt.Println(``)

	fmt.Println(`smaug.Debug:`)
	smaug.Debug.Dump("Will not Dump!")
	smaug.Debug.Print("Will not Print\n")
	smaug.Debug.Printf("Will not %q", "Printf")
	smaug.Debug.Println("Will not Println")
	fmt.Println(``)

	fmt.Println(`smaug.Info:`)
	smaug.Info.Dump("Will not Dump!")
	smaug.Info.Print("Will not Print\n")
	smaug.Info.Printf("Will not %q", "Printf")
	smaug.Info.Println("Will not Println")
	fmt.Println(``)

	fmt.Println(`smaug.Warn:`)
	smaug.Warn.Dump("Will not Dump!")
	smaug.Warn.Print("Will not Print\n")
	smaug.Warn.Printf("Will not %q", "Printf")
	smaug.Warn.Println("Will not Println")
	fmt.Println(``)

	fmt.Println(`smaug.Error:`)
	smaug.Error.Dump("Will Dump!")
	smaug.Error.Print("Will Print\n")
	smaug.Error.Printf("Will %q", "Printf")
	smaug.Error.Println("Will Println")
	fmt.Println(``)
}
