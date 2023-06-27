package loсalcache_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"githud.com/test-task/internal/cache"
	loсalcache "githud.com/test-task/internal/cache/localcache"
	"githud.com/test-task/pkg/e"
)

var _ cache.Cache = (*loсalcache.LocalCache)(nil)

// тест вункции loсalcache.New
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
			result := loсalcache.New(tC.data...)
			if !reflect.DeepEqual(result.Cache, tC.expected) {
				t.Errorf("invalid func loсalcache.New\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}
		})
	}
}

// тест метода lokalcache.Single
func TestSingl(t *testing.T) {
	testCases := []struct {
		title       string
		domain      string
		expected    *cache.Website
		errExpected error
	}{
		{
			"test1",
			"youtube.com",
			&cache.Website{Domain: "youtube.com", Delay: time.Second},
			nil,
		},
		{
			"test2",
			"facebook.com",
			&cache.Website{Domain: "facebook.com", Delay: time.Second},
			nil,
		},
		{
			"test3",
			"ya.ru",
			nil,
			e.Err(loсalcache.ErrSignal, errors.New(loсalcache.ErrNotFound)),
		},
	}

	c := loсalcache.New(
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

	c.Update(cache.New("youtube.com", time.Second))
	c.Update(cache.New("facebook.com", time.Second))

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, err := c.Single(tC.domain)
			if !reflect.DeepEqual(err, tC.errExpected) {
				t.Errorf("invalid method lokalcache.Single\nerr expected:\n%v\nerr:\n%v", tC.errExpected, err)
				return
			}
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("invalid method lokalcache.Single\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}
		})
	}
}

// тест метода cache.Update
func TestUpdate(t *testing.T) {
	testCases := []struct {
		title       string
		data        *cache.Website
		expected    map[string]time.Duration
		errExpected error
	}{
		{
			"test1",
			&cache.Website{Domain: "google.com", Delay: time.Millisecond * 100},
			map[string]time.Duration{
				"google.com": time.Millisecond * 100, "youtube.com": time.Hour,
			},
			nil,
		},
		{
			"test2",
			&cache.Website{Domain: "youtube.com", Delay: time.Millisecond * 250},
			map[string]time.Duration{
				"google.com": time.Millisecond * 100, "youtube.com": time.Millisecond * 250,
			},
			nil,
		},
		{
			"test3",
			&cache.Website{Domain: "ya.ru", Delay: time.Millisecond * 250},
			map[string]time.Duration{
				"google.com": time.Millisecond * 100, "youtube.com": time.Millisecond * 250,
			},
			e.Err(loсalcache.ErrUpdate, errors.New(loсalcache.ErrNotFound)),
		},
	}

	cache := loсalcache.New(
		"google.com", "youtube.com",
	)

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			err := cache.Update(tC.data)
			if !reflect.DeepEqual(err, tC.errExpected) {
				t.Errorf("invalid method loсalcache.Updata\nerr xpected:\n%v\nerr:\n%v", tC.errExpected, err)
				return
			}
			if !reflect.DeepEqual(cache.Cache, tC.expected) {
				t.Errorf("invalid method loсalcache.Updata\nexpected:\n%v\nresult:\n%v", tC.expected, cache.Cache)
			}
		})
	}
}

// тест метода loсalcache.Min
func TestMin(t *testing.T) {
	testCases := []struct {
		title string
		// data     *cache.Website
		expected    *cache.Website
		errExpected error
	}{
		{
			"test1",
			cache.New("bing.com", time.Millisecond*85),
			nil,
		},
	}

	k := loсalcache.New(
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
	err := k.Update(cache.New("google.com", time.Millisecond*150))
	if err != nil {
		t.Error(err)
	}

	err = k.Update(cache.New("google.ca", time.Millisecond*90))
	if err != nil {
		t.Error(err)
	}

	err = k.Update(cache.New("bing.com", time.Millisecond*85))
	if err != nil {
		t.Error(err)
	}

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, _ := k.Min()
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("invalid method lokalcache.Min\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}

		})
	}
}

// тест метода lokalcache.Max
func TestMax(t *testing.T) {
	testCases := []struct {
		title string
		// data     *cache.Website
		expected    *cache.Website
		errExpected error
	}{
		{
			"test1",
			cache.New("google.com", time.Millisecond*150),
			nil,
		},
	}

	k := loсalcache.New(
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
	err := k.Update(cache.New("google.com", time.Millisecond*150))
	if err != nil {
		t.Error(err)
	}

	err = k.Update(cache.New("google.ca", time.Millisecond*90))
	if err != nil {
		t.Error(err)
	}

	err = k.Update(cache.New("bing.com", time.Millisecond*85))
	if err != nil {
		t.Error(err)
	}

	err = k.Update(cache.New("bing.com", time.Hour*6))
	if err != nil {
		t.Error(err)
	}

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			result, _ := k.Max()
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("invalid method lokalcache.Max\nexpected:\n%v\nresult:\n%v", tC.expected, result)
			}

		})
	}
}

// тест метода lokalcache.list
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

	k := loсalcache.New("google.com", "youtube.com", "facebook.com", "baidu.com",
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
				t.Errorf("invalid func loсalcache.List\nexpected:\n%v\nresult:\n%v", len(tC.expected), len(result))
			}
		})
	}
}
