package pgbar

import (
	"fmt"
	"sync"
	"time"
)

type Bar struct {
	mu       sync.Mutex
	line     int
	prefix   string
	total    int
	width    int
	advance  chan bool
	done     chan bool
	currents map[string]int
	current  int
	before   int
	rate     int
	speed    int
	cost     int
	estimate int
	fast     int
	slow     int
	srcUnit  string
	dstUnit  string
	change   int
	closed   bool
}

var (
	bar1 string
	bar2 string
)

const (
	defaultFast = 20
	defaultSlow = 5
)

func initBar(width int) {
	for i := 0; i < width; i++ {
		bar1 += "="
		bar2 += "-"
	}
}

func NewBar(line int, prefix string, total int) *Bar {
	if total <= 0 {
		return nil
	}

	if line <= 0 {
		gMaxLine++
		line = gMaxLine
	}

	bar := &Bar{
		line:     line,
		prefix:   prefix,
		total:    total,
		fast:     defaultFast,
		slow:     defaultSlow,
		width:    100,
		advance:  make(chan bool),
		done:     make(chan bool),
		currents: make(map[string]int),
		change:   1,
	}

	initBar(bar.width)
	go bar.updateCost()
	go bar.run()

	return bar
}

func (b *Bar) SetUnit(src string, dst string, change int) {
	b.srcUnit = src
	b.dstUnit = dst
	b.change = change

	if b.change == 0 {
		b.change = 1
	}
}

func (b *Bar) SetSpeedSection(fast, slow int) {
	if fast > slow {
		b.fast, b.slow = fast, slow
	} else {
		b.fast, b.slow = slow, fast
	}
}

func (b *Bar) Add(n ...int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	step := 1
	if len(n) > 0 {
		step = n[0]
	}

	b.current += step

	lastRate := b.rate
	lastSpeed := b.speed

	b.count()

	if lastRate != b.rate || lastSpeed != b.speed {
		b.advance <- true
	}

	if b.rate >= 100 && !b.closed {
		b.closed = true
		close(b.done)
		close(b.advance)
	}
}

func (b *Bar) count() {
	now := time.Now()
	nowKey := now.Format("20060102150405")
	befKey := now.Add(time.Minute * -1).Format("20060102150405")
	b.currents[nowKey] = b.current
	if v, ok := b.currents[befKey]; ok {
		b.before = v
	}
	delete(b.currents, befKey)

	b.rate = b.current * 100 / b.total
	if b.cost == 0 {
		b.speed = b.current * 100
	} else if b.before == 0 {
		b.speed = b.current * 100 / b.cost
	} else {
		b.speed = (b.current - b.before) * 100 / 60
	}

	if b.speed != 0 {
		b.estimate = (b.total - b.current) * 100 / b.speed
	}
}

func (b *Bar) updateCost() {
	for {
		select {
		case <-time.After(time.Second):
			b.cost++
			b.mu.Lock()
			b.count()
			b.mu.Unlock()
			b.advance <- true
		case <-b.done:
			return
		}
	}
}

func (b *Bar) run() {
	for range b.advance {
		printf(b.line, "\r%s", b.barMsg())
	}
}

func (b *Bar) barMsg() string {
	unit := ""
	change := 1
	if b.srcUnit != "" {
		unit = b.srcUnit
	}

	if b.dstUnit != "" {
		unit = b.dstUnit
		change = b.change
	}

	prefix := fmt.Sprintf("%s", b.prefix)
	rate := fmt.Sprintf("%3d%%", b.rate)
	speed := fmt.Sprintf("%3.2f %s ps", 0.01*float64(b.speed)/float64(change), unit)
	cost := b.timeFmt(b.cost)
	estimate := b.timeFmt(b.estimate)
	ct := fmt.Sprintf(" (%d/%d)", b.current, b.total)
	barLen := b.width - len(prefix) - len(rate) - len(speed) - len(cost) - len(estimate) - len(ct) - 10
	bar1Len := barLen * b.rate / 100
	bar2Len := barLen - bar1Len

	realBar1 := bar1[:bar1Len]
	var realBar2 string
	if bar2Len > 0 {
		realBar2 = ">" + bar2[:bar2Len-1]
	}

	msg := fmt.Sprintf(`%s %s%s [%s%s] %s %s in: %s`, prefix, rate, ct, realBar1, realBar2, speed, cost, estimate)
	switch {
	case b.speed <= b.slow*100:
		return "\033[0;31m" + msg + "\033[0m"
	case b.speed > b.slow*100 && b.speed < b.fast*100:
		return "\033[0;33m" + msg + "\033[0m"
	default:
		return "\033[0;32m" + msg + "\033[0m"
	}
}

func (b *Bar) timeFmt(cost int) string {
	var h, m, s int
	h = cost / 3600
	m = (cost - h*3600) / 60
	s = cost - h*3600 - m*60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
