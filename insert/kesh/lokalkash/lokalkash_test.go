package lokalkash_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"githud.com/test-task/insert/kesh"
	"githud.com/test-task/insert/kesh/lokalkash"
	"githud.com/test-task/pkg/e"
)

var _ kesh.Kesh = (*lokalkash.LokalKash)(nil)

// тест вункции lokalkash.New
func TestNew(t *testing.T) {
	testCases := []struct {
		title    string
		data     []string
		expected map[string]time.Duration
		//errExpected error
	}{
		{
			"test1",
			[]string{
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
			},
			map[string]time.Duration{
				"google.com": time.Hour, "youtube.com": time.Hour, "facebook.com": time.Hour, "baidu.com": time.Hour,
				"wikipedia.org": time.Hour, "qq.com": time.Hour, "taobao.com": time.Hour, "yahoo.com": time.Hour,
				"tmall.com": time.Hour, "amazon.com": time.Hour, "google.co.in": time.Hour, "twitter.com": time.Hour,
				"sohu.com": time.Hour, "jd.com": time.Hour, "live.com": time.Hour, "instagram.com": time.Hour,
				"sina.com.cn": time.Hour, "weibo.com": time.Hour, "google.co.jp": time.Hour, "reddit.com": time.Hour,
				"vk.com": time.Hour, "360.cn": time.Hour, "login.tmall.com": time.Hour, "blogspot.com": time.Hour,
				"yandex.ru": time.Hour, "google.com.hk": time.Hour, "netflix.com": time.Hour, "linkedin.com": time.Hour,
				"pornhub.com": time.Hour, "google.com.br": time.Hour,
				"twitch.tv": time.Hour, "pages.tmall.com": time.Hour, "csdn.net": time.Hour, "yahoo.co.jp": time.Hour,
				"mail.ru": time.Hour, "aliexpress.com": time.Hour, "alipay.com": time.Hour, "office.com": time.Hour,
				"google.fr": time.Hour, "google.ru": time.Hour, "google.co.uk": time.Hour, "microsoftonline.com": time.Hour,
				"google.de": time.Hour, "ebay.com": time.Hour, "microsoft.com": time.Hour, "livejasmin.com": time.Hour,
				"t.co": time.Hour, "bing.com": time.Hour, "xvideos.com": time.Hour, "google.ca": time.Hour,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result := lokalkash.New(tC.data...)
			if !reflect.DeepEqual(result.Kash, tC.expected) {
				t.Errorf("invalid func lokalkash.New\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}
		})
	}
}

// тест метода lokalkash.Singl
func TestSingl(t *testing.T) {
	testCases := []struct {
		title       string
		domain      string
		expected    *kesh.Website
		errExpected error
	}{
		{
			"test1",
			"youtube.com",
			&kesh.Website{Domain: "youtube.com", Delay: time.Second},
			nil,
		},
		{
			"test2",
			"facebook.com",
			&kesh.Website{Domain: "facebook.com", Delay: time.Second},
			nil,
		},
		{
			"test3",
			"ya.ru",
			nil,
			e.Err(lokalkash.ErrSignal, errors.New(lokalkash.ErrNotFound)),
		},
	}

	kash := lokalkash.New(
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
	)

	kash.Updata(kesh.New("youtube.com", time.Second))
	kash.Updata(kesh.New("facebook.com", time.Second))

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, err := kash.Singl(tC.domain)
			if !reflect.DeepEqual(err, tC.errExpected) {
				t.Errorf("invalid method lokalkash.Singl\nerr xpected:\n%v\nerr:\n%v", tC.errExpected, err)
				return
			}
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("invalid method lokalkash.Singl\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}
		})
	}
}

// тест метода kash.Updata
func TestUpdata(t *testing.T) {
	testCases := []struct {
		title       string
		data        *kesh.Website
		expected    map[string]time.Duration
		errExpected error
	}{
		{
			"test1",
			&kesh.Website{Domain: "google.com", Delay: time.Millisecond * 100},
			map[string]time.Duration{
				"google.com": time.Millisecond * 100, "youtube.com": time.Hour,
			},
			nil,
		},
		{
			"test2",
			&kesh.Website{Domain: "youtube.com", Delay: time.Millisecond * 250},
			map[string]time.Duration{
				"google.com": time.Millisecond * 100, "youtube.com": time.Millisecond * 250,
			},
			nil,
		},
		{
			"test3",
			&kesh.Website{Domain: "ya.ru", Delay: time.Millisecond * 250},
			map[string]time.Duration{
				"google.com": time.Millisecond * 100, "youtube.com": time.Millisecond * 250,
			},
			e.Err(lokalkash.ErrUpdata, errors.New(lokalkash.ErrNotFound)),
		},
	}

	kash := lokalkash.New(
		"google.com", "youtube.com",
	)

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			err := kash.Updata(tC.data)
			if !reflect.DeepEqual(err, tC.errExpected) {
				t.Errorf("invalid method lokalkash.Updata\nerr xpected:\n%v\nerr:\n%v", tC.errExpected, err)
				return
			}
			if !reflect.DeepEqual(kash.Kash, tC.expected) {
				t.Errorf("invalid method lokalkash.Updata\nexpected:\n%v\nresult:\n%v", tC.expected, kash.Kash)
			}
		})
	}
}

// тест метода lokalkash.Min
func TestMin(t *testing.T) {
	testCases := []struct {
		title string
		// data     *kesh.Website
		expected    *kesh.Website
		errExpected error
	}{
		{
			"test1",
			kesh.New("bing.com", time.Millisecond*85),
			nil,
		},
	}

	k := lokalkash.New(
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
	)
	err := k.Updata(kesh.New("google.com", time.Millisecond*150))
	if err != nil {
		t.Error(err)
	}

	err = k.Updata(kesh.New("google.ca", time.Millisecond*90))
	if err != nil {
		t.Error(err)
	}

	err = k.Updata(kesh.New("bing.com", time.Millisecond*85))
	if err != nil {
		t.Error(err)
	}

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, _ := k.Min()
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("invalid method lokalkash.Min\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}

		})
	}
}

// тест метода lokalkash.Max
func TestMax(t *testing.T) {
	testCases := []struct {
		title string
		// data     *kesh.Website
		expected    *kesh.Website
		errExpected error
	}{
		{
			"test1",
			kesh.New("google.com", time.Millisecond*150),
			nil,
		},
	}

	k := lokalkash.New(
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
	)
	err := k.Updata(kesh.New("google.com", time.Millisecond*150))
	if err != nil {
		t.Error(err)
	}

	err = k.Updata(kesh.New("google.ca", time.Millisecond*90))
	if err != nil {
		t.Error(err)
	}

	err = k.Updata(kesh.New("bing.com", time.Millisecond*85))
	if err != nil {
		t.Error(err)
	}

	err = k.Updata(kesh.New("bing.com", time.Hour*6))
	if err != nil {
		t.Error(err)
	}

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, _ := k.Max()
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("invalid method lokalkash.Max\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}

		})
	}
}

// тест метода lokalkash.list
func TestList(t *testing.T) {
	testCases := []struct {
		title string
		//data     []string
		expected []string
		//errExpected error
	}{
		{
			"test1",
			[]string{
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
			},
		},
	}

	k := lokalkash.New("google.com", "youtube.com", "facebook.com", "baidu.com",
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
		"bing.com", "xvideos.com", "google.ca")

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, _ := k.List()
			if len(result) != len(tC.expected) {
				t.Errorf("invalid func lokalkash.List\nexpected:\n%v\nresult:\n%v", len(tC.expected), len(result))
			}
		})
	}
}
