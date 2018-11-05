package spider

import (
	"net/http"
	"fmt"
	"net/http/httputil"
)

type MatchReg func() int

type Spider struct {
	RawUrl      string
	Url         string   // 页面地址
	MatchFunc   MatchReg // 正则表达式
	PageContent string   // 页面内容
}

func (s *Spider) GetPage(url string) error {
	if s.Url == "" {
		panic("invalid url")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

/**
处理正则
 */
func (s *Spider) ProcessPage() error {
	s.MatchFunc()
	return nil
}
