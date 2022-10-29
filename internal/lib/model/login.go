package model

type Loginresp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}
type List struct {
	ID                string      `json:"id"`
	Card              string      `json:"card"`
	Name              string      `json:"name"`
	IDCard            string      `json:"idCard"`
	Gender            int         `json:"gender"`
	Birthday          string      `json:"birthday"`
	JoinTime          string      `json:"joinTime"`
	Wallet            string      `json:"wallet"`
	Saving            string      `json:"saving"`
	FillScore         int         `json:"fillScore"`
	TotalFillScore    int         `json:"totalFillScore"`
	ConsumeScore      int         `json:"consumeScore"`
	TotalConsumeScore int         `json:"totalConsumeScore"`
	Role              interface{} `json:"role"`
	RoleName          string      `json:"roleName"`
	Dept              interface{} `json:"dept"`
	DeptName          string      `json:"deptName"`
	SubDept           interface{} `json:"subDept"`
	SubDeptName       interface{} `json:"subDeptName"`
	Tel               string      `json:"tel"`
	Mobile            string      `json:"mobile"`
	Email             string      `json:"email"`
	Qq                string      `json:"qq"`
	Status            int         `json:"status"`
	Weixin            string      `json:"weixin"`
	HwUpdateFlag      int         `json:"hw_update_flag"`
	SkedbUpdateFlag   int         `json:"skedb_update_flag"`
	ROWNUMBER         string      `json:"ROW_NUMBER"`
	Renegeinfo        interface{} `json:"renegeinfo"`
}
type Hash struct {
	Userid      string `json:"userid"`
	AccessToken string `json:"access_token"`
	Expire      string `json:"expire"`
}
type Data struct {
	List List `json:"list"`
	Hash Hash `json:"_hash_"`
}
