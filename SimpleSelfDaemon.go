package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var globalPid = -1

// 初始值为一个不存在的pid

var frequency = 3

// 该全局变量设定每隔几秒钟检测一次子进程存活，默认为3秒

func isProcessAlive(pid int) (bool, error) {
	// 输入pid参数检测对应进程是否存活

	pidStr := strconv.Itoa(pid)
	cmd := exec.Command("tasklist", "/svc")
	// 命令行查询进程列表
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// 将输出解析为行
	lines := strings.Split(string(output), "\n")

	// 遍历行查找指定的PID
	for _, line := range lines {
		// 检查行中是否包含PID
		if strings.Contains(line, pidStr) {
			return true, nil
			// 进程存活，返回true
		}
	}

	// 如果未找到指定PID，则返回false
	return false, nil

}

func main() {
	// 检查是否带有命令行参数
	if len(os.Args) == 1 {
		// 获取当前可执行文件的名称
		executable, err := os.Executable()
		if err != nil {
			fmt.Println("[-]Failed to get executable path:", err)
			return
		}

		// 每隔三秒钟检查新进程是否存活
		for {
			time.Sleep(time.Duration(frequency) * time.Second)

			// 检查进程是否存活
			exists, err := isProcessAlive(globalPid)
			if err != nil {
				fmt.Println("[-]Error:", err)
				return
			}

			if exists {
				fmt.Println("[+]Process exists. pid:" + strconv.Itoa(globalPid))
			} else {
				fmt.Println("[*]Process does not exist. Restarting...")
				// 重新启动进程
				cmd := exec.Command(executable, "-real")
				err = cmd.Start()
				globalPid = cmd.Process.Pid
				if err != nil {
					fmt.Println("[-]Failed to restart process:", err)
					return
				}
			}

		}
	} else {
		if os.Args[1] == "-real" {
			// 如果带有参数"-real"，则说明是被作为子进程创建，执行真正的程序功能
			fmt.Println("[*]Running real program functionality...")
			// 以下部分替换为真正的程序功能
			// 此处被注释的示例代码是每隔10秒打开一次计算器，注释后仅仅每隔10秒空循环一次
			for {
				time.Sleep(10 * time.Second)
				// test := exec.Command("cmd", "/c", "calc.exe")
				// test.Start()
			}
			// 程序功能结束
		}
	}
}
