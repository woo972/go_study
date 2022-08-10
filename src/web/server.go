package main

import (
	"net/http"
)

type Server struct {
	*router
	middlewares  []MiddleWare
	startHandler HandlerFunc
}

func NewServer() *Server {
	return &Server{
		router: &router{make(map[string]map[string]HandlerFunc)},
		middlewares: []MiddleWare{
			logHandler,
			recoverHandler,
			staticHandler,
			parseFormHandler,
			parseJsonBodyHandler,
		}}
}

func (s *Server) Run(addr string) {
	// startHandler를 라우터 핸들러 함수로 지정
	s.startHandler = s.router.handler()

	// 등록된 미들웨어를 라우터 핸들러 앞에 하나씩 추가
	// middleware를 선언한 역순으로 적용하기 위함
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		s.startHandler = s.middlewares[i](s.startHandler)
	}

	// 웹 서버 시작
	if err := http.ListenAndServe(addr, s); err != nil {
		panic(err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		Params:         make(map[string]interface{}),
		ResponseWriter: w,
		Request:        r,
	}
	for k, v := range r.URL.Query() {
		c.Params[k] = v[0]
	}
	s.startHandler(c)
}

func (s *Server) Use(middlewares ...MiddleWare) {
	s.middlewares = append(s.middlewares, middlewares...)
}
