package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)
	go setList(channel)

	for v := range channel {
		fmt.Println("recebendo: ", v)
		time.Sleep(time.Second)
	}
}

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("enviando: ", i)
	}
	close(channel)
}

// 	var m sync.Mutex
// 	i := 0
// 	for x := 0; x < 10000; x++ {
// 		go func() {
// 			m.Lock()
// 			i++
// 			m.Unlock()
// 		}()
// 	}

// go ChangeNumber(&i, 5)
// go ChangeNumber(&i, 10)
// go ChangeNumber(&i, 20)

// 	time.Sleep(time.Second * 2)
// 	fmt.Println(i)
// }

// func ChangeNumber(i *int, newNumber int) {
// 	*i = newNumber

// 	wg.Add(3)
// 	callDatabase(&wg)
// 	callAPI(&wg)
// 	processInternal(&wg)

// 	wg.Wait()
// }

// func callDatabase(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finalizado callDatabase")
// 	wg.Done()
// }

// func callAPI(wg *sync.WaitGroup) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("finalizado callAPI")
// 	wg.Done()
// }

// func processInternal(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finalizado processInternal")
// 	wg.Done()
// }

// 	for i := 0; i < 10; i++ {
// 		go showMessage(strconv.Itoa(i))
// 	}

// 	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))
// }

//	func showMessage(message string) {
//		fmt.Println(message)
//	}
//}
