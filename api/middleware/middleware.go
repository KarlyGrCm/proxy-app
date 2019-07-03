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
			fmt.Println("PUSHED", tmp)
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

func SortedAppend(que []*Queue, request *Queue) []*Queue {
	index := sort.Search(len(que), func(i int) bool {
		return que[i].Weight > request.Weight && que[i].Priority > request.Priority
	})
	que = append(que, &Queue{})
	copy(que[index+1:], que[index:])
	que[index] = request
	return que
}

func customSorting(queue []*Queue) []*Queue {
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
	for _, row := range repo.Read() {
		fmt.Println("FROM SOURCE", row.Domain)
		Que = append(Que, row)
	}
	customSorting(Que)
	fmt.Println("QUEUE", Que)
	c.Next()
}
