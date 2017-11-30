package main

import (
	"sync"
	"time"

	"github.com/qianlnk/pgbar"
)

func main() {
	pgb := pgbar.New("test")
	pgbar.Println("1111111")
	b := pgb.NewBar("1st", 20000)
	pgbar.Println("2222222")
	// b2 := pgb.NewBar("2st", 10000)
	// pgbar.Println("3333333")
	// b3 := pgb.NewBar("3st", 30000)
	// b.SetSpeedSection(900, 100)
	// b2.SetSpeedSection(900, 100)
	// b3.SetSpeedSection(900, 100)

	// pgb1 := pgbar.New("test")
	// pgbar.Println("1111111")
	// b4 := pgb1.NewBar("1st", 20000)
	// pgbar.Println("2222222")
	// b5 := pgb1.NewBar("2st", 10000)
	// pgbar.Println("3333333")
	// b6 := pgb1.NewBar("3st", 30000)
	// b4.SetSpeedSection(900, 100)
	// b5.SetSpeedSection(900, 100)
	// b6.SetSpeedSection(900, 100)

	// b7 := pgbar.NewBar(0, "7st", 4000)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add()
			time.Sleep(time.Second / 100)
		}
	}()

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 10000; i++ {
	// 		b2.Add()
	// 		time.Sleep(time.Second / 1000)
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 30000; i++ {
	// 		b3.Add()
	// 		time.Sleep(time.Second / 1000)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 20000; i++ {
	// 		b4.Add()
	// 		time.Sleep(time.Second / 1000)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 10000; i++ {
	// 		b5.Add()
	// 		time.Sleep(time.Second / 1000)
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 30000; i++ {
	// 		b6.Add()
	// 		time.Sleep(time.Second / 300)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 30000; i++ {
	// 		b7.Add()
	// 		time.Sleep(time.Second / 50)
	// 	}
	// }()
	wg.Wait()
}
