package parser

import (
	"joewt.com/joe/learngo/crawler/engine"
	"regexp"
	"strconv"
	"joewt.com/joe/learngo/crawler/model"
)

var ageRe 		 = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe 	 = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe	 = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe	 = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xingzuoRe    = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var hukouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe		 = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe      = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe      = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

//个人信息页数据的解析
func ParserProfile(contents []byte,url string, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents,heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents,weightRe))
	if err == nil {
		profile.Weight = weight
	}
	//profile.Age         = extractString(contents,ageRe)
	profile.Marriage 	= extractString(contents,marriageRe)
	profile.Income   	= extractString(contents,incomeRe)
	profile.Gender   	= extractString(contents,genderRe)
	profile.Xingzuo  	= extractString(contents,xingzuoRe)
	profile.Education   = extractString(contents,educationRe)
	profile.Occupation  = extractString(contents,occupationRe)
	profile.Hukou       = extractString(contents,hukouRe)
	profile.House       = extractString(contents,houseRe)
	profile.Car			= extractString(contents,carRe)

	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url: 		url,
				Type: 		"zhenai",
				Id: 		extractString([]byte(url),idUrlRe),
				Payload: 	profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents,-1)

	for _, m := range matches {
		name := string(m[2])
		url  := string(m[1])
		result.Requests = append(result.Requests,
				engine.Request{
					Url: string(m[1]),
					ParserFunc: ProfileParser(name, url),
				},
			)
	}
	return result

}

func ProfileParser(name string, url string) engine.ParserFunc {
	return func(c []byte) engine.ParserResult{
		return ParserProfile(c, url, name)
	}
}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}