package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que declaration
var Que []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	var final []*Queue
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0
	tmp := &Queue{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		if scanner.Text() == "" {
			count = 0
			final = append(final, tmp)
			tmp = &Queue{}

			fmt.Println("OUT", scanner.Text())
			continue
		}
		switch count {
		case 1:
			tmp.Domain = scanner.Text()
		case 2:
			tmp.Weight, _ = strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
		case 3:
			tmp.Priority, _ = strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
		}
		fmt.Println("IN", scanner.Text())
	}

	return final
}

// MockQueue should mock an Array of Queues
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   1,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 1,
		},
	}
}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	fmt.Println("FROM HEADER", domain)
	for _, row := range repo.Read() {
		fmt.Println("FROM SOURCE", row.Domain)

		//  ALGORITHM HERE...
		//  USE QUE

	}
	Que = append(Que, domain)

	c.Next()
}
