package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

/* type postRequst struct {
	accessToken     string
	gaugeID         string
	subjectsAnswers string
}
*/

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

func main() {
	router := gin.Default()
	router.POST("/first", PostHandler)

	http.ListenAndServe(":8000", router)
}

func PostHandler(c *gin.Context) {
	/*连接数据库*/
	db, err := gorm.Open("mysql", "root:zzj210@tcp(127.0.0.1:3308)/world?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	/*关闭连接数据库*/
	defer db.Close()

	var gauge Gauge
	var company Company
	var reportSet Reportsetting

	/*获取量表ID,员工ID,员工所属企业名称*/
	gaugeID := c.PostForm("gauge_id")
	//db.First(&gauge, "id = ?", gaugeID)
	gint, _ := strconv.Atoi(gaugeID)
	db.Where("id = ?", gint).Find(&gauge)
	if gauge.GaugeName == "" {
		log.Printf("量表%s不存在", gaugeID)
		return
		//fmt.Errorf("量表%s不存在", gaugeID)
	}

	db.Where("id = ?")

	subjectsAnswers := c.PostForm("subjectsanswers")
	subjectsAnswersArr, errs := json.Marshal(subjectsAnswers)
	if errs != nil {
		log.Println("解析JSON字符串错！")
		return
	}

	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, subjectsAnswers)
	return
}
