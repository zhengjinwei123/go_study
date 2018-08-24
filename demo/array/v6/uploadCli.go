package main

import (
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"path/filepath"
)

func postFile(filename string,targetUrl string) error{
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter,err := bodyWriter.CreateFormFile("uploadfile","README.md")
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	fh,err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	defer fh.Close()

	_,err = io.Copy(fileWriter,fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp,err := http.Post(targetUrl,contentType,bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("resp.status:",resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main(){
	targetUrl := "http://localhost:9090/upload"
	fileName,_ := filepath.Abs("./")
	fileName += "/README.md"
	err := postFile(fileName,targetUrl)
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("upload file success")
}
