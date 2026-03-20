package main

import (
	"time"
	"syscall"
)

func main() {
	for {
		ahora := time.Now()
		// Meta: Hoy a las 17:00 (5 PM)
		meta := time.Date(ahora.Year(), ahora.Month(), ahora.Day(), 17, 00, 0, 0, ahora.Location())

		if ahora.After(meta) {
			meta = meta.AddDate(0, 0, 1)
		}

		time.Sleep(time.Until(meta))

		syscall.NewLazyDLL("user32.dll").NewProc("LockWorkStation").Call()

		time.Sleep(1 * time.Minute)
	}
}
