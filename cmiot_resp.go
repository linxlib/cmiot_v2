package cmiot_v2

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// 卡状态
type UserStatusRealSingle struct {
	*BaseResponse
	Result []struct {
		Status string `json:"STATUS"`
	} `json:"result"`
}

// 开关机状态
type UserOnOffRealSingle struct {
	*BaseResponse
	Result []struct {
		Status string `json:"status"`
	} `json:"result"`
}

// 生命周期
type UserLifeCyCle struct {
	*BaseResponse
	Result []struct {
		LifeCycle string `json:"lifecycle"`
		OpenTime  string `json:"opentime"`
	} `json:"result"`
}

type Gprs struct {
	Left       string `json:"left"` //剩余
	ProdId     string `json:"prodid"`
	ProdInstId string `json:"prodinstid"`
	ProdName   string `json:"prodname"`
	Total      string `json:"total"` //总流量
	Used       string `json:"used"`  //已使用
}

type GprsBatch struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Msisdn    string `json:"msisdn"`
	ProdId    string `json:"prodId"`
	PodInstId string `json:"podInstId"`
	ProdName  string `json:"prodName"`
	GprsTotal string `json:"gprsTotal"`
	GprsUsed  string `json:"gprsUsed"`
	GprsLeft  string `json:"gprsLeft"`
}

type GprsBatchResult struct {
	*BaseResponse
	Result []GprsBatch `json:"result"`
}

type GprsRealTimeInfo struct {
	*BaseResponse
	Result []struct {
		Gprs []*Gprs `json:"gprs"`
	} `json:"result"`
}

type CardUnusualNum struct {
	*BaseResponse
	Result []struct {
		CardUnusualNum string `json:"cardUnusualNum"`
	} `json:"result"`
}

type GprsRealSingle struct {
	APN        string `json:"APN"`
	GPRSSTATUS string `json:"GPRSSTATUS"`
	IP         string `json:"IP"`
	RAT        string `json:"RAT"`
}
type GprsRealSingleResp struct {
	*BaseResponse
	Result []*GprsRealSingle `json:"result"`
}

type GprsUsedInfoSingle struct {
	TotalGps string `json:"total_gprs"`
}

type GprsUsedInfoSingleResp struct {
	*BaseResponse
	Result []*GprsUsedInfoSingle `json:"result"`
}

type BalanceRealSingle struct {
	Balance string `json:"balance"`
}

type BalanceRealSingleResp struct {
	*BaseResponse
	Result []*BalanceRealSingle `json:"result"`
}

type GroupUserInfo struct {
	Total string `json:"total"`
}

type GroupUserInfoResp struct {
	*BaseResponse
	Result []*GroupUserInfo `json:"result"`
}

type GprsUsedInfoSingleByDate struct {
	Gprs string `json:"gprs"`
}
type GprsUsedInfoSingleByDateResp struct {
	*BaseResponse
	Result []*GprsUsedInfoSingleByDate `json:"result"`
}

/*
"normalnum": "28596", "othernum": "3", "sleepnum": "2001", "testnum": "564"
*/
type QueryCardCount struct {
	NormalNum int `json:"normalnum"`
	OtherNum  int `json:"othernum"`
	SleepNum  int `json:"sleepnum"`
	TestNum   int `json:"testnum"`
}

type QueryCardCountResp struct {
	*BaseResponse
	Result []*QueryCardCount `json:"result"`
}
