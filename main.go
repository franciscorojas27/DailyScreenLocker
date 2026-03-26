package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sys/windows"
)

type Job struct {
	Name     string
	GoalDate time.Time
}

func LockWorkStation() {
	user := windows.NewLazySystemDLL("user32.dll")
	lockWorkStation := user.NewProc("LockWorkStation")
	ret, _, err := lockWorkStation.Call()

	if ret == 0 {
		fmt.Printf("\nError trying to lock the screen: %v\n", err)
	} else {
		fmt.Println("\n[!] Windows was blocked")
	}
}

func main() {
	var mu sync.RWMutex

	now := time.Now()
	goal := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())

	if now.After(goal) {
		goal = goal.AddDate(0, 0, 1)
	}

	job := Job{
		Name:     "BlockScreen",
		GoalDate: goal,
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			mu.RLock()
			target := job.GoalDate
			mu.RUnlock()

			timeLeft := time.Until(target)
			fmt.Printf("\r[+] Next block screen: %v    ", timeLeft.Round(time.Second))
		}
	}()

	for {
		mu.RLock()
		target := job.GoalDate
		mu.RUnlock()

		time.Sleep(time.Until(target))

		LockWorkStation()

		mu.Lock()
		job.GoalDate = job.GoalDate.AddDate(0, 0, 1)
		mu.Unlock()

		time.Sleep(2 * time.Second)
	}
}
