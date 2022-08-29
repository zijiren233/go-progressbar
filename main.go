package progressbar

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type progress struct {
	updateChan  chan uint
	maxProgress uint
}

func NewProgressBar(maxProgress uint) *progress {
	return &progress{updateChan: make(chan uint), maxProgress: maxProgress}
}

func (c *progress) ProgressBar() {
	c.progressBar()
}

func (c *progress) progressBar() {
	ctx, cf := context.WithCancel(context.Background())
	wate := sync.WaitGroup{}
	wate.Add(1)
	go icon(ctx, &wate)
	for i := range c.updateChan {
		fmt.Printf(" [%s%s] %.0f%%\r", strings.Repeat("=", int(i)), strings.Repeat(" ", int(c.maxProgress-i)), float32(i)/float32(c.maxProgress)*100)
		if i >= c.maxProgress {
			break
		}
	}
	cf()
	wate.Wait()
}

func (c *progress) Update(Progress int) {
	if uint(Progress) >= c.maxProgress {
		c.updateChan <- c.maxProgress
	} else {
		c.updateChan <- uint(Progress)
	}
}

func icon(ctx context.Context, wate *sync.WaitGroup) {
	defer wate.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf(" \n")
			return
		default:
			switch time.Now().Second() % 4 {
			case 0:
				fmt.Printf("\\\r")
			case 1:
				fmt.Printf("|\r")
			case 2:
				fmt.Printf("/\r")
			case 3:
				fmt.Printf("-\r")
			}
		}
	}
}
