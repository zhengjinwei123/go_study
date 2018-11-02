package spider

import (
	"net/http"
	"fmt"
	"net/http/httputil"
)

type Spider struct {
	Url         string // 页面地址
	Reg         string // 正则表达式
	PageContent string // 页面内容
	ResultList  []string // 结果列表
}

func (s *Spider) GetPage() error {
	if s.Url == "" {
		panic("invalid url")
	}

	req, err := http.NewRequest(http.MethodGet, s.Url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	s.PageContent = string(content)
	return nil
}

func (s *Spider) ProcessPage() error {
	if s.Reg == "" {
		panic("invalid reg")
	}

	return nil
}
