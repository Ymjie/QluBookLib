package model

type Bookresp struct {
	Status    int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      BookData    `json:"data"`
	Checkinfo interface{} `json:"checkinfo"`
}
type BeginTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
type EndTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
type UpdateTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
type ExamTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
type ParentAreaInfo struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	ParentID    int         `json:"parentId"`
	Levels      int         `json:"levels"`
	IsValid     int         `json:"isValid"`
	Comment     string      `json:"comment"`
	Sort        int         `json:"sort"`
	Type        int         `json:"type"`
	Color       interface{} `json:"color"`
	Enname      string      `json:"enname"`
	NameMerge   string      `json:"nameMerge"`
	EnnameMerge string      `json:"ennameMerge"`
	ROWNUMBER   string      `json:"ROW_NUMBER"`
}
type AreaInfo struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	ParentID       int              `json:"parentId"`
	Levels         int              `json:"levels"`
	IsValid        int              `json:"isValid"`
	Comment        string           `json:"comment"`
	Sort           int              `json:"sort"`
	Type           int              `json:"type"`
	Color          interface{}      `json:"color"`
	Enname         string           `json:"enname"`
	NameMerge      string           `json:"nameMerge"`
	EnnameMerge    string           `json:"ennameMerge"`
	ROWNUMBER      string           `json:"ROW_NUMBER"`
	ParentAreaInfo []ParentAreaInfo `json:"parentAreaInfo"`
}
type SpaceOpenTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
type SpaceCloseTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
type BookRuleInfo struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	NeedExam        int            `json:"needExam"`
	ReserveTime     int            `json:"reserveTime"`
	AttentionTime   int            `json:"attentionTime"`
	AlarmTime       int            `json:"alarmTime"`
	OpenTime        int            `json:"openTime"`
	CloseTime       int            `json:"closeTime"`
	MinTime         int            `json:"minTime"`
	MaxTime         int            `json:"maxTime"`
	ContinueTime    int            `json:"continueTime"`
	CancelTime      int            `json:"cancelTime"`
	LeaveTime       int            `json:"leaveTime"`
	BookDay         int            `json:"bookDay"`
	UpdateTime      int            `json:"updateTime"`
	SignIn          int            `json:"signIn"`
	SignInPerson    int            `json:"signInPerson"`
	SuperSignIn     int            `json:"superSignIn"`
	SignOut         int            `json:"signOut"`
	SignOutDelay    int            `json:"signOutDelay"`
	Light           int            `json:"light"`
	Power           int            `json:"power"`
	AutoSignOutTime int            `json:"autoSignOutTime"`
	SpaceOpenTime   SpaceOpenTime  `json:"spaceOpenTime"`
	SpaceCloseTime  SpaceCloseTime `json:"spaceCloseTime"`
	UpdateCount     int            `json:"updateCount"`
	ROWNUMBER       string         `json:"ROW_NUMBER"`
}
type RenegeRuleInfo struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	SignInCount       int    `json:"signInCount"`
	SignOutCount      int    `json:"signOutCount"`
	LeaveNoBackCount  int    `json:"leaveNoBackCount"`
	LeaveNoCheckCount int    `json:"leaveNoCheckCount"`
	LateCount         int    `json:"lateCount"`
	TotalCount        int    `json:"totalCount"`
	SignInHour        int    `json:"signInHour"`
	SignOutHour       int    `json:"signOutHour"`
	LeaveNoBackHour   int    `json:"leaveNoBackHour"`
	LeaveNoCheckHour  int    `json:"leaveNoCheckHour"`
	LateHour          int    `json:"lateHour"`
	TotalHour         int    `json:"totalHour"`
	IsValid           int    `json:"isValid"`
	ROWNUMBER         string `json:"ROW_NUMBER"`
}
type CategoryInfo struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Type           int            `json:"type"`
	SpaceCount     int            `json:"spaceCount"`
	MinPerson      int            `json:"minPerson"`
	MaxPerson      int            `json:"maxPerson"`
	BookRule       int            `json:"bookRule"`
	RenegeRule     int            `json:"renegeRule"`
	Comment        string         `json:"comment"`
	IsValid        int            `json:"isValid"`
	ROWNUMBER      string         `json:"ROW_NUMBER"`
	BookRuleInfo   BookRuleInfo   `json:"bookRuleInfo"`
	RenegeRuleInfo RenegeRuleInfo `json:"renegeRuleInfo"`
}
type SpaceInfo struct {
	ID           int          `json:"id"`
	No           string       `json:"no"`
	Name         string       `json:"name"`
	Area         int          `json:"area"`
	Category     int          `json:"category"`
	Status       int          `json:"status"`
	Detail       string       `json:"detail"`
	ROWNUMBER    string       `json:"ROW_NUMBER"`
	AreaInfo     AreaInfo     `json:"areaInfo"`
	CategoryInfo CategoryInfo `json:"categoryInfo"`
}
type Renegeinfo struct {
	Renege int         `json:"renege"`
	Count  interface{} `json:"count"`
}
type BookList struct {
	ID            int         `json:"id"`
	No            string      `json:"no"`
	Booker        string      `json:"booker"`
	SpaceCategory int         `json:"spaceCategory"`
	Space         string      `json:"space"`
	IsSingle      int         `json:"isSingle"`
	MemberCount   int         `json:"memberCount"`
	BeginTime     BeginTime   `json:"beginTime"`
	EndTime       EndTime     `json:"endTime"`
	Title         string      `json:"title"`
	Application   string      `json:"application"`
	IsPublic      int         `json:"isPublic"`
	UpdateTime    UpdateTime  `json:"updateTime"`
	ExamTime      ExamTime    `json:"examTime"`
	Examinant     interface{} `json:"examinant"`
	ExamResult    string      `json:"examResult"`
	SignIn        int         `json:"signIn"`
	SignOut       int         `json:"signOut"`
	Status        int         `json:"status"`
	ROWNUMBER     string      `json:"ROW_NUMBER"`
	SpaceInfo     SpaceInfo   `json:"spaceInfo"`
	StatusName    string      `json:"statusName"`
	Renegeinfo    Renegeinfo  `json:"renegeinfo"`
	Starttime     string      `json:"starttime"`
	Endingtime    string      `json:"endingtime"`
	Segment       string      `json:"segment"`
	Linkurl       string      `json:"linkurl"`
	PData         string      `json:"p_data"`
}

type BookData struct {
	List BookList `json:"list"`
	Hash Hash     `json:"_hash_"`
}
