/*
 * @Author: 7erry
 * @Date: 2024-11-07 13:19:33
 * @LastEditTime: 2025-03-04 16:58:41
 * @Description:
 */
package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/internal/utils"
)

func TestGetTimeDiff(t *testing.T) {
	f := NewFucker()
	tf := struct {
		sum   time.Duration
		count int
		avg   time.Duration
	}{
		time.Duration(0),
		0,
		time.Duration(0),
	}
	for i := 0; i < 20; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			time_diff, err := f.GetTimeDiff()
			utils.CheckIfError(err)
			tf.sum += time_diff
			tf.count += 1
			utils.Info("%v", tf.sum/time.Duration(tf.count))
		})
		time.Sleep(config.TIME_INTERVAL * time.Millisecond)
	}
}
