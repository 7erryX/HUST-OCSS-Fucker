/*
 * @Author: 7erry
 * @Date: 2024-10-17 18:09:20
 * @LastEditTime: 2025-03-04 16:45:16
 * @Description:
 */
package client

import (
	"errors"
	"time"

	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/internal/client/course"
	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/internal/client/user"
	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/internal/utils"
	"github.com/imroc/req/v3"
)

var (
	ErrGetTimeDiffFailed = errors.New("[*] Get Time difference Failed ")
)

type Fucker struct {
	Client *req.Client
	Token  string
}

func NewFucker() *Fucker {
	return &Fucker{
		Client: req.C().ImpersonateChrome().SetTimeout(10 * time.Second),
	}
}

type Profile struct {
	Content string
}

func (c *Fucker) GetCapchaImage() ([]byte, string, error) {
	return user.GetCapchaImage(c.Client)
}

func (c *Fucker) Login(username string, password string, code string, uuid string) error {
	t, err := user.Login(c.Client, username, password, code, uuid)
	c.SetToken(t)
	return err
}

func (c *Fucker) SetToken(token string) {
	c.Token = token
	c.Client.SetCommonBearerAuthToken(token)
}

func (c *Fucker) GetProfile() (Profile, error) {
	p, err := user.GetProfile(c.Client)
	return parseProfile(p), err
}

func parseProfile(raw []byte) Profile {
	//	err = json.Unmarshal(resp.Bytes(), res)
	return Profile{Content: string(raw)}
}

func (c *Fucker) GetCourses() (*[]course.Course, error) {
	return course.GetCourses(c.Client)
}

func (c *Fucker) SelectCourse(target *course.Course) error {
	return course.SelectCourse(c.Client, target)
}

// * TimeDiff 实际上是 Client 发送请求时的 Client 本地时间与 Server 收到请求时的 Server 本地时间的时间差
func (c *Fucker) GetTimeDiff() (time.Duration, error) {
	c_date := time.Now()
	// resp, err := c.Client.R().Get("http://222.20.126.201/student/student/course")
	//* 这个 URL 好像不被限流为 200 次每接口每人每 12h 的请求频率
	resp, err := c.Client.R().Get("http://222.20.126.201/student/index")
	//* fmt.Println(c_date)
	//* fmt.Println(time.Now())
	utils.CheckIfError(err)
	//* fmt.Println(resp.TotalTime())

	date_str, ok := resp.Header["Date"]
	if !ok {
		return -1, ErrGetTimeDiffFailed
	}
	s_date, err := time.Parse(time.RFC1123, date_str[0])
	utils.CheckIfError(err)
	s_date = s_date.Local()
	//* utils.Info("Client Time:%s\nServer Time:%s", c_date, s_date)
	return c_date.Sub(s_date), nil
}
