package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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
var Que []*Queue

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
			final = append(final, tmp)
			tmp = &Queue{}
		}
		fmt.Println(tmp)
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

func sortedAppend(que []*Queue, request *Queue) []*Queue {
	index := sort.Search(len(que), func(i int) bool {
		if que[i].Weight < request.Weight {
			return que[i].Priority < request.Priority
		}
		if que[i].Priority < request.Priority {
			return que[i].Weight < request.Weight
		}
		return false
	})
	que = append(que, &Queue{})
	copy(que[index+1:], que[index:])
	que[index] = request
	return que
}

func CustomSorting(queue []*Queue) []*Queue {
	sort.Slice(queue, func(i, j int) bool {
		if queue[i].Weight < queue[j].Weight {
			return true
		}
		if queue[i].Weight > queue[j].Weight {
			return false
		}
		return queue[i].Priority < queue[j].Priority
	})
	return queue
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

	// Method 1: Insert and then sort by Weight and Priority
	Que = repo.Read()
	CustomSorting(Que)
	for _, row := range Que {
		fmt.Println("FROM SOURCE and Ordered", row.Domain)
		fmt.Println(row.Priority, row.Weight)
	}

	var OrderedQueue []*Queue
	// Method 2, read each value and pushed ordered
	for _, row := range repo.Read() {
		OrderedQueue = sortedAppend(OrderedQueue, row)
	}
	for _, row := range OrderedQueue {
		fmt.Println("FROM SOURCE and Ordered 2", row.Domain)
		fmt.Println(row.Priority, row.Weight)
	}

	c.Next()
}
