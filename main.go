package main

import (
	"fmt"
	"ncc-select/client"
	"os"
	"strings"
)

func main() {
	c := client.NewClient()

	config, err := os.ReadFile("config")
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return
	}

	lines := strings.Split(string(config), "\n")
	username := strings.TrimSpace(lines[0])
	password := strings.TrimSpace(lines[1])

	// 获取验证码图片
	img, uuid, err := c.GetCapchaImage()
	os.WriteFile("captcha.jpg", img, 0644)
	if err != nil {
		fmt.Println("获取验证码图片失败: ", err)
		return
	}
	fmt.Printf("输入验证码, 在当前目录下captcha.jpg中查看验证码图片: ")
	var code string
	fmt.Scanln(&code)

	err = c.Login(username, password, code, uuid)
	if err != nil {
		fmt.Println("登录失败: ", err)
		return
	}

	fmt.Println("登录成功")

	courses, err := c.GetCourses()
	if err != nil {
		fmt.Println("获取课程列表失败: ", err)
		return
	}

	fmt.Println("课程列表:")
	for _, course := range courses {
		fmt.Printf("%+v\n", course)
	}

}
