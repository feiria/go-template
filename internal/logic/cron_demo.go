package logic

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Task1 struct {
}

func (t Task1) Run() {
	fmt.Println(time.Now(), "I am Task1")
}

type Task2 struct {
}

func (t Task2) Run() {
	fmt.Println(time.Now(), "I am Task2")
}

func CronExample() {
	c := cron.New(cron.WithSeconds())

	EntryID, err := c.AddJob("*/5 * * * * *", Task1{})
	fmt.Println(time.Now(), EntryID, err)

	EntryID, err = c.AddJob("*/10 * * * * *", Task2{})
	fmt.Println(time.Now(), EntryID, err)

	c.Start()
	select {}
}
