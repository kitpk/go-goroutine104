package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Data string `json:"data"`
}

type PubSub struct {
	subs []chan Message
	mu   sync.Mutex
}

func (ps *PubSub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan Message, 1)
	ps.subs = append(ps.subs, ch)
	return ch
}

func (ps *PubSub) Publisher(msg *Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, sub := range ps.subs {
		sub <- *msg
	}
}

func (ps *PubSub) UnSubscribe(ch chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for i, sub := range ps.subs {
		if sub == ch {
			ps.subs = append(ps.subs[:i], ps.subs[i+1:]...)
			close(ch)
			break
		}
	}
}

func main() {
	app := fiber.New()

	pubsub := &PubSub{}

	app.Post("/publisher", func(c *fiber.Ctx) error {
		msg := new(Message)
		if err := c.BodyParser(&msg); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		pubsub.Publisher(msg)
		return c.JSON(&fiber.Map{
			"message": "add to subscriber",
		})
	})

	sub1 := pubsub.Subscribe()
	go func() {
		for msg := range sub1 {
			fmt.Println("Received message sub1:", msg)
		}
	}()

	sub2 := pubsub.Subscribe()
	go func() {
		for msg := range sub2 {
			fmt.Println("Received message sub2:", msg)
		}
	}()

	// pubsub.UnSubscribe(sub1)
	// pubsub.UnSubscribe(sub2)

	app.Listen(":8080")
}
