/**
	根据url获取网页结果，以便解析器处理
 */

package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
	"fmt"
)
var rateLimiter = time.Tick(10*time.Millisecond)
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status code: %s\n",resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	//获取到的网页数据是gbk格式的 所以需要转为utf8的
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}


func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v\n",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes,"")
	return e
}
