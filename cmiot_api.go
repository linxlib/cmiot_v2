package cmiot_v2

import (
	"errors"
	"strconv"
)

//EBID_gprsrealsingle 在线信息实时查询
func (c *CMIOTClient) GetGprsRealSingle(msisdn string) (*GprsRealSingle, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_gprsrealsingle, c.apiId, c.apiPassword),
		Msisdn:      msisdn,
	}
	var data GprsRealSingleResp
	err := c.call(EBID_gprsrealsingle, qry, &data)
	if err != nil {
		return nil, err
	}
	if data.Status == "0" {
		return data.Result[0], nil
	} else {
		return nil, errors.New(data.Message)
	}
}

// EBID_userstatusrealsingle
func (c *CMIOTClient) GetSimStatus(Msisdn string) (string, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_userstatusrealsingle, c.apiId, c.apiPassword),
		Msisdn:      Msisdn,
	}
	var data UserStatusRealSingle
	err := c.call(EBID_userstatusrealsingle, qry, &data)
	if err != nil {
		return "", err
	}

	if data.Status == "0" {
		return data.Result[0].Status, nil
	} else {
		return "", errors.New(data.Message)
	}

}

//EBID_gprsusedinfosingle
func (c *CMIOTClient) GetGprsUsedInfSingle(Msisdn string) (*GprsUsedInfoSingle, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_gprsusedinfosingle, c.apiId, c.apiPassword),
		Msisdn:      Msisdn,
	}
	var data GprsUsedInfoSingleResp
	err := c.call(EBID_gprsusedinfosingle, qry, &data)
	if err != nil {
		return nil, err
	}

	if data.Status == "0" {
		return data.Result[0], nil
	} else {
		return nil, errors.New(data.Message)
	}
}

//EBID_gprsrealtimeinfo
func (c *CMIOTClient) GetSimDataMargin(Msisdn string) (*Gprs, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_gprsrealtimeinfo, c.apiId, c.apiPassword),
		Msisdn:      Msisdn,
	}
	var data GprsRealTimeInfo
	err := c.call(EBID_gprsrealtimeinfo, qry, &data)
	if err != nil {
		return nil, err
	}

	if data.Status == "0" {
		if len(data.Result) < 1 {
			return nil, errors.New("data Result len is zero")
		}
		if len(data.Result[0].Gprs) < 1 {
			return nil, errors.New("data Gprs len is zero")
		}
		return data.Result[0].Gprs[0], nil
	} else {
		return nil, errors.New(data.Message)
	}

}

//EBID_onandoffrealsingle
func (c *CMIOTClient) GetOnOff(Msisdn string) (bool, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_onandoffrealsingle, c.apiId, c.apiPassword),
		Msisdn:      Msisdn,
	}
	var data UserOnOffRealSingle
	err := c.call(EBID_onandoffrealsingle, qry, &data)
	if err != nil {
		return false, err
	}

	if data.Status == "0" {
		return data.Result[0].Status == "1", nil
	} else {
		return false, errors.New(data.Message)
	}
}

//EBID_querycardlifecycle
func (c *CMIOTClient) QueryCardLifeCycle(Msisdn string) (bool, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_querycardlifecycle, c.apiId, c.apiPassword),
		Msisdn:      Msisdn,
	}
	var data UserLifeCyCle
	err := c.call(EBID_querycardlifecycle, qry, &data)
	if err != nil {
		return false, err
	}

	if data.Status == "0" {
		return data.Result[0].LifeCycle == "00", nil
	} else {
		return false, errors.New(data.Message)
	}
}

//EBID_queryabnormalcardcount
func (c *CMIOTClient) GetUnnormalCount() (int64, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_queryabnormalcardcount, c.apiId, c.apiPassword),
	}
	var data CardUnusualNum
	err := c.call(EBID_queryabnormalcardcount, qry, &data)
	if err != nil {
		return -1, err
	}
	if data.Status == "0" {
		i, _ := strconv.ParseInt(data.Result[0].CardUnusualNum, 10, 64)
		return i, nil
	} else {
		return -1, errors.New(data.Message)
	}
}

//EBID_balancerealsingle
func (c *CMIOTClient) GetBalanceRealSingle(msidsn string) (*BalanceRealSingle, error) {
	qry := BaseQuery{
		PublicQuery: newPublicQuery(EBID_balancerealsingle, c.apiId, c.apiPassword),
		Msisdn:      msidsn,
	}
	var data BalanceRealSingleResp
	err := c.call(EBID_balancerealsingle, qry, &data)
	if err != nil {
		return nil, err
	}

	if data.Status == "0" {
		return data.Result[0], nil
	} else {
		return nil, errors.New(data.Message)
	}
}

//EBID_groupuserinfo
func (c *CMIOTClient) GetGroupUserInfo() (*GroupUserInfo, error) {
	qry := BaseQueryWithDate{
		PublicQuery: newPublicQuery(EBID_groupuserinfo, c.apiId, c.apiPassword),
	}
	var data GroupUserInfoResp
	err := c.call(EBID_groupuserinfo, qry, &data)
	if err != nil {
		return nil, err
	}

	if data.Status == "0" {
		return data.Result[0], nil
	} else {
		return nil, errors.New(data.Message)
	}
}

//EBID_gprsusedinfosinglebydate
func (c *CMIOTClient) GetGprsUsedInfoSingleByDate(msidsn string, queryDate string) (*GprsUsedInfoSingleByDate, error) {
	qry := BaseQueryWithDate1{
		PublicQuery: newPublicQuery(EBID_gprsusedinfosinglebydate, c.apiId, c.apiPassword),
		Msisdn:      msidsn,
		QueryDate:   queryDate,
	}
	var data GprsUsedInfoSingleByDateResp
	err := c.call(EBID_gprsusedinfosinglebydate, qry, &data)
	if err != nil {
		return nil, err
	}

	if data.Status == "0" {
		return data.Result[0], nil
	} else {
		return nil, errors.New(data.Message)
	}
}

//EBID_querycardcount
func (c *CMIOTClient) QueryCardCount() (*QueryCardCount, error) {
	qry := BaseQueryNoCardId{
		PublicQuery: newPublicQuery(EBID_querycardcount, c.apiId, c.apiPassword),
	}
	var data QueryCardCountResp
	err := c.call(EBID_querycardcount, qry, &data)
	if err != nil {
		return nil, err
	}

	if data.Status == "0" {
		return data.Result[0], nil
	} else {
		return nil, errors.New(data.Message)
	}
}
