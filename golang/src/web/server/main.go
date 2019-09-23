package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

/*

测试表单上传数据的服务端代码

*/

func main() {
	http.HandleFunc("/test", func(r http.ResponseWriter, q *http.Request) {
		//fmt.Println(fmt.Sprintf("data %+v ", q.Body))
		/*
		 方式一：
		 用来获取通过Get方法提交的数据，数据会被带在url中，如：'?name=zhangsan'
		*/
		// Note: get 请求需要调用这个方法；
		q.ParseForm()
		fmt.Println("get:  ", q.Form.Get("name"))

		/*
		 方式二：
		  用来获取post提交的数据，数据会被解析到body体中
		*/
		// postFormValue可以在普通的form中和同时传文件及数据的表单中传数据；
		fmt.Println("post:   ", q.PostFormValue("name"))
	})

	// 文件上传
	// 问题1： This file resides outside the working directory. Collaborators might not have the same file path.
	//  这个不是问题；
	http.HandleFunc("/fs/upClient", func(r http.ResponseWriter, q *http.Request) {

		// 默认值内存够的情况下 不需要指定
		/*if err := q.ParseMultipartForm(32 << 50); err != nil {
			fmt.Println("parse form error:", err)
			return
		}*/

		fmt.Println(q.PostFormValue("module"))

		file, head, err := q.FormFile("context")
		if err != nil {
			fmt.Println("form file error: ", err)
			return
		}

		defer file.Close()

		f, err := os.Create(filepath.Join("/tmp/", head.Filename))
		if err != nil {
			fmt.Println("create error :", err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		fmt.Println("success to upload file . ")
	})

	// TODO delete
	http.HandleFunc("/api/inout/data", func(writer http.ResponseWriter, request *http.Request) {
		data := request.PostFormValue("data")
		fmt.Println(data)

		status := request.PostFormValue("status")
		fmt.Println(status)

		writer.Write([]byte(`{"message":"ok"}`))
	})

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		panic(err)
	}
}
