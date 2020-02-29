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
	PdPartnerByIdCache gcache.Cache
	PdPartnerGoodsByCodeCache gcache.Cache
	PdPartnerGoodsByTpCache gcache.Cache
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
func (this *Cache) GetPdPartnerGoodsByCode(code string) (interface{}, error) {
	if this.PdPartnerGoodsByCodeCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.PdPartnerGoodsByCodeCache.Get(code)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetPdPartnerGoodsByCode(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.PdPartnerGoodsByCodeCache = gcache.New(config.GConfig.Cache.CardClassByNameMaxKey).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemovePdPartnerGoodsByCode(tp string) {
	if this.PdPartnerGoodsByCodeCache == nil {
		return
	}
	this.PdPartnerGoodsByCodeCache.Remove(tp)
}

//**************************************************/
func (this *Cache) GetPdPartnerGoodsById(id int64) (interface{}, error) {
	if this.PdPartnerGoodsByTpCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.PdPartnerGoodsByTpCache.Get(id)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetPdPartnerGoodsById(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.PdPartnerGoodsByTpCache = gcache.New(config.GConfig.Cache.CardClassByNameMaxKey).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemovePdPartnerGoodsById(id int64) {
	if this.PdPartnerGoodsByTpCache == nil {
		return
	}
	this.PdPartnerGoodsByTpCache.Remove(id)
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

//////////////////////////////////////////////////////////////////

func (this *Cache) GetPdPartnerById(id int64) (interface{}, error) {
	if this.PdPartnerByIdCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.PdPartnerByIdCache.Get(id)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetPdPartnerById(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.PdPartnerByIdCache = gcache.New(100).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemovePdPartnerById(id int64) {
	if this.PdPartnerByIdCache == nil {
		return
	}
	this.PdPartnerByIdCache.Remove(id)
}
