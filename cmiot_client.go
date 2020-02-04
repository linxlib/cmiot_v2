package cmiot_v2

import "encoding/json"

type CMIOTClient struct {
	client      *CMIOTHttp
	debug       bool
	forwarderIp string

	apiId       string
	apiPassword string
	ecid        string
}

func NewCMIOTClient(apiId string, apiPassword string, ecid string) *CMIOTClient {
	return &CMIOTClient{apiId: apiId, apiPassword: apiPassword, ecid: ecid,
		client: NewCMIOTHttp(),
	}
}

func (c *CMIOTClient) setProxy(proxyUrl string) {
	c.client.setProxy(proxyUrl)
}

func (c *CMIOTClient) unsetProxy() {
	c.client.unsetProxy()
}

func (c *CMIOTClient) SetDebug(debug bool) {
	c.client.isDebug = debug
}

func (c *CMIOTClient) call(ebid string, query interface{}, body interface{}) error {
	url := buildUrl(ebid)
	qry, err := encodeQuery(query)
	if err != nil {
		return err
	}
	url += "?" + qry

	result, err := c.client.request("GET", url, "", 30000, c.forwarderIp)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), body)
	if err != nil {
		return err
	}
	return nil
}

func (c *CMIOTClient) SetForwarderId(ip string) bool {
	c.forwarderIp = ip
	return true
}
