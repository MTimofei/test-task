// фаил содержит глобальнве переменые
package config

import "flag"

const Key string = "admin"

var (
	Host    *string = flag.String("host", "0.0.0.0:80", "address for starting the server")
	Timeout int     = *flag.Int("timeout", 10, "sed timeout")
)

var Domains = []string{
	"google.com", "youtube.com", "facebook.com", "baidu.com",
	"wikipedia.org", "qq.com", "taobao.com", "yahoo.com",
	"tmall.com", "amazon.com", "google.co.in", "twitter.com",
	"sohu.com", "jd.com", "live.com", "instagram.com",
	"sina.com.cn", "weibo.com", "google.co.jp", "reddit.com",
	"vk.com", "360.cn", "login.tmall.com", "blogspot.com",
	"yandex.ru", "google.com.hk", "netflix.com",
	"linkedin.com", "pornhub.com", "google.com.br", "twitch.tv",
	"pages.tmall.com", "csdn.net", "yahoo.co.jp", "mail.ru",
	"aliexpress.com", "alipay.com", "office.com", "google.fr",
	"google.ru", "google.co.uk", "microsoftonline.com", "google.de",
	"ebay.com", "microsoft.com", "livejasmin.com", "t.co",
	"bing.com", "xvideos.com", "google.ca",
}
