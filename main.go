package main

import (
	"fmt"
	"ncc-select/client"
	"time"
)

func main() {
	c := client.NewClient()

	fmt.Println("输入TOKEN: ")
	var token string
	fmt.Scanln(&token)
	c.SetToken(token)

	courses, err := c.GetCourses()
	if err != nil {
		fmt.Println("获取课程列表失败: ", err)
		return
	}

	fmt.Println("课程列表:")
	for _, course := range courses {
		fmt.Printf("%+v\n", course)
	}

	fmt.Println("按回车键开始选课")
	fmt.Scanln()

	for {
		startTime := time.Now()
		err = c.SelectCourse(client.Course{
			CourseId: 1141,
		})
		end := time.Now()
		if err != nil {
			fmt.Printf("选课失败: %s, 耗时: %s, now: %s\n", err, end.Sub(startTime), end)
			time.Sleep(100 * time.Millisecond)
		} else {
			fmt.Println("选课成功")
			break
		}
	}

}
