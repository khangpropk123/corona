package model

type Message struct {
	Messages []Text `json:"messages"`
}

type Text struct {
	Text string `json:"text"`
}

type CoronaResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
	SoCaNhiem string `json:"soCaNhiem"`
	TuVong string `json:"tuVong"`
	NghiNhiem string `json:"nghiNhiem"`
	BinhPhuc string	`json:"binhPhuc"`
	CachLy string `json:"cachLy"`
}
