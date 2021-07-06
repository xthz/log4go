package log4go

import "testing"

func TestInfo(t *testing.T) {
	for i := 0; i < 10; i++ {
		Info("你好, 世界", i)
		Error("不好, 世界", i)
		Logger.Info("asss")
	}
}
