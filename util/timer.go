package util

import (
	"fmt"
	"reflect"
	"time"
)

// SolutionFunc 인터페이스 (모든 솔루션 함수 지원)
type SolutionFunc interface{}

// 자동 실행 시간 측정 래퍼 함수 (파라미터 지원)
func RunWithTimer(name string, fn SolutionFunc, args ...interface{}) interface{} {
	start := time.Now()

	// 함수 실행 (반환값을 저장)
	result := callFunction(fn, args...)

	elapsed := time.Since(start)
	fmt.Printf("%s Execution Time: %s\n", name, elapsed)

	return result
}

// Reflection을 이용해 함수 호출 및 결과 반환
func callFunction(fn SolutionFunc, args ...interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)

	// 인자들을 reflect.Value 타입으로 변환
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// 함수 실행 및 반환값 처리
	out := fnValue.Call(in)

	// 반환값이 없는 경우 nil 반환
	if len(out) == 0 {
		return nil
	}

	// 단일 반환값만 지원
	return out[0].Interface()
}
