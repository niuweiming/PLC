package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// 获取数据
	url := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"

	// 创建一个http.Client对象
	client := &http.Client{}

	// 创建一个http.Request对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("创建请求失败: %s", err)
	}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("发送GET请求失败: %s", err)
	}
	defer resp.Body.Close()

	// 读取并打印响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应体失败: %s", err)
	}
	result := string(body)
	// 将输入数据按空格和换行分割成字符串切片
	parts := strings.Fields(result)

	// 计算总和和数量
	var sum float64
	var count int
	for _, part := range parts {
		value, err := strconv.ParseFloat(part, 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			return
		}
		sum += value
		count++
	}

	// 计算平均值
	average := sum / float64(count)
	fmt.Printf("PLC Average: %.2f\n", average)
}
