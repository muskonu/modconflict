//go:build linux
// +build linux

package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func System(s string) {
	// 要执行的命令
	cmd := exec.Command("bash", "-c", s)

	// 捕获命令输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 捕获错误输出（可选）
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Stderr:", stderr.String())
		return
	}

	fmt.Println(out.String())
}
