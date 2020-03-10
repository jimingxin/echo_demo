package main

import (
	"time"
	"sync"
	"github.com/labstack/echo"
	"strconv"
	"net/http"
)

type (
	Stats struct {
		Uptime time.Time `json:"uptime"`
		RequestCount uint64 `json:"requestCount"`
		Statuses map[string]int `json:"statuses"`
		mutex sync.RWMutex // 读写锁
	}
)

func NewStatus() *Stats {
	return &Stats{
		Uptime:time.Now(),
		Statuses:make(map[string]int),
	}
}

func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount ++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK,s)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer,"Echo/3.0")
		return next(c)
	}
}

func main()  {
	e := echo.New()

	e.Debug = true

	// Status
	s := NewStatus()
	e.Use(s.Process)
	e.GET("/status",s.Handle)

	// Server header
	e.Use(ServerHeader)

	// handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"Hello World")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8087"))
}