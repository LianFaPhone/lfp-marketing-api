package db

import (
	"LianFaPhone/lfp-marketing-api/config"
	"errors"
	"github.com/bluele/gcache"
	"time"
)

var GCache Cache

type Cache struct {
	BlacklistAreaCache gcache.Cache
}

func (this *Cache) Init() {
}

//**************************************************/
func (this *Cache) GetBlacklistArea(code string) (interface{}, error) {
	if this.BlacklistAreaCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.BlacklistAreaCache.Get(code)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetBlacklistArea(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.BlacklistAreaCache = gcache.New(config.GConfig.Cache.SponsorMaxKey).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveBlacklistArea(tp string) {
	if this.BlacklistAreaCache == nil {
		return
	}
	this.BlacklistAreaCache.Remove(tp)
}
