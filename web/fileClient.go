package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// url := "http://127.0.0.1:8080/upload"
	// filename := "test"
	// if err := uploadFile(filename, url); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("upload ok")

	if err := downloadFile("test1", "http://127.0.0.1:8080/download/"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("download ok")
}

func downloadFile(filename string, url string) error {
	url = url + filename

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)

	lfile, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	io.Copy(lfile, resp.Body)
	return nil
}

func uploadFile(filename string, url string) error {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	fWriter, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		return err
	}

	lfile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer lfile.Close()

	_, err = io.Copy(fWriter, lfile)
	if err != nil {
		return err
	}

	contentType := writer.FormDataContentType()
	writer.Close()

	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil

}
