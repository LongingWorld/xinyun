package model

type Gauge struct {
	GaugeID        int
	GaugeName      string
	ShowName       string
	Describe       string
	GaugeType      int
	IsRandom       int
	CompletionTime int
	Guidance       string
	Status         int
	TemplateID     int
	CreatedTime    string
	UpdatedTime    string
	DeletedTime    string
}

type Reportsetting struct {
	ReportSetID       int
	GaugeID           int
	HideReportInstr   int
	ReportInstr       string
	HideShowMethod    int
	ShowMethod        int
	HideDismmisionDes int
	DismmisinDes      string
	HideReportCom     int
	ReportCom         string
	HideDimSuggest    int
	HideCliches       int
	Cliches           string
	ReferScore        float64
	ItScoreDesc       string
	GtScoreDesc       string
	CreatedTime       string
	UpdatedTime       string
	DeletedTime       string
}

type Company struct {
	CompanyID       int
	CompanyName     string
	CompanyUserName string
	ComLoginPasswd  string
	UserPhone       string
	TemplateID      int
	ContractID      string
	Remarks         string
	StartTime       string
	EndTime         string
	AdminID         int
	CreatedTime     string
	UpdatedTime     string
	DeletedTime     string
}

type StaffAnswer struct {
	ID           int
	StaffAnsID   int
	CompanyID    int
	GaugeID      int
	StartTime    string
	EndTime      string
	Score        int
	StaffID      int
	IsFinish     int
	CompanyTimes int
	CreatedTime  string
	UpdatedTime  string
	DeletedTime  string
}

type Subject struct {
	ID          int
	GaugeID     int
	SubjectName string
	Sort        int
	Number      int
	CreatedTime string
	UpdatedTime string
	DeletedTime string
}

type SubjectAnswer struct {
	ID          int
	SubjectID   int
	OpetionName string
	Image       string
	Fraction    int
	MapValue    string
	Sort        int
	CreatedTime string
	UpdatedTime string
}

type StaffAnswerOpetion struct {
	ID            int
	StaffAnswerID int
	SubjectID     int
	SubjectAnsID  int
	StaffID       int
	CreatedTime   string
	UpdatedTime   string
	DeletedTime   string
}

type ReportStaff struct {
	ID             int
	StaffAnsID     int
	Name           string
	HideRepInstr   int
	RepInstr       string
	HideDesc       int
	ItGtDesc       string
	Describe       string
	HideShowMethod int
	ShowMethod     int
	HideCliches    int
	Cliches        string
	HideComment    int
	Comment        string
	HideDimSuggest int
	GaugeID        int
	StaffID        int
	StaffName      string
	StaffAge       int
	Position       string
	Marriage       int
	CompanyID      int
	CompanyName    string
	Number         string
	Status         int
	TotalScore     float32
	TemplateID     int
	GenerateDate   string
	CreatedTime    string
	UpdatedTime    string
	DeletedTime    string
}
