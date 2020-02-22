package controllers

import (
//	s3 "LianFaPhone/bas-tools/sdk.aws.s3"
	"github.com/kataras/iris"
	"net/url"
	"strings"

	"crypto/md5"
	"fmt"
	"io/ioutil"
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-marketing-api/config"
	"go.uber.org/zap"
	. "LianFaPhone/lfp-base/log/zap"
	"mime/multipart"
	"os"
	"io"
	"LianFaPhone/lfp-marketing-api/sdk"
)

type UploadFile struct {
	Controllers
}

type LogoFileAddr struct {
	File string `json:"file"`
	Addr string `json:"addr"`
}

func ParseFiles(ctx iris.Context) ([]*multipart.FileHeader, error){
	err := ctx.Request().ParseMultipartForm(32 << 23)
	if err != nil {
		return nil,err
	}
	files := ctx.Request().MultipartForm.File["file"]
	if len(files) == 0 {
		return nil, fmt.Errorf("nil file")
	}
	return files,nil
}

func (this *UploadFile) PhotoUpload(ctx iris.Context) {
	files,err := ParseFiles(ctx)
	if err != nil {
		ZapLog().Error( "ParseFiles err", zap.Error(err), zap.Any("headers", ctx.Request().Header))
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "ParseMultipartForm_ERRORS")
		return
	}
	results := make([]LogoFileAddr, 0)
	for i := 0; i < len(files); i++ {
		filename := files[i].Filename
		file, err := files[i].Open()
		if err != nil {
			ZapLog().Error("FileOpen  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "FileOpen_ERRORS")
			return
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			ZapLog().Error( "FileReadAll  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_ReadAll_ERRORS")
			return
		}
		if len(content) > 30*1024*1024 {
			ZapLog().Error( "file too large err")
			this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "file too large")
			return
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			ZapLog().Error( "FileSeek  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_Seek_ERRORS")
			return
		}
		filenameMd5 := fmt.Sprintf("%X", md5.Sum(content))
		filenameMd5 = filenameMd5 +GetFileExt(filename)
		localfile, err := os.Create("./photos/"+filenameMd5)
		if err != nil {
			ZapLog().Error( "FileSeek  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_Seek_ERRORS")
			return
		}
		_, err = io.Copy(localfile, file)
		if err != nil {
			ZapLog().Error( "FileSeek  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_Seek_ERRORS")
			return
		}
		results = append(results, LogoFileAddr{filenameMd5, "./photos/"+filenameMd5})
	}
	this.Response(ctx, results)

}

func (this *UploadFile) PhotoDownload(ctx iris.Context) {
	fileName := ctx.FormValue("url")
	if len(fileName) <= 0 {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err")
		return
	}

	fileName = GetFileName(fileName)

	file, err := os.Open("./photos/"+fileName)
	if err != nil {
		ZapLog().Error( "file open  err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "file open err")
		return
	}
	defer file.Close()

	Content, err := ioutil.ReadAll(file)
	if err != nil {
		ZapLog().Error( "File read  err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_read err")
		return
	}
	ctx.ResponseWriter().Header().Add("Content-Type", "application/octet-stream")
	ctx.ResponseWriter().Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	ctx.Write(Content)
}

func (this *UploadFile) UpSso(ctx iris.Context) {
	files,err := ParseFiles(ctx)
	if err != nil {
		ZapLog().Error( "ParseFiles err", zap.Error(err), zap.Any("headers", ctx.Request().Header))
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "ParseMultipartForm_ERRORS")
		return
	}

	s3Sdk,err := sdk.NewSsoSdk(config.GConfig.Aliyun.OssEndpoint, config.GConfig.Aliyun.AccessKeyId, config.GConfig.Aliyun.AccessKeySecret)
	if err != nil {
		ZapLog().Error( "NewS3Sdk  err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "NewS3Sdk_ERRORS")
		return
	}
	defer s3Sdk.Close()

	buket,err := s3Sdk.GetOrInitBucket(config.GConfig.Aliyun.BucketName)
	if err != nil {
		ZapLog().Error( "GetOrInitBucket  err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "NewS3Sdk_ERRORS")
		return
	}

	results := make([]LogoFileAddr, 0)
	for i := 0; i < len(files); i++ {
		filename := files[i].Filename

		file, err := files[i].Open()
		if err != nil {
			ZapLog().Error( "FileOpen  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "FileOpen_ERRORS")
			return
		}
		defer file.Close()

		path := ""
		if len(config.GConfig.Aliyun.UpfilePath) > 0 {
			path = config.GConfig.Aliyun.UpfilePath + "/"
		}
		err = s3Sdk.PutObject(buket, path+filename, file)
		if err != nil {
			ZapLog().Error( "S3 UpLoad  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), "S3_UpLoad_ERRORS")
			return
		}
		addr := strings.TrimPrefix(config.GConfig.Aliyun.OssEndpoint, "http://")
		addr = strings.TrimPrefix(config.GConfig.Aliyun.OssEndpoint, "https://")
		addr = "https://"+config.GConfig.Aliyun.BucketName+"."+addr+"/"+path+url.PathEscape(filename)
		//注意oss开启https
		results = append(results, LogoFileAddr{filename, addr})
	}

	this.Response(ctx, results)
}

func (this *UploadFile) UpSsoCardClassPic(ctx iris.Context) {
	files,err := ParseFiles(ctx)
	if err != nil {
		ZapLog().Error( "ParseFiles err", zap.Error(err), zap.Any("headers", ctx.Request().Header))
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "ParseMultipartForm_ERRORS")
		return
	}

	s3Sdk,err := sdk.NewSsoSdk(config.GConfig.Aliyun.OssEndpoint, config.GConfig.Aliyun.AccessKeyId, config.GConfig.Aliyun.AccessKeySecret)
	if err != nil {
		ZapLog().Error( "NewS3Sdk  err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "NewS3Sdk_ERRORS")
		return
	}
	defer s3Sdk.Close()

	buket,err := s3Sdk.GetOrInitBucket(config.GConfig.Aliyun.BucketName)
	if err != nil {
		ZapLog().Error( "GetOrInitBucket  err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "NewS3Sdk_ERRORS")
		return
	}

	results := make([]LogoFileAddr, 0)
	for i := 0; i < len(files); i++ {
		filename := files[i].Filename

		file, err := files[i].Open()
		if err != nil {
			ZapLog().Error( "FileOpen  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "FileOpen_ERRORS")
			return
		}
		defer file.Close()
		path := ""
		if len(config.GConfig.Aliyun.CardclasspicPath) > 0 {
			path = config.GConfig.Aliyun.CardclasspicPath + "/"
		}

		err = s3Sdk.PutObject(buket, path+filename, file)
		if err != nil {
			ZapLog().Error( "S3 UpLoad  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), "S3_UpLoad_ERRORS")
			return
		}
		addr := strings.TrimPrefix(config.GConfig.Aliyun.OssEndpoint, "http://")
		addr = strings.TrimPrefix(config.GConfig.Aliyun.OssEndpoint, "https://")
		addr = "https://"+config.GConfig.Aliyun.BucketName+"."+addr+"/"+path+url.PathEscape(filename)
		//注意oss开启https
		results = append(results, LogoFileAddr{filename, addr})
	}

	this.Response(ctx, results)
}

