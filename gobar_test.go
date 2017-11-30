package pgbar

import (
	"testing"
	"time"
)

func TestPrintf(t *testing.T) {
	//printf(&Point{2, 2}, "test %d", 1)
	Printf("test %d\n\n", 1)
	Printf("test %d\n", 2)
	Println("test ", 3)
	Println("test ", 4)
	Println("test ", 5)
	Println("test ", 6)
	//fmt.Println(gCurrentLine, gMaxLine)
	for i := 0; i < 100; i++ {
		printf(2, "\rtest %d", 7+i)
		time.Sleep(time.Second / 3)
	}

}
