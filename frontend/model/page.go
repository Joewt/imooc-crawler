package model

type SearchResult struct {
	Hits     int
	Start    int
	Query 	 string
	PrevFrom int
	NextFrom int
	Items []interface{}
}
