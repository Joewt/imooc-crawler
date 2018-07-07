package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}


type ParserResult struct {
	Requests []Request
	Items []Item
}


type Item struct {
	Url 	string
	Type    string
	Id  	string
	Payload interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}





