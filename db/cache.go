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
	CardClassByNameCache gcache.Cache
	CardClassByTpCache gcache.Cache
	BsProvinceByNameCache gcache.Cache
	BsCityByNameCache gcache.Cache
	BsAreaByNameCache gcache.Cache
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

//**************************************************/
func (this *Cache) GetCardClassByName(code string) (interface{}, error) {
	if this.CardClassByNameCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.CardClassByNameCache.Get(code)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetCardClassByName(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.CardClassByNameCache = gcache.New(config.GConfig.Cache.CardClassByNameMaxKey).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveCardClassByName(tp string) {
	if this.CardClassByNameCache == nil {
		return
	}
	this.CardClassByNameCache.Remove(tp)
}

//**************************************************/
func (this *Cache) GetCardClassById(id int) (interface{}, error) {
	if this.CardClassByTpCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.CardClassByTpCache.Get(id)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetCardClassById(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.CardClassByTpCache = gcache.New(config.GConfig.Cache.CardClassByNameMaxKey).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveCardClassById(id int) {
	if this.CardClassByTpCache == nil {
		return
	}
	this.CardClassByTpCache.Remove(id)
}

///////////////provice
func (this *Cache) GetProviceByName(name string) (interface{}, error) {
	if this.BsProvinceByNameCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.BsProvinceByNameCache.Get(name)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetProvinceByName(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.BsProvinceByNameCache = gcache.New(10).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveProvinceByName(name string) {
	if this.BsProvinceByNameCache == nil {
		return
	}
	this.BsProvinceByNameCache.Remove(name)
}

//////////city
func (this *Cache) GetCityByName(name string) (interface{}, error) {
	if this.BsCityByNameCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.BsCityByNameCache.Get(name)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetCityByName(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.BsCityByNameCache = gcache.New(10).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveCityByName(id string) {
	if this.BsCityByNameCache == nil {
		return
	}
	this.BsCityByNameCache.Remove(id)
}

/////////////area
func (this *Cache) GetAreaByName(name string) (interface{}, error) {
	if this.BsAreaByNameCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.BsAreaByNameCache.Get(name)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetAreaByName(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.BsAreaByNameCache = gcache.New(10).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveAreaByName(id string) {
	if this.BsAreaByNameCache == nil {
		return
	}
	this.BsAreaByNameCache.Remove(id)
}

