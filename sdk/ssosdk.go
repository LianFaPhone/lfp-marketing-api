package sdk

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

// 参数说明：区域，密钥Id，密钥，token（一般为空）
func NewSsoSdk(region, accessKeyID, accessKeySecret string) (*SsoSdk, error) {
	client, err := oss.New(region, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &SsoSdk{
		mClient : client,
	}, nil
	//GSsoSdk.mClient = client
	//return &GSsoSdk, nil
}

//var GSsoSdk SsoSdk

type SsoSdk struct {
	mClient *oss.Client
}

func (this *SsoSdk) GetOrInitBucket(name string) (*oss.Bucket, error) {
	exist, err := this.IsBucketExist(name)
	if err != nil {
		return nil, err
	}
	if !exist {
		if err = this.CreateBucket(name); err != nil {
			return nil, err
		}
	}
	bucket, err := this.mClient.Bucket(name)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func (this *SsoSdk) Close() {

}

//创建bukect, 默认为标准存储类型，私有权限。根据实际情况 还得改
func (this *SsoSdk) CreateBucket(name string) error {
	err := this.mClient.CreateBucket(name)
	return err
}

//创建bukect
func (this *SsoSdk) IsBucketExist(name string) (bool, error) {
	isExist, err := this.mClient.IsBucketExist(name)
	if err != nil {
		return isExist, err
	}
	return isExist, err
}

func (this *SsoSdk) PutObject(bucket *oss.Bucket, fileName string, body io.Reader) error {
	err := bucket.PutObject(fileName, body)
	return err
}

func (this *SsoSdk) GetObject(bucket *oss.Bucket, fileName string) (io.Reader, error) {
	return bucket.GetObject(fileName)

}
