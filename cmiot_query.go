package cmiot_v2

type PublicQuery struct {
	AppID   string `query:"appid"`
	TransID string `query:"transid"`
	EBID    string `query:"ebid"`
	Token   string `query:"token"`
}

type BaseQuery struct {
	*PublicQuery
	Msisdn string `query:"msisdn"`
}

func newPublicQuery(ebid string, appId string, appPassword string) *PublicQuery {
	token, transId := getTokenAndTransId(appId, appPassword)
	return &PublicQuery{
		AppID:   appId,
		TransID: transId,
		EBID:    ebid,
		Token:   token,
	}
}

// 单卡实时信息Query
type UserStatusRealSingleQuery struct {
	*PublicQuery
	Msisdn string `query:"msisdn"`
}

// 单卡流量实时信息
type GprsRealTimeInfoQuery struct {
	*PublicQuery
	Msisdn string `query:"msisdn"`
}

// 批量流量实时信息
type GprsBatchRealTimeInfoQuery struct {
	*PublicQuery
	Msisdns string `query:"msisdns"`
}
