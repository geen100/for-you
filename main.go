package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type NotificationSender interface {
	SendNotification(message string) error
}

type MacNotificationSender struct{}

func (m MacNotificationSender) SendNotification(message string) error {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "休憩通知"`, message))
	return cmd.Run()
}

func startTimer(interval time.Duration, resetChan chan bool, sender NotificationSender) {
	timer := time.NewTimer(interval)
	hoursPassed := 1

	for {
		select {
		case <-timer.C:
			message := fmt.Sprintf("%d時間が経過しました。そろそろ休憩しませんか？", hoursPassed)
			_ = sender.SendNotification(message)
			hoursPassed++
			timer.Reset(interval)
		case <-resetChan:
			fmt.Println("スリープ解除が検出されました。タイマーをリセットします。")
			hoursPassed = 1
			timer.Reset(interval)
		}
	}
}

func monitorSleepWake(resetChan chan bool) {
	cmd := exec.Command("pmset", "-g", "log")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("コマンドの実行に失敗しました:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("コマンドの開始に失敗しました:", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Wake from") {
			resetChan <- true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("コマンドの読み取り中にエラーが発生しました:", err)
	}
}

func main() {
	resetChan := make(chan bool)
	sender := MacNotificationSender{}

	go startTimer(1*time.Hour, resetChan, sender)

	go monitorSleepWake(resetChan)

	select {}
}
