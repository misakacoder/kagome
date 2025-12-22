package http

import (
	"fmt"
	"github.com/misakacoder/kagome/cond"
	"github.com/misakacoder/kagome/json"
	"github.com/misakacoder/kagome/str"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	Client = &http.Client{
		Timeout: 30 * time.Second,
	}
	Logger func(format string, v ...interface{}) = nil
)

type MultipartFile string

func Get[T any](url string, headers map[string]string) (T, error) {
	return Execute[T](url, http.MethodGet, headers, nil)
}

func PostForm[T any](url string, headers map[string]string, data url.Values) (T, error) {
	headers = cond.RequireNonNilElse(headers, make(map[string]string, 1))
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return Execute[T](url, http.MethodPost, headers, strings.NewReader(data.Encode()))
}

func PostJSON[T any](url string, headers map[string]string, data string) (T, error) {
	headers = cond.RequireNonNilElse(headers, make(map[string]string, 1))
	headers["Content-Type"] = "application/json"
	return Execute[T](url, http.MethodPost, headers, strings.NewReader(data))
}

func PostFile[T any](url string, headers map[string]string, data map[string]any) (T, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	go func() {
		defer pw.Close()
		defer writer.Close()
		for k, v := range data {
			if multipartFile, ok := v.(MultipartFile); ok {
				path := string(multipartFile)
				filename := filepath.Base(path)
				part, err := writer.CreateFormFile(k, filename)
				if err != nil {
					logger("create form file error: %v", err)
					return
				}
				file, err := os.Open(path)
				if err != nil {
					logger("open file error: %v", err)
					return
				}
				_, err = io.Copy(part, file)
				file.Close()
				if err != nil {
					logger("copy file error: %v", err)
					return
				}
			} else {
				writer.WriteField(k, fmt.Sprintf("%v", v))
			}
		}
	}()
	headers = cond.RequireNonNilElse(headers, make(map[string]string, 1))
	headers["Content-Type"] = writer.FormDataContentType()
	return Execute[T](url, http.MethodPost, headers, pr)
}

func Execute[T any](url string, method string, headers map[string]string, data io.Reader) (T, error) {
	var param string
	if Logger != nil {
		reader, ok := data.(*strings.Reader)
		if ok {
			bytes, _ := io.ReadAll(reader)
			param = string(bytes)
			data = strings.NewReader(param)
		}
	}
	var result T
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return result, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	response, err := Client.Do(request)
	if err != nil {
		return result, err
	}
	body := response.Body
	defer body.Close()
	bytes, err := io.ReadAll(body)
	if err != nil {
		return result, err
	}
	responseString := string(bytes)
	if Logger != nil {
		joiner := str.NewJoiner("\n", "\n", "")
		joiner.Append(fmt.Sprintf("Request Url: %s", url))
		joiner.Append("Request Headers: ")
		for k, v := range request.Header {
			joiner.Append(fmt.Sprintf("    %s: %s", k, strings.Join(v, ",")))
		}
		joiner.Append("Request Body: ")
		joiner.Append(fmt.Sprintf("    %s", param))
		joiner.Append("Response Headers: ")
		for k, v := range response.Header {
			joiner.Append(fmt.Sprintf("    %s: %s", k, strings.Join(v, ",")))
		}
		joiner.Append("Response Body: ")
		joiner.Append(fmt.Sprintf("    %s", responseString))
		Logger("%s", joiner.String())
	}
	return json.ParseObject[T](responseString)
}

func logger(format string, args ...any) {
	if Logger != nil {
		Logger(format, args...)
	}
}
