package algorithm

import (
	"algorithm/util"
	"log"
	"os"
	"testing"
)

func TestCreateNewFolder(t *testing.T) {
	// 오늘 날짜 폴더 생성 실행
	pwd, _ := os.Getwd()
	if err := util.CreateTodayFolder(pwd); err != nil {
		log.Fatalf("폴더 생성 오류: %v", err)
	}
}
