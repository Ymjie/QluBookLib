package model

type BookTimeIdresp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Datas  `json:"data"`
}

type Lists struct {
	SpaceID    int       `json:"spaceId"`
	SpaceName  string    `json:"spaceName"`
	Area       int       `json:"area"`
	BookTimeID int       `json:"bookTimeId"`
	BeginTime  BeginTime `json:"beginTime"`
	EndTime    string    `json:"endTime"`
	Status     int       `json:"status"`
	Day        string    `json:"day"`
	StartTime  string    `json:"startTime"`
	ID         int       `json:"id"`
}
type Datas struct {
	List []Lists `json:"list"`
}
