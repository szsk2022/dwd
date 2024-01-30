package main

import (
	"bufio"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

/*
	@Author:@szsk2022
	@Email：admin@sunzishaokao.com
	@URL:https://gitee.com/szsk2/defender
*/

// readLines 读取文件并将其内容作为一段行返回。
func readLines(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

// 初始化termbox
func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}

// pause 按任意键继续(termbox)
func pause() {
	fmt.Print("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}
