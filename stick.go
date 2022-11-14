package stick

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Bar struct {
	Task       int64     // tasks to complete
	Done       int64     // completed task
	Length     int64     // length of progress bar
	percent    float32   // percentage
	doneBar    string    // bar to be displayed
	BarStr     string    // bar string to use
	maxTaskChr int       // Maximum number of task character
	startTime  time.Time // timestamp for start
	lastIter   time.Time // timestamp for previous iteration
}

func Init(task int64) Bar {
	var bar Bar
	bar.BarStr = "█"
	bar.Length = 20
	bar.Task = task
	bar.startTime = time.Now()
	bar.maxTaskChr = len(strconv.Itoa(int(task)))
	return bar
}

func (b *Bar) calcPercent() {
	b.percent = float32(b.Done) / float32(b.Task)
}

func (b *Bar) calcEnd() time.Duration {
	estSec := time.Since(b.lastIter).Seconds() * float64(b.Task-b.Done)
	// estSec := time.Since(b.startTime).Seconds() / float64(b.Done) * float64(b.Task-b.Done)
	return time.Second * time.Duration(estSec)
}

func (b *Bar) updateBar() {
	b.calcPercent()
	b.doneBar = strings.Repeat(b.BarStr, int(b.percent*float32(b.Length)))
}

func (b *Bar) printBar() {
	elapsed := formatDuration(time.Since(b.startTime))
	estEnd := formatDuration(b.calcEnd())
	fmt.Printf("\r %3.0f%%|%-*s|%*d/%*d [%s<%s]",
		b.percent*100,
		b.Length,
		b.doneBar,
		b.maxTaskChr,
		b.Done,
		b.maxTaskChr,
		b.Task,
		elapsed,
		estEnd)
}

func (b *Bar) Add(done int64) {
	b.Done = b.Done + done
	b.updateBar()
	b.printBar()
	if b.Done == b.Task {
		b.complete()
	}
	b.calcEnd()
	b.lastIter = time.Now()
}

func (bar *Bar) complete() {
	fmt.Println()
}

func formatDuration(d time.Duration) string {
	min := int(d.Seconds()/60) % 60
	sec := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}
