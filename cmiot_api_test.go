package cmiot_v2

import (
	"fmt"
	"testing"
)

const (
	APPID    = "SIPVG623MUTK9Q"
	PASSWORD = "q1j2WLzJ1"
	ECID     = "571A57177722637"
)

func TestCMIOTClient_QueryCardCount(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")

	c.QueryCardCount()
}

func TestCMIOTClient_GetGprsUsedInfoSingleByDate(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")

	c.GetGprsUsedInfoSingleByDate("1440473631364", "20200206")
}

func TestCMIOTClient_GetGroupUserInfo(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetGroupUserInfo("20200204")
}

func TestCMIOTClient_GetBalanceRealSingle(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetBalanceRealSingle("1440066339952")
}

func TestCMIOTClient_GetUnnormalCount(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetUnnormalCount()
}

func TestCMIOTClient_QueryCardLifeCycle(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.QueryCardLifeCycle("1440066339952")
}

func TestCMIOTClient_GetOnOff(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetOnOff("1440066339952")
}

func TestCMIOTClient_GetSimDataMargin(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetSimDataMargin("1440303458165")
}

func TestCMIOTClient_GetSimStatus(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	fmt.Println(c.GetSimStatus("1440066339952"))
}

func TestCMIOTClient_GetGprsUsedInfSingle(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetGprsUsedInfSingle("1440066339952")
}

func TestCMIOTClient_GetGprsRealSingle(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.GetGprsRealSingle("1440066339952")
}

/*
0—msisdn ⚫ 1—imsi
⚫ 2—iccid
*/
func TestCMIOTClient_QueryCardInfo(t *testing.T) {
	c := NewCMIOTClient(APPID, PASSWORD, ECID)
	c.SetDebug(true)
	c.SetForwarderId("111.231.118.185")
	c.QueryCardInfo("89860406111980053250", 2)
}
