package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/sundayfun/go-web/util/filter"
)

func ImageFromImageUrl(imageUrl string, fileName string) error {
	resp, err := http.Get(imageUrl)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	file, err := os.Create("../resource/picture/" + fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	io.Copy(file, bytes.NewReader(body))
	return nil
}

func ImageUrlFromHtml(html string) {
	str := filter.ReImageURL.FindAllString(html, -1)
	for i, val := range str {
		fmt.Println(i, val)
		err := ImageFromImageUrl(val, strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func HtmlFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

const (
	// URL = "http://image.baidu.com/search/index?tn=baiduimage&ipn=r&ct=201326592&cl=2&lm=-1&st=-1&fm=result&fr=&sf=1&fmq=1525178045789_R&pv=&ic=0&nc=1&z=&se=1&showtab=0&fb=0&width=&height=&face=0&istype=2&ie=utf-8&hs=2&word=baby"
	// URL = "http://www.tooopen.com/img"
	URL = "https://www.zhihu.com/question/66313867/answer/242883999"
)

func done() {
	html, err := HtmlFromUrl(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("../resource/haha.html", []byte(html), 0664)
	ImageUrlFromHtml(html)
}

func main() {
	done()
	// ImageFromImageUrl("http://img5.imgtn.bdimg.com/it/u=1011441765,2690175992&fm=27&gp=0.jpg", "vimi")
	return
}
