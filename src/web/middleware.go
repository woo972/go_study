package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

type MiddleWare func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		requestStartTime := time.Now()

		next(ctx)

		requestEndTime := time.Now().Sub(requestStartTime)

		log.Printf("[%s] %q elapsed: %v\n",
			ctx.Request.Method,
			ctx.Request.URL.String(),
			requestEndTime)
	}
}

func recoverHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() { // defer = finally
			if err := recover(); err != nil {
				http.Error(c.ResponseWriter,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next(c)
	}
}

func parseFormHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		c.Request.ParseForm()
		fmt.Println(c.Request.PostForm)
		for k, v := range c.Request.PostForm {
			fmt.Printf("parseform key: %s / value: %s\n", k, v)
			if len(v) > 0 {
				c.Params[k] = v[0]
			}
		}
		next(c)
	}
}

func parseJsonBodyHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		var m map[string]interface{}
		fmt.Printf("request body: %v\n", c.Request.Body)
		if json.NewDecoder(c.Request.Body).Decode(&m); len(m) > 0 {
			for k, v := range m {
				fmt.Printf("parsejsonbody key: %s / value: %s\n", k, v)
				c.Params[k] = v
			}
		}
		next(c)
	}
}

func staticHandler(next HandlerFunc) HandlerFunc {
	var (
		dir       = http.Dir(".")
		indexFile = "index.html"
	)
	return func(c *Context) {
		// http 메서드가 GET이나 HEAD가 아니면 바로 다음 핸들러 수행
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			next(c)
			return
		}

		file := c.Request.URL.Path
		// URL 경로에 해당하는 파일 열기 시도
		f, err := dir.Open(file)
		if err != nil {
			// URL 경로에 해당하는 파일 열기에 실패하면 바로 다음 핸들러 수행
			next(c)
			return
		}
		defer f.Close()

		fileInfo, err := f.Stat()
		if err != nil {
			// 파일의 상태가 정상이 아니면 바로 다음 핸들러 수행
			next(c)
			return
		}

		// URL 경로가 디렉터리면 indexFile을 사용
		if fileInfo.IsDir() {
			// 디렉터리 경로를 URL로 사용하면 경로 끝에 “/“를 붙여야 함
			if !strings.HasSuffix(c.Request.URL.Path, "/") {
				http.Redirect(c.ResponseWriter, c.Request, c.Request.URL.Path+"/", http.StatusFound)
				return
			}

			// 디렉터리를 가리키는 URL 경로에 indexFile 이름을 붙여서 전체 파일 경로 생성
			file = path.Join(file, indexFile)

			// indexFile 열기 시도
			f, err = dir.Open(file)
			if err != nil {
				next(c)
				return
			}
			defer f.Close()

			fileInfo, err = f.Stat()
			if err != nil || fileInfo.IsDir() {
				// indexFile 상태가 정상이 아니면 바로 다음 핸들러 수행
				next(c)
				return
			}
		}

		// file의 내용 전달(next 핸들러로 제어권을 넘기지 않고 요청 처리를 종료함)
		http.ServeContent(c.ResponseWriter, c.Request, file, fileInfo.ModTime(), f)
	}
}
