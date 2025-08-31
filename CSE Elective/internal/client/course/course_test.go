package course

import (
	"testing"
	"time"

	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/imroc/req/v3"
)

func TestGetCourses(t *testing.T) {
	c := req.C().ImpersonateChrome().SetTimeout(10 * time.Second)
	c.SetCommonBearerAuthToken(config.TOKEN)
	courses, err := GetCourses(c)
	if err != nil {
		t.Errorf("GetCourses() failed: %v", err)
		return
	}

	t.Logf("GetCourses() success: courses=%+v", courses)
}

func TestGetCourseNumber(t *testing.T) {
	c := req.C().ImpersonateChrome().SetTimeout(10 * time.Second)
	c.SetCommonBearerAuthToken(config.TOKEN)
	err := GetCourseClassNumber(c, &Course{CourseId: 1146})
	if err != nil {
		panic(err)
	}
}

func TestClient_SelectCourse(t *testing.T) {
	c := req.C().ImpersonateChrome().SetTimeout(10 * time.Second)
	c.SetCommonBearerAuthToken(config.TOKEN)
	target := Course{
		CourseId: 1146,
	}
	err := SelectCourse(c, &target)
	if err != nil {
		t.Errorf("SelectCourse() failed: %v", err)
		return
	}

	t.Logf("SelectCourse() success")
}
