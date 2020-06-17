package main

import (
	"encoding/json"
	"fmt"
)

type Storage struct {
	Collection string
	Id         string
	Type       string
	Name       string
	Data       string
}

func Test() string {
	files := make([]*Storage, 0)
	file := &Storage{
		Collection: "collection",
		Id:         "111111",
		Type:       "file",
		Name:       "file1",
		Data:       "data",
	}
	files = append(files, file)

	file2 := &Storage{
		Collection: "collection",
		Id:         "22222",
		Type:       "file",
		Name:       "file2",
		Data:       "data",
	}
	files = append(files, file2)

	data, err := json.Marshal(files)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))

	return string(data)
}

func Test2() {
	var files []*Storage

	data, err := json.Marshal(files)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))
}

func Test3() {
	var files []*Storage
	files = nil

	data, err := json.Marshal(files)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))
}

func Test4() {
	files := make([]*Storage, 0)

	data, err := json.Marshal(files)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))
}

func Test5(data string) []*Storage {
	if data == "[]" {
		return nil
	}

	files := make([]*Storage, 0)
	body := []byte(data)
	err := json.Unmarshal(body, &files)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Println("===files===", files)
	fmt.Println("===files===", *files[0])
	return files
}

func main() {
	data := Test()
	Test2()
	Test3()
	Test4()
	Test5(data)

}
