// Code generated by goctl. DO NOT EDIT.
package types

type PicItem struct {
	ID     int    `json:"id"`
	Pid    string `json:"pid"`
	ItemId string `json:"itemId"`
	Text   string `json:"text"`
	Src    string `json:"src"`
}

type PicList struct {
	Page  int       `json:"page"`
	Total int       `json:"total"`
	List  []PicItem `json:"list"`
}

type PicsReq struct {
	Page int `path:"page"`
}

type PicsRes struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data PicList `json:"data"`
}

type StarReq struct {
	ID int `path:"id"`
}

type StarRes struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []PicItem `json:"data"`
}

type DetailReq struct {
	ID int `path:"id"`
}

type DetailRes struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data PicItem `json:"data"`
}
