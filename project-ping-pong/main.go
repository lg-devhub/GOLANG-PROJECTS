package main
import (
	"fmt"
	"sync"
)

func ping(pingChan chan bool, pongChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-pingChan
		fmt.Println("ping")
		// Passa a vez para o pong
		pongChan <- true
	}
}

func pong(pingChan chan bool, pongChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		<-pongChan
		fmt.Println("pong")
		// Só devolve a vez se não for a última rodada
		if i < 4 {
			pingChan <- true
		}
	}
}

func main() {
	pingChan := make(chan bool)
	pongChan := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)
	
	go ping(pingChan, pongChan, &wg)
	go pong(pingChan, pongChan, &wg)
	// Começa o jogo dando a vez para o ping
	pingChan <- true
	wg.Wait()
}