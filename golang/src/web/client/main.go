package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

func main() {
	//Get()
	//PostForm()
	uploadFile()
}

/*
uploadFile 以form表单的形式进行文件上传；
*/
func uploadFile() {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("context", "test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = part.Write([]byte("test"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Warning  这里的Close有 flush的效果，所以一定要在这里关闭，不能在defer中关闭
	if err := writer.Close(); err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:8081/fs/upClient", writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Get() {
	_, err := http.Get("http://localhost:8081/test?name=cloudzp")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func PostForm() {
	_, err := http.Post("http://localhost:8081/test", "application/x-www-form-urlencoded",
		strings.NewReader("name=cloudzp"))
	if err != nil {
		fmt.Println("failed to do post ,due to : ", err)
		return
	}
}
