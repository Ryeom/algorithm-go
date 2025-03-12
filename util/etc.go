package util

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CreateTodayFolder 함수는 algorithm 폴더 하위에 오늘 날짜에 해당하는 폴더를 생성함.
func CreateTodayFolder(basePath string) error {
	// 현재 날짜 가져오기
	now := time.Now()
	year := now.Format("2006")      // 연도 (YYYY)
	monthDay := now.Format("01-02") // 월-일 (MM-DD)

	// 연도 폴더 경로
	yearPath := filepath.Join(basePath, year)
	// 월-일 폴더 경로
	todayPath := filepath.Join(yearPath, monthDay)

	// 연도 폴더가 없으면 생성
	if _, err := os.Stat(yearPath); os.IsNotExist(err) {
		if err := os.Mkdir(yearPath, os.ModePerm); err != nil {
			return fmt.Errorf("연도 폴더 생성 실패: %w", err)
		}
	}

	// 오늘 날짜 폴더가 없으면 생성
	if _, err := os.Stat(todayPath); os.IsNotExist(err) {
		if err := os.Mkdir(todayPath, os.ModePerm); err != nil {
			return fmt.Errorf("오늘 날짜 폴더 생성 실패: %w", err)
		}
	}

	fmt.Printf("폴더 생성 완료: %s\n", todayPath)
	return nil
}
