package utils

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}

type Message struct {
	Body   string `json:"body"`
	Status string `json:"status"`
}

var (
	clients = make(map[string]*Client)
	mu      sync.Mutex
)

func RateLimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		if _, exists := clients[ip]; !exists {
			clients[ip] = &Client{Limiter: rate.NewLimiter(2, 4), LastSeen: time.Now()}
		}
		clients[ip].LastSeen = time.Now()
		mu.Unlock()
		if !clients[ip].Limiter.Allow() {
			msg := Message{
				Body:   "Rate limit exceeded. Try again in 2 seconds",
				Status: "429, Error!!!!!!!",
			}
			c.AbortWithStatusJSON(429, msg)
			return

		}

		c.Next()

	}
}

func CleanupClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.LastSeen) > 2*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}
