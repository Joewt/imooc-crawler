package parser

import (
	"joewt.com/joe/learngo/crawler/engine"
	"regexp"
)

const cityListRe =  `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
//所有城市解析
func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)
	result := engine.ParserResult{}

	for _, c := range matches {
		//result.Items = append(result.Items,"City: "+string(c[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:		string(c[1]),
			ParserFunc:	ParserCity,
		})
		//fmt.Printf("City: %s, URL: %s\n",c[2],c[1])
	}
	return result
}