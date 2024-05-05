package main

import (
	"fmt"
	"os"
	"bufio"
)
	


func main() {
	fmt.Println("Hello World")

	f, err := os.Open("test.txt")

	if err != nil{
		fmt.Println("File Opening Error")
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {                // テキスト行がなくなるまで、ループする
		line := scanner.Text()          // テキスト １行を読み取る
		// テキスト一行ごとの処理
		fmt.Println(line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}