// go开发基础作业2
package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"sync"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值。

func add10(num *int) {
	*num += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。

func multiplyBy2(nums *[]int) {
	for i, num := range *nums {
		(*nums)[i] = num * 2
	}
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数

func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func printNumbers() {
	go printOdd()
	go printEven()
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

func taskScheduler(tasks []func()) {
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func(idx int, t func()) {
			start := time.Now()
			defer func() {
				duration := time.Since(start)
				fmt.Printf("task %d took %s\n", idx, duration)
				wg.Done()
			}()
			t()
		}(i, task)
	}
	wg.Wait()
}


// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {
	Area() float64
	Perimeter() float64
}


type Rectangle struct {
	Width  float64
	Height float64
}


func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}


// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息

type Person struct {
	Name string
	Age  int
}


type Employee struct {
	person Person
	EmployeeID string
}


func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %s\n", e.person.Name, e.person.Age, e.EmployeeID)
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

func generateNumbers() {
    ch := make(chan int)
    go func() {
        defer close(ch)  // 发送完后关闭通道
        for i := 1; i <= 10; i++ {
            ch <- i
        }
    }()
    go func() {
        for i := range ch {  // 通道关闭后，range 循环会结束
            fmt.Println(i)
        }
    }()
    time.Sleep(time.Second)  // 只需要很短时间让协程启动
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func generateNumbersWithBuffer() {
    ch := make(chan int, 100)
    
    // 生产者：异步发送，利用缓冲
    go func() {
        defer close(ch)
        for i := 1; i <= 100; i++ {
            ch <- i
            fmt.Printf("Sent: %d\n", i)  // 可以看到发送速度
        }
    }()
    
    // 消费者：同步接收
    go func() {
        for i := range ch {
            fmt.Printf("Received: %d\n", i)  // 可以看到接收速度
            time.Sleep(time.Millisecond * 10)  // 模拟处理时间
        }
    }()
    
    time.Sleep(time.Second * 10)  // 给足够时间观察
}


// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func mutexCounter() {
    var counter int
    var mu sync.Mutex
    var wg sync.WaitGroup
    
    // 启动10个协程
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            // 每个协程递增1000次
            for j := 0; j < 1000; j++ {
                mu.Lock()        // 获取锁
                counter++        // 临界区：递增操作
                mu.Unlock()      // 释放锁
            }
            
            fmt.Printf("Worker %d completed\n", workerID)
        }(i)
    }
    
    wg.Wait()  // 等待所有协程完成
    fmt.Printf("Final counter value: %d\n", counter)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func atomicCounter() {
    var counter int64  // 必须是int64，atomic包要求
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            for j := 0; j < 1000; j++ {
                atomic.AddInt64(&counter, 1)  // 原子递增操作
            }
            
            fmt.Printf("Worker %d completed\n", workerID)
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter)
}

func main() {
	// num := 10
	// add10(&num)
	// fmt.Println(num)

	// nums := []int{1, 2, 3, 4, 5}
	// multiplyBy2(&nums)
	// fmt.Println(nums)

	// printNumbers()

	// rectangle := Rectangle{Width: 10, Height: 20}
	// circle := Circle{Radius: 10}
	// fmt.Println(rectangle.Area())
	// fmt.Println(rectangle.Perimeter())
	// fmt.Println(circle.Area())
	// fmt.Println(circle.Perimeter())

	// employee := Employee{person: Person{Name: "John", Age: 30}, EmployeeID: "123456"}
	// employee.PrintInfo()

	// generateNumbers()
	// generateNumbersWithBuffer()
	// mutexCounter()
	atomicCounter()

}
