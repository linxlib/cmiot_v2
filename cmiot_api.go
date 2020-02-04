package cmiot_v2

import (
	"errors"
	"strconv"
)

func (c *CMIOTClient) GetSimStatus(Msisdn string) (string, error) {
	qry := UserStatusRealSingleQuery{
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

func (c *CMIOTClient) GetSimDataMargin(Msisdn string) (*Gprs, error) {
	qry := GprsRealTimeInfoQuery{
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
