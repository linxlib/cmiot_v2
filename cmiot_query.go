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

type BaseQueryWithDate struct {
	*PublicQuery
	Msisdn    string `query:"msisdn"`
	QueryDate string `query:"query_date"`
}

type BaseQueryWithDate1 struct {
	*PublicQuery
	Msisdn    string `query:"msisdn"`
	QueryDate string `query:"queryDate"`
}

type BaseQueryWithCardInfo struct {
	*PublicQuery
	CardInfo string `query:"card_info"`
	CardType int    `query:"type"`
}

type BaseQueryNoCardId struct {
	*PublicQuery
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
