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
	}
}
