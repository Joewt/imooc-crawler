package parser

import (
	"joewt.com/joe/learngo/crawler/engine"
	"regexp"
)

var  (
	profileRe =  regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe  = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)
//解析城市数据
func ParserCity(contents []byte) engine.ParserResult {
	matches := profileRe.FindAllSubmatch(contents,-1)
	result := engine.ParserResult{}
	for _, c := range matches {
		url  := string(c[1])
		name := string(c[2])
		result.Requests = append(result.Requests,engine.Request{
			Url:		url,
			ParserFunc:	func(c []byte) engine.ParserResult{
				return ParserProfile(c,url,name)
			},
		})
		//fmt.Printf("User: %s, URL: %s\n",c[2],c[1])
	}

	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
				engine.Request{
					Url:  string(m[1]),
					ParserFunc:ParserCity,
				},
			)
	}



	return result
}
