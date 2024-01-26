// Package config 站点配置信息
package config

import "OriginBoost/pkg/config"



func init() {
    config.Add("app", func() map[string]interface{} {
        return map[string]interface{}{
			"appid": config.Env("APPID", "******"),
			"appkey": config.Env("APPKEY", "******"),
			"endpoint": config.Env("ENDPOINT", "http://api.fanyi.baidu.com"),
			"path": config.Env("PATH", "/api/trans/vip/translate"),
        }
    })
}