package model

type postRequst struct {
	accessToken     string
	gaugeID         int
	subjectsAnswers string
}

type postRespone struct {
	resCode    int
	resData    string
	resMessage string
}
