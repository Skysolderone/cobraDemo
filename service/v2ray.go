package service

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

// 开启
func StartV2ray(cmd *cobra.Command, args []string) {
	exs := exec.Command("sh", "/root/start_v2ray.sh")

	err := exs.Start()
	if err != nil {
		fmt.Println("exec err:", err)
		return
	}
	fmt.Println("successfull ")
}

// 关闭
func CloseV2ray(cmd *cobra.Command, args []string) {
	exs := exec.Command("sh", "-c", "pgrep  v2ray")
	pid, err := exs.Output()
	index := bytes.Index(pid, []byte("\n"))
	pid = pid[:index]
	if err != nil {
		fmt.Println("CLOSE ERR", err)
		return
	}
	id, err := strconv.Atoi(string(pid))
	if err != nil {
		fmt.Println("pares err", err)
		return
	}

	err = syscall.Kill(id, syscall.SIGKILL)
	if err != nil {
		fmt.Println("KILL pid err:", id)
		return
	}
	fmt.Println("stop success")
}
