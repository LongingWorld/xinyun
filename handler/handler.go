package handler

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"xinyun/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func PostHandler(c *gin.Context) {
	fmt.Println("@@@@@@@Start@@@@@@@")
	/*连接数据库*/
	// db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxj2018?charset=utf8&parseTime=true&loc=Local")
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxj2018?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	/*关闭连接数据库*/
	defer db.Close()

	var gauge model.Gauge
	// var company Company
	//var reportSet model.Reportsetting

	/*获取量表ID*/
	gaugeID := c.PostForm("gauge_id")
	//db.First(&gauge, "id = ?", gaugeID)
	gint, _ := strconv.Atoi(gaugeID)
	fmt.Printf("*****gauge_id is %d****\n", gint)
	db.Table("xy_gauge").Where("id = ?", gint).Find(&gauge)
	fmt.Println(gauge)
	var count int
	var name string
	//db.Model(&model.Gauge{}).Where("id = ?", gint).Count(&count) //获取记录数
	db.Table("xy_gauge").Select("xy_gauge.id").Where("id = ?", gint).Count(&count)
	db.Table("xy_gauge").Select("xy_gauge.name").Where("id = ?", gint).First(&name)
	fmt.Printf("*****count is %d name is %s****\n", count, gauge.GaugeName)
	if count == 0 {
		log.Printf("量表%s不存在", gaugeID)
		return
		//fmt.Errorf("量表%s不存在", gaugeID)
	}

	c.JSON(http.StatusOK, gauge)
	return

	// /*获取员工答题ID*/
	// staffAnsID := c.PostForm("service_use_staff_id")
	// staAnsID, _ := strconv.Atoi(staffAnsID)

	// /*获取答题答案JSON字符串*/
	// subjectsAnswers := c.PostForm("subjectsanswers")

	// /*解析JSON字符串存入map中*/
	// var subjectsAnswersArr = make(map[string]string)
	// errs := json.Unmarshal([]byte(subjectsAnswers), &subjectsAnswersArr)
	// if errs != nil {
	// 	log.Println("解析JSON字符串错！")
	// 	return
	// }
	// if len(subjectsAnswersArr) == 0 {
	// 	log.Println("请答题！")
	// }
	// fmt.Println(subjectsAnswersArr)

	// /*读取redis获取token,通过token匹配用户信息BEGIN*/
	// /*读取redis获取token,通过token匹配用户信息END*/

	// /*获取xy_reportsetting报告设置表信息*/
	// db.Where("id = ?", gaugeID).Find(&reportSet)

	// var staffAnswer []model.StaffAnswer

	// /*获取员工答题信息*/
	// db.Where("service_use_staff_id = ? AND staff_id = ? AND is_finish = ? AND deleted_at = ?", staAnsID, 123, 2, "").Find(&staffAnswer)
	// if len(staffAnswer) == 0 {
	// 	log.Printf("员工%d答题信息表记录不存在！", 123)
	// 	return
	// }

	// /*获取员工答题题目列表*/
	// subjectIDs := getKeysFromMap(subjectsAnswersArr)
	// if len(subjectIDs) == 0 {
	// 	log.Println("请答完所有题目！")
	// 	return
	// }

	// /*将取出的题目列表subjectIDs与量表对应的题目列表做比较，量表对应的所有题目存在于subjectIDs中*/
	// var subCount int
	// db.Model(&model.Subject{}).Where("gauge_id = ?", gaugeID).Not("id", subjectIDs).Count(&subCount)
	// if subCount > 0 {
	// 	log.Println("请答完所有题目！")
	// 	return
	// }

	// /*开启事务*/
	// db.Begin()
	// /*更新staffanswer表isfinish和答题结束时间*/
	// db.Model(&model.StaffAnswer{}).
	// 	Where("service_use_staff_id = ? AND staff_id = ? AND is_finish = ? AND deleted_at = ?", staffAnsID, 123, 2, "").
	// 	Updates(map[string]interface{}{"is_finish": 1, "deleted_at": time.Now().Format("2006-01-02 15:04:05")})

	// /*写入xy_staff_auswer_option员工答案信息表*/
	// for key, value := range subjectsAnswersArr {
	// 	k, _ := strconv.Atoi(key)
	// 	v, _ := strconv.Atoi(value)
	// 	staffAnsOpe := model.StaffAnswerOpetion{
	// 		SubjectID:     k,
	// 		SubjectAnsID:  v,
	// 		StaffAnswerID: staAnsID,
	// 		StaffID:       123,
	// 	}
	// 	db.Create(&staffAnsOpe)
	// }

	// /*生成员工报告*/
	// var reportStaff model.ReportStaff
	// reportStaff = model.ReportStaff{
	// 	StaffAnsID:   staAnsID,
	// 	Name:         gauge.ShowName,
	// 	HideRepInstr: reportSet.HideReportInstr,
	// 	RepInstr:     reportSet.ReportInstr,
	// 	HideDesc:     reportSet.HideDismmisionDes,
	// 	// ItGtDesc       string
	// 	Describe:       reportSet.DismmisinDes,
	// 	HideShowMethod: reportSet.HideShowMethod,
	// 	ShowMethod:     reportSet.ShowMethod,
	// 	HideCliches:    reportSet.HideCliches,
	// 	Cliches:        reportSet.Cliches,
	// 	HideComment:    reportSet.HideReportCom,
	// 	Comment:        reportSet.ReportCom,
	// 	HideDimSuggest: reportSet.HideDimSuggest,
	// 	GaugeID:        gauge.GaugeID,
	// 	StaffID:        123,
	// 	StaffName:      "staff_id",
	// 	StaffAge:       18,
	// 	// Position       string
	// 	// Marriage       int
	// 	CompanyID:   1,
	// 	CompanyName: "company_name",
	// 	// Number         :createRepNO("SNO"),
	// 	Status: 1,
	// 	// TotalScore     float32
	// 	TemplateID: gauge.TemplateID,
	// 	// GenerateDate   string
	// 	// CreatedTime    string
	// 	// UpdatedTime    string
	// 	// DeletedTime    string
	// }

	// /*存入员工报告表*/
	// db.Create(&reportStaff)

	// /*根据模板报告ID TemplateID，计算对应维度得分*/
	// if gauge.TemplateID == 4 {
	// 	/*理清$staffDimensionScore、$explainStaff，主要写入的即是这两张表*/
	// } else {
	// 	log.Printf("%d未知的模板报告！", gauge.TemplateID)
	// }

	// /*计算员工报告中答题得分*/
	// var totalScore float32
	// db.Table("xy_subject_answer").Select("SUM(xy_subject_answer.fraction)").
	// 	Joins("left join xy_staff_answer_opetion on xy_staff_answer_opetion.subject_answer_id = xy_subject_answer.id").
	// 	Where("xy_staff_answer_opetion.subject_answer_id = ?", staAnsID).Find(&totalScore)

	// /*生成员工报告数据表reportStaffData数据
	// 查询出xy_norm_explain中的f.id(常模说明项ID),
	// f.name(说明项名称),f.score_introduce(分值说明),f.coach_proposal(辅导建议)
	// 然后以json格式写入reportStaffData->report_data_extra以及reportStaffData->report_data 中
	// */

	// /*更新员工报告report_staff答题得分*/
	// /*更新员工答题表xy_staff_answer得分*/

	// //若返回json数据，可以直接使用gin封装好的JSON方法
	// c.JSON(http.StatusOK, subjectsAnswers)
	// return
}

/*获取map中的key值，并存放在slice中*/
func getKeysFromMap(m map[string]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func createRepNO(s string) string {
	result := []byte(s)
	timeSub := strconv.FormatInt(time.Now().Unix(), 20)
	now := []byte(timeSub)
	randstr := []byte(strconv.FormatInt(randInt64(1000, 9999), 4))
	result = append(result, now...)
	result = append(result, randstr...)
	return string(result)
}

func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
