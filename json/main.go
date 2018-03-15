// Compact Json Format

package main

import (
	"encoding/json"
	"fmt"
	//	"go/types"
	"reflect"
	"strconv"
	"strings"
)

const (
	TYPE = iota
	VALUE
)

func parseType(typeField string) reflect.Type {
	lowerCaseType := strings.ToLower(typeField)
	switch lowerCaseType {
	case "int":
		//		fmt.Printf("reflect.TypeOf(int(0)): (%T)%v\ntypes.Int: (%T)%v\nreflect.Int: (%T)%v\n", reflect.TypeOf(int(0)), reflect.TypeOf(int(0)), types.Int, types.Int, reflect.Int, reflect.Int)
		return reflect.TypeOf(int(0))

		// reflect.Kind Type 반환
		//		return reflect.Int
		// types.BasicKind Type 반환
		//		return types.Int

	case "string":
		return reflect.TypeOf("")

	default:
		fmt.Println("@parseType() Exception: Parsing을 지원하지 않는 타입")
		return nil
	}
}

func main() {
	var jsonSample = []byte(`{"Struct":{
				"id": ["int", "987"],
				"issuer": ["string", "issuerVal"]
			}
		}`)

	var unmarshaledJson interface{}
	errByUnmarshal := json.Unmarshal([]byte(jsonSample), &unmarshaledJson)
	if errByUnmarshal != nil {
		// TODO Exception Handling
	}

	fmt.Printf("unmarshaledJson: (%T)%+v\n\n\n", unmarshaledJson, unmarshaledJson)

	// Empty Interface는 Map Type이 아니다! => Key, Value 접근 불가능
	// => Map Type으로 Casting
	castedJson := unmarshaledJson.(map[string]interface{})

	// Key "Struct"에 매핑된 Value를 uncastedValueOfStruct에 할당
	uncastedValueOfStruct, exists := castedJson["Struct"]
	if !exists {
		// TODO Exception Handling
	}
	/* 깔때사용
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	func ValueOf(i interface{}) Value {
		if i == nil {
			return Value{}
		}

		// TODO: Maybe allow contents of a Value to live on the stack.
		// For now we make the contents always escape to the heap. It
		// makes life easier in a few places (see chanrecv/mapassign
		// comment below).
		escapes(i)

		return unpackEface(i)
	}
	*/

	//--------------------------------------------------------------------------
	// Slice Type을 받으면 루프 돌면서 map으로 Casting
	// uncastedValueOfStruct의 Type은 interface{}, Underlying Type은 []interface{}
	//--------------------------------------------------------------------------
	fmt.Printf("uncastedValueOfStruct: (%T)%+v\n", uncastedValueOfStruct, uncastedValueOfStruct)
	valueOfStruct := uncastedValueOfStruct.(map[string]interface{})

	// id: [x, x] => map[string]interface{} => map[string][]interface{} => map[string][]string
	var valueOfFieldName []interface{}
	var nameField, typeField string
	var structFields []reflect.StructField
	var fieldValues = make(map[string]interface{}) // Field의 값을 보관하기 위한 Map

	// id:
	for k, v := range valueOfStruct {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
		nameField = strings.Title(k)
		valueOfFieldName = v.([]interface{})

		// id: [int, 987]
		for i, value := range valueOfFieldName {
			switch i {
			case TYPE: // e.g. int, string
				typeField = value.(string)
			case VALUE: // e.g. 987, "issuerVal"
				fieldValues[nameField] = value.(string)
			default:
				// TODO Exception Handling
			}
		}

		fmt.Println(k + "의 Name: " + nameField)
		fmt.Println(k + "의 Type: " + typeField)
		fmt.Println("fieldValues[" + k + "]: " + fieldValues[nameField].(string))

		//--------------------------------------------------------------------------
		// StructField 생성해서 structFields Slice에 넣기
		//--------------------------------------------------------------------------
		//		structField := reflect.StructField{Name: nameField, Type: parseType(typeField), Tag: reflect.StructTag("json:" + strings.ToLower(nameField))}
		//		fmt.Printf("%+v\n", structField)
		//		structFields = append(structFields, structField)
		structFields = append(structFields, reflect.StructField{Name: nameField, Type: parseType(typeField), Tag: reflect.StructTag("json:" + strings.ToLower(nameField))})
	}

	fmt.Printf("len(structFields): %v\n", len(structFields))
	fmt.Printf("structFields: %v\n", structFields)

	//--------------------------------------------------------------------------
	// struct 정의 및 instantiate
	//--------------------------------------------------------------------------
	structDef := reflect.StructOf(structFields)

	// New()로 struct 정의, Elem()으로 instantiate.
	// Elem()은 instance를 Value Type으로 반환한다.
	instance := reflect.New(structDef).Elem()

	// Data 넣기 전 출력
	fmt.Printf("생성 직후 instance: %+v\n", instance)

	//--------------------------------------------------------------------------
	// Setter method를 이용하여 Field 채우기
	// 필드를 이름으로 검색, 필드의 타입에 따라 분기 처리
	// Value 필드에 대해 string 값을 int로 파싱하여 SetInt 한다.
	//--------------------------------------------------------------------------
	for _, field := range structFields {
		switch instance.FieldByName(field.Name).Kind() {
		case reflect.String: // Most cases
			instance.FieldByName(field.Name).SetString(fieldValues[field.Name].(string))
		case reflect.Int:
			valueParsedToInt, err := strconv.Atoi(fieldValues[field.Name].(string))
			if err != nil {
				// TODO Exception Handling
			}
			instance.FieldByName(field.Name).SetInt(int64(valueParsedToInt))
		default:
			fmt.Printf("@main():Exception: Field 중에 String, Int 이외의 Type이 존재")
		}
	}
	fmt.Printf("Set 직후 instance: %+v\n", instance)
	fmt.Printf("instance.FieldByName(Id).Type(): %v\n", instance.FieldByName("Id").Type())
	fmt.Printf("instance.FieldByName(Id).Int(): %v\n", instance.FieldByName("Id").Int())
	fmt.Printf("instance.FieldByName(Issuer).Type(): %v\n", instance.FieldByName("Issuer").Type())
	fmt.Printf("instance.FieldByName(Issuer).String(): %v\n", instance.FieldByName("Issuer").String())

	// 아래와 동일
	//	marshaledJson, errByMarshal := json.Marshal((&instance).Interface())
	marshaledJson, errByMarshal := json.Marshal(instance.Addr().Interface())
	if errByMarshal != nil {
		fmt.Printf("@main():Error: %v", errByMarshal)
		return
	}
	fmt.Printf("marshaledJson: %s\n", marshaledJson)

	var reunmarshaledJson interface{}
	json.Unmarshal(marshaledJson, &reunmarshaledJson)
	fmt.Printf("reunmarshaledJson: (%T)%+v\n\n", reunmarshaledJson, reunmarshaledJson)
}
