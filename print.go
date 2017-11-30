package pgbar

import (
	"fmt"
	"strings"
	"sync"

	"github.com/qianlnk/to"
)

var (
	mu           sync.Mutex
	gSrcLine     = 0 //起点行
	gCurrentLine = 0 //当前行
	gMaxLine     = 0 //最大行
)

func move(line int) {
	//fmt.Println("\n\n\n\n", gCurrentLine, line)
	fmt.Printf("\033[%dA\033[%dB", gCurrentLine, line)
	gCurrentLine = line
}

func print(line int, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()

	move(line)
	var realArgs []interface{}
	realArgs = append(realArgs, "\r")
	realArgs = append(realArgs, args...)
	fmt.Print(realArgs...)
	move(gMaxLine)
}

func printf(line int, format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()

	move(line)
	fmt.Printf("\r"+format, args...)
	move(gMaxLine)
}

func println(line int, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()

	move(line)
	var realArgs []interface{}
	realArgs = append(realArgs, "\r")
	realArgs = append(realArgs, args...)
	fmt.Print(realArgs...)
	move(gMaxLine)
}

func Print(args ...interface{}) {
	mu.Lock()
	lf := countLF("", args...)
	if gMaxLine == 0 {
		gMaxLine += lf + 1
	} else {
		gMaxLine += lf
	}
	mu.Unlock()

	print(gMaxLine, args...)
}

func Printf(format string, args ...interface{}) {
	mu.Lock()

	lf := countLF(format, args...)
	if gMaxLine == 0 {
		gMaxLine += lf + 1
	} else {
		gMaxLine += lf
	}
	mu.Unlock()

	printf(gMaxLine, format, args...)
}

func Println(args ...interface{}) {
	mu.Lock()

	lf := countLF("", args...)
	lf++
	if gMaxLine == 0 {
		gMaxLine += lf + 1
	} else {
		gMaxLine += lf
	}
	mu.Unlock()

	println(gMaxLine, args...)
}

func countLF(format string, args ...interface{}) int {
	var count int
	count = strings.Count(format, "\n")
	for _, arg := range args {
		count += strings.Count(to.String(arg), "\n")
	}

	return count
}
