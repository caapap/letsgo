package main

import (
	"fmt"
	"sync"
)

// ProducerConsumer 实现生产者-消费者模式
func ProducerConsumer() {
	// 创建一个带缓冲的channel，容量为20（2个生产者各生产10个数据）
	dataChan := make(chan int, 20)
	
	// 创建两个WaitGroup，分别用于生产者和消费者
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup
	
	// 启动2个生产者
	for i := 1; i <= 2; i++ {
		producerWg.Add(1)
		go func(producerID int) {
			defer producerWg.Done()
			// 每个生产者生产10个数据
			for j := 1; j <= 10; j++ {
				data := (producerID-1)*10 + j
				dataChan <- data
				fmt.Printf("生产者 %d 生产数据: %d\n", producerID, data)
			}
		}(i)
	}
	
	// 启动3个消费者
	for i := 1; i <= 3; i++ {
		consumerWg.Add(1)
		go func(consumerID int) {
			defer consumerWg.Done()
			// 消费者持续从channel中读取数据
			for data := range dataChan {
				fmt.Printf("消费者 %d 消费数据: %d\n", consumerID, data)
			}
		}(i)
	}
	
	// 等待所有生产者完成
	producerWg.Wait()
	
	// 关闭channel，通知消费者没有更多数据了
	close(dataChan)
	
	// 等待所有消费者完成
	consumerWg.Wait()
	
	fmt.Println("所有数据已处理完毕，程序退出")
}

func main() {
	ProducerConsumer()
} 




