/*
 * @Author: 7erry
 * @Date: 2024-10-17 13:44:18
 * @LastEditTime: 2024-10-23 13:41:02
 * @Description: 持续发包抢课
 */
package main

import (
	"fmt"
	"time"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/course"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/utils"
)

func main() {

	c := client.NewFucker()

	t := config.TOKEN
	if t == "" {
		utils.Info("[*] 输入 TOKEN:")
		fmt.Scanln(&t)
	}
	c.SetToken(t)

	id := config.COURSE_ID
	if id == -1 {
		utils.Info("[*] 输入课程 ID:")
		fmt.Scanln(&id)
	}

	for {
		startTime := time.Now()
		err := c.SelectCourse(&course.Course{
			CourseId: id,
		})
		end := time.Now()
		if err != nil {
			fmt.Printf("[!] 选课失败: %s, 耗时: %s, now: %s\n", err, end.Sub(startTime), end)
			time.Sleep(100 * time.Millisecond)
		} else {
			utils.Info("[*] 选课成功")
			break
		}
	}
}
