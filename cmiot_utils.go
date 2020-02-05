package cmiot_v2

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func getSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := strings.ToLower(hex.EncodeToString(bytes))
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}

func getApiNameByEBID(ebid string) string {
	switch ebid {
	case EBID_gprsrealsingle:
		return "gprsrealsingle"
	case EBID_userstatusrealsingle:
		return "userstatusrealsingle"
	//case EBID_cardinfo:
	//	return "cardinfo"
	case EBID_gprsusedinfosingle:
		return "gprsusedinfosingle"
	case EBID_onandoffrealsingle:
		return "onandoffrealsingle"
	//case EBID_batchsmsusedbydate:
	//	return "batchsmsusedbydate"
	//case EBID_batchgprsusedbydate:
	//	return "batchgprsusedbydate"
	case EBID_balancerealsingle:
		return "balancerealsingle"
	//case EBID_smsusedinfosingle:
	//	return "smsusedinfosingle"
	//case EBID_smsstatusreset:
	//	return "smsstatusreset"
	case EBID_groupuserinfo:
		return "groupuserinfo"
	//case EBID_smsusedbydate:
	//	return "smsusedbydate"
	case EBID_gprsrealtimeinfo:
		return "gprsrealtimeinfo"
	case EBID_arrearageuserinfo:
		return "arrearageuserinfo"
	//case EBID_querycardprodinfo:
	//	return "querycardprodinfo"
	case EBID_gprsusedinfosinglebydate:
		return "gprsusedinfosinglebydate"
	case EBID_querycardcount:
		return "querycardcount"
	case EBID_queryabnormalcardcount:
		return "queryabnormalcardcount"
	//case EBID_querygprsonlinecardcount:
	//	return "querygprsonlinecardcount"
	//case EBID_querysmsopenstatus:
	//	return "querysmsopenstatus"
	//case EBID_querygprsopenstatus:
	//	return "querygprsopenstatus"
	//case EBID_queryapnopenstatus:
	//	return "queryapnopenstatus"
	case EBID_querycardlifecycle:
		return "querycardlifecycle"
	//case EBID_useropenservice:
	//	return "useropenservice"
	//case EBID_querygroupuserchangenum:
	//	return "querygroupuserchangenum"
	//case EBID_querycardopentime:
	//	return "querycardopentime"
	//case EBID_batchquerycardinfo:
	//	return "batchquerycardinfo"
	//case EBID_batchquerycardlifecycle:
	//	return "batchquerycardlifecycle"
	//case EBID_batchquerymonthsmsinfo:
	//	return "batchquerymonthsmsinfo"
	//case EBID_batchqueryrealtimegprsinfo:
	//	return "batchqueryrealtimegprsinfo"
	//case EBID_querygprsshareused:
	//	return "querygprsshareused"
	//case EBID_batchquerygprsshareused:
	//	return "batchquerygprsshareused"

	default:
		return ""
	}
}

func buildUrl(ebid string) string {
	apiName := getApiNameByEBID(ebid)
	return APIBASE + apiName
}

func getCardStatusPB(cardStatus string) string {
	switch cardStatus {
	case "00":
		return "正常"
	case "01":
		return "单向停机"
	case "02":
		return "停机"
	case "03":
		return "预销号"
	case "04":
		return "销号"
	case "05":
		return "过户"
	case "06":
		return "休眠"
	case "07":
		return "待激活"
	case "99":
		return "号码不存在"
	default:
		return ""
	}
}

func encodeQuery(query interface{}) (string, error) {
	var qry string
	q := newQueryEncode(nil)
	if err := qryEncode(query, q); err != nil {
		return "", err
	}
	qry = q.end()
	return qry, nil
}
