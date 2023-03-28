package initialization

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

var larkClient *lark.Client
var FeishuBaseUrl = "https://open.f.mioffice.cn"

func LoadLarkClient(config Config) {
	customBaseUrlOption := func(cfg *larkcore.Config) {
		cfg.BaseUrl = FeishuBaseUrl
	}
	larkClient = lark.NewClient(config.FeishuAppId, config.FeishuAppSecret, customBaseUrlOption)

}

func GetLarkClient() *lark.Client {
	return larkClient
}
