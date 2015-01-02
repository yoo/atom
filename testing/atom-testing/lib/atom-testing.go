package main

import (
	"fmt"

	"github.com/JohannWeging/atom"
	"honnef.co/go/js/console"
)

var currentFunc string
var encounteredError bool

func TestConfirmNoCallbacks() {
	start("TestConfirmNoCallbacks")

	options := &atom.ConfirmOption{}
	options.Message = "Test"
	options.DetailedMessage = "Testing atom.Confirm without callbacks"
	options.Buttons = make([]atom.Button, 2)
	options.Buttons[0].Name = ""
	options.Buttons[1].Name = "Click Here!"
	index := atom.Confirm(options)

	if !(index == 0 || index == 1) {
		logError("Wrong button index returned.")
	}

	stop()
}

func TestInDevMode() {
	start("TestInDevMode")

	options := confirmOptions("Dev Mode?", "Is Atom running in dev mode?", "Yes", "No")
	index := atom.Confirm(options)
	devMode := atom.InDevMode()

	if index == 0 && !devMode {
		logError("Was false, expected true.")
	} else if index == 1 && devMode {
		logError("Was true, expected false.")
	}

	stop()
}

func TestInSafeMode() {
	start("TestInSafeMode")

	options := confirmOptions("Dev Safe?", "Is Atom running in safe mode?", "Yes", "No")
	index := atom.Confirm(options)
	safeMode := atom.InSafeMode()

	if index == 0 && !safeMode {
		logError("Was false, expected true.")
	} else if index == 1 && safeMode {
		logError("Was true, expected false.")
	}

	stop()
}

func TestInSpecMode() {
	start("TestInSpecMode")

	options := confirmOptions("Dev Spec?", "Is Atom running in spec mode?", "Yes", "No")
	index := atom.Confirm(options)
	specMode := atom.InSafeMode()

	if index == 0 && !specMode {
		logError("Was false, expected true.")
	} else if index == 1 && specMode {
		logError("Was true, expected false.")
	}

	stop()
}

func TestGetVersion() {
	start("TestGetVersion")

	version := atom.GetVersion()
	detial := fmt.Sprintf("Is your atom version: %s?", version)
	options := confirmOptions("Version?", detial, "Yes", "No")
	index := atom.Confirm(options)

	if index != 0 {
		logError("Returned the wrong version: " + version + ".")
	}

	stop()
}

func start(funcName string) {
	currentFunc = funcName
	console.Log("Running: ", currentFunc, "\n")
}

func stop() {
	if encounteredError {
		console.Log("Finished with error: ", currentFunc, "\n")
	} else {
		console.Log("Finished successfully: ", currentFunc, "\n")
	}
	encounteredError = false
	currentFunc = ""
}

func logError(errMsg string) {
	encounteredError = true
	console.Error("Error in ", currentFunc, ":\n", errMsg, "\n")
}

func confirmOptions(msg, detail string, buttons ...string) *atom.ConfirmOption {
	options := &atom.ConfirmOption{}
	options.Message = msg
	options.DetailedMessage = detail
	options.Buttons = make([]atom.Button, len(buttons))
	for i, buttonName := range buttons {
		options.Buttons[i].Name = buttonName
	}

	return options
}

func main() {
	atom.Activate(activate)
}

func activate() {
	atom.Commands.Add("atom-workspace", "atom-testing:test", test)
}

func test() {
}
