package main

import (
	"encoding/json"
	"fmt"
	//	"go/types"
	"reflect"
)

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"},
		{"Name": "depth1",    "Order": {
				"depth2": "depth2Val"
			}
		}
	]`)

	//	type Animal struct {
	//		Name  string
	//		Order string
	//	}
	//	var animals []Animal
	//	err := json.Unmarshal(jsonBlob, &animals)
	//	if err != nil {
	//		fmt.Println("error:", err)
	//	}
	//	fmt.Printf("%+v", animals)

	var f []interface{}
	err := json.Unmarshal(jsonBlob, &f)
	if err != nil {
		fmt.Println("error:", err)
	}
	// 예제는 맵의 슬라이스 형태
	fmt.Printf("%+v\n", f)                                   // 전체
	fmt.Printf("%+v\n", f[0])                                // 첫번째
	fmt.Printf("reflect.TypeOf(f): %v\n", reflect.TypeOf(f)) // 타입

	// assemble struct
	for i, v := range f {
		// reflect.TypeOf(f)가 Slice 타입이면,
		//	if reflect.TypeOf(f) == Slice {
		//	if reflect.TypeOf(f) == Array {
		//	if reflect.TypeOf(f) == reflect.TypeOf([]interface{}) {
		if reflect.TypeOf(f) == f.(type) {
			// 그리고 빈 Slice가 아니라면,
			// Do something

			// Slice의 첫 번째 요소의 Map에서 Key를 얻는다.
			//			f[0]
		}
		// Slice 타입이 아니면
		//		else {}
	}
}
