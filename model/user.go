package model

type UserInfo struct {
	Name string
	Age  int
}

type UserInfoReq struct {
	ID int `json:"id"`
}

type UserInfoResp struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
