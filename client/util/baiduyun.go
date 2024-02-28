package util

import (
	"encoding/json"
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

// 审核成功返回true
func CheckText(text string) bool {
	//如果是百度云ak sk,使用下面的客户端
	client := censor.NewCloudClient("ALTAKPGpFFPrZ3qBdl6TqGJaqk", "68018a53853c48679b1a29df3102a6d3")
	res := client.TextCensor(text)
	var t T
	json.Unmarshal([]byte(res), &t)
	return t.ConclusionType == 1
}

type T struct {
	LogId          int64  `json:"log_id"`
	Conclusion     string `json:"conclusion"`
	ConclusionType int    `json:"conclusionType"`
}
