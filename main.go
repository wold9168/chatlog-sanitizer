package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// 定义一个正则表达式来匹配 "QQ名: MM-DD HH:MM:SS" 格式的行
	// ^      - 匹配行的开始
	// .+     - 匹配一个或多个任意字符 (QQ名)
	// :\s    - 匹配一个冒号和一个空格
	// \d{2}-\d{2} \d{2}:\d{2}:\d{2} - 匹配 "01-21 05:23:41" 这种时间格式
	// $      - 匹配行的结束
	headerRegex := regexp.MustCompile(`^.+:\s\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)

	// 使用 bufio.Scanner 从标准输入逐行读取，效率高且方便
	scanner := bufio.NewScanner(os.Stdin)

	// 循环扫描输入的每一行
	for scanner.Scan() {
		line := scanner.Text()

		// 如果当前行匹配了头部的正则表达式，就跳过（不输出）
		if headerRegex.MatchString(line) {
			continue
		}

		// 如果不匹配，说明是消息正文，直接打印到标准输出
		fmt.Println(line)
	}

	// 检查扫描过程中是否发生了错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取标准输入时出错:", err)
	}
}
