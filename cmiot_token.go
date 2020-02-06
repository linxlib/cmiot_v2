package cmiot_v2

import (
	"fmt"
	"math/rand"
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

func randInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func getTransid(appId string) string {
	autoid := randInt(1, 99999998)
	return fmt.Sprintf("%s%s%.8d", appId, time.Now().Format("20060102150405"), autoid)
}
