package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// コマンド1の実行
	cmd1 := exec.Command("echo", "Command 1")
	if err := cmd1.Start(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// コマンド1の実行
	cmd2 := exec.Command("touch", "c.txt") //.Output()
	if err := cmd2.Start(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Printf("%s", out)
	// コマンド1の待機
	if err := cmd1.Wait(); err != nil {
		fmt.Println("Error", err)
	}

	if err := cmd2.Wait(); err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Both commands have completed execution")
}
