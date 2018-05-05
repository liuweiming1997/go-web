package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/util/file"
	"github.com/sundayfun/go-web/util/filter"
)

const (
	SavePicturePath = "./resource/picture/"

	// URL = "http://image.baidu.com/search/index?tn=baiduimage&ipn=r&ct=201326592&cl=2&lm=-1&st=-1&fm=result&fr=&sf=1&fmq=1525178045789_R&pv=&ic=0&nc=1&z=&se=1&showtab=0&fb=0&width=&height=&face=0&istype=2&ie=utf-8&hs=2&word=baby"
	// URL = "http://www.tooopen.com/img"
	URL = "https://image.baidu.com/search/index?tn=baiduimage&ipn=r&ct=201326592&cl=2&lm=-1&st=-1&fm=index&fr=&hs=0&xthttps=111111&sf=1&fmq=&pv=&ic=0&nc=1&z=&se=1&showtab=0&fb=0&width=&height=&face=0&istype=2&ie=utf-8&word=%E7%99%BD%E7%9F%B3%E8%8C%89%E8%8E%89%E5%A5%88%E5%A5%88&oq=%E7%99%BD%E7%9F%B3%E8%8C%89%E8%8E%89&rsp=0"
)

var (
	same = make(map[string]bool)
)

func ImageFromImageUrl(imageUrl string, fileName string) error {
	if same[imageUrl] {
		return fmt.Errorf("exist")
	}
	// if _, ok := same[imageUrl]; ok {
	// 	return fmt.Errorf("exist")
	// }

	same[imageUrl] = true
	resp, err := http.Get(imageUrl)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	file, err := file.NewFileWriter(SavePicturePath+fileName, false)
	defer file.Close()
	if err != nil {
		return err
	}
	err = file.WriteByte(body)
	return err
	// file, err := os.Create(SavePicturePath + fileName)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }
	// io.Copy(file, bytes.NewReader(body))
	// return nil
}

func ImageUrlFromHtml(html string) {
	str := filter.ReImageURL.FindAllString(html, -1)
	for i, val := range str {
		fmt.Println(i, val)
		err := ImageFromImageUrl(val, strconv.Itoa(i))
		if err != nil {
			logrus.Error(err)
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

func done() {
	html, err := HtmlFromUrl(URL)
	if err != nil {
		logrus.Error(err)
		return
	}
	ioutil.WriteFile("./resource/haha.html", []byte(html), 0664)
	ImageUrlFromHtml(html)
}

func main() {
	err := file.CreateDirIfNotExit(SavePicturePath)
	if err != nil {
		logrus.Error(err)

	}
	done()
	// ImageFromImageUrl("http://img5.imgtn.bdimg.com/it/u=1011441765,2690175992&fm=27&gp=0.jpg", "vimi")
	return
}
