package cmiot_v2

import (
	"fmt"
	"time"
)

func getToken(transId string, appId string, appPassword string) string {
	s := fmt.Sprintf("%s%s%s", appId, appPassword, transId)
	return getSHA256HashCode([]byte(s))
}

func getTokenAndTransId(appId string, appPassword string) (token string, transId string) {
	transId = getTransid(appId)
	token = getToken(transId, appId, appPassword)
	return token, transId
}

func getTransid(appId string) string {
	autoid := 1
	return fmt.Sprintf("%s%s%.8d", appId, time.Now().Format("20060102")+"120344", autoid)
}
