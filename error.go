package berr

import "log"

// Runs the callback when `err != nil`.
//
// Preferebly only use it to share error handling between different calls.
func Try(err error, callback func(error)) {
	if err != nil {
		callback(err)
	}
}

func TryOrPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func TryOrFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
