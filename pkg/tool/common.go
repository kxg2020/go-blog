package tool

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// md5
func Md5Encrypt(str string)string  {
	ctx := md5.New()
	ctx.Write([]byte(str))
	ret := ctx.Sum(nil)
	return hex.EncodeToString(ret)
}

//sha1
func Sha1Encrypt(str string)string  {
	ctx := sha1.New()
	ctx.Write([]byte(str))
	ret := ctx.Sum(nil)
	return hex.EncodeToString(ret)
}

// dir or file exist
func DirOrFileExist(path string)(bool,error)  {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false,err
	}
	return false, err
}

// http
func HttpRequest(url string,method string,data map[string]string)(res string)  {
	client  := &http.Client{}
	var postStr string
	for key,val := range data{
		postStr += (key + "=" + val) + "&"
	}
	req,err := http.NewRequest(method, url, strings.NewReader(postStr))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		resp.Body.Close()
		res = ""
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body)
}

// http.Get
func HttpGet(url string)string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

// http.Post
func HttpPost(link string,data map[string]string)string {
	var postStr string
	for key,val := range data{
		postStr += (key + "=" + val) + "&"
	}
	resp, err := http.Post(link, "application/x-www-form-urlencode", strings.NewReader(postStr))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HttpPostJson(link string,data []byte) []byte {
	request, err := http.NewRequest("POST", link, bytes.NewReader(data))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return respBytes
}