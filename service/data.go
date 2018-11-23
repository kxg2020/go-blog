package service

type departmentDetail struct{
	Id       	 int    			`json:"id"`
	Name     	 string 			`json:"name"`
	ParentId 	 int    			`json:"parentid"`
	Order        int    			`json:"order"`
}

type DepartmentList struct{
	ErrCode      int    			`json:"errcode"`
	ErrMsg       string 			`json:"errmsg"`
	Department   []departmentDetail `json:"department"`
}


type Token struct{
	ErrCode      int    			`json:"errcode"`
	ErrMsg       string 			`json:"errmsg"`
	Access_token string 			`json:"access_token"`
	Expires_in   int    			`json:"expires_in"`
}
