package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/lint"
)

sampleJson :=
"{	\"attr1\": \"val1\", 	\"attr2\": \"val2\" }"

//func main() {
//	// BFS vs. DFS
//	// Depth 구분?
//	// 구분자: 큰따옴표, 콜론, 컴마, 중괄호, 슬라이스(대괄호)
//	// json parser 사용? vs. 직접 개발?
//	/*
//	for i := 0; i < length; ++i {
				
//	}
//	*/
//}