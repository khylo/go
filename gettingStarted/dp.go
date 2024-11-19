package main

import (	
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type Philosopher struct {
	id int
	host *Host
	left *Chopstick
	right *Chopstick
}

type Chopstick struct {
	id int
	sync.Mutex
}

type Host struct {
	eating chan int
}

func (p *Philosopher) eat(wg *sync.WaitGroup) {
	for i:=0; i<3; i++ {
		//fmt.Printf("  %d) is hungry\n", p.id)
		p.host.eating <- p.id // Add to Host chan of eaters
		
		//fmt.Printf("  %d) Locking c %d \n", p.id, p.left.id)
		p.left.Lock()
		//fmt.Printf("  %d) Locking c %d \n", p.id, p.right.id)
		p.right.Lock()
		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		//fmt.Printf("  %d) UnLocking c %d \n", p.id, p.right.id)
		p.right.Unlock()
		//fmt.Printf("  %d) UnLocking c %d \n", p.id, p.left.id)
		p.left.Unlock()
		//fmt.Printf("  %d) Tell host we've stopped eating \n",p.id)
		<-p.host.eating  // Subtract from Host chan of eaters
		fmt.Printf("finishing eating %d \n", p.id)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) //wait before eating again
	}
	wg.Done()
}

func createTable() []Philosopher {
	p := make([]Philosopher, 5)
	c := make([]Chopstick, 5)
	h := Host{eating: make(chan int, 2)}
	for i := 0; i < 5; i++ {
		c[i] = Chopstick{id: i+1}
	}
	for i := 0; i < 5; i++ {
		//fmt.Printf("P idx %d is seated. with left=%d and right=%d\n", i, i, (i+1)%5)
		p[i] = Philosopher{id: i+1, host: &h, left: &c[i], right: &c[((i+1)%5)]}
		//fmt.Printf("Philosopher %d is seated. with left=%d and right=%d\n", p[i].id, p[i].left.id, p[i].right.id)
	}
	return p
}

func main() {
	philosophers := createTable()
	var wg sync.WaitGroup
	for _,p := range philosophers {
		//fmt.Printf("  Adding to wg for %d",p.id)
		wg.Add(1)
		go p.eat(&wg)
	}
	wg.Wait()
	//time.Sleep(10 * time.Second)
	fmt.Println("All done!")
}