/*
 * @Author: 7erry
 * @Date: 2024-10-17 13:44:18
 * @LastEditTime: 2024-10-23 13:43:00
 * @Description: 获取课表并保存到本地
 */
package main

import (
	"fmt"
	"os"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/utils"
	"gopkg.in/yaml.v2"
)

func main() {

	c := client.NewFucker()

	t := config.TOKEN
	if t == "" {
		utils.Info("[*] 输入 TOKEN:")
		fmt.Scanln(&t)
	}
	c.SetToken(t)

	targets, err := c.GetCourses()
	if err != nil {
		utils.Info("[!] 获取课程列表失败: ", err)
		return
	}

	utils.Info("课程列表:")
	for index, target := range *targets {
		utils.Info("[%d] %+v\n", index, target)
	}
	data, err := yaml.Marshal(&targets)
	utils.CheckIfError(err)
	os.WriteFile(
		config.COURSE_LIST,
		data,
		0666,
	)
}
