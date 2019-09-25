package controllers

//import (
////	s3 "LianFaPhone/bas-tools/sdk.aws.s3"
//	"github.com/kataras/iris"
//
//	"crypto/md5"
//	"fmt"
//	"io/ioutil"
//	"strings"
//	apibackend "LianFaPhone/lfp-api/errdef"
//	"LianFaPhone/lfp-marketing-api/config"
//	"go.uber.org/zap"
//	. "LianFaPhone/lfp-base/log/zap"
//)
//
//type UploadFile struct {
//	Controllers
//}
//
//type LogoFileAddr struct {
//	File string `json:"file"`
//	Addr string `json:"addr"`
//}
//
//func (this *UploadFile) HandlePicFiles(ctx iris.Context) {
//	this.handleFilesToAwsS3(ctx, config.GConfig.Aws.PicRegion, config.GConfig.Aws.PicBucket, config.GConfig.Aws.PicBucketPath, "", config.GConfig.Aws.PicTimeout)
//}
//
//func (this *UploadFile) handleFilesToAwsS3(ctx iris.Context, region, bucket, bucketPath, filePrefix string, timeout int) {
//	//设置内存大小
//	err := ctx.Request().ParseMultipartForm(32 << 23)
//	if err != nil {
//		ZapLog().Error( "ParseMultipartForm err", zap.Error(err), zap.Any("headers", ctx.Request().Header))
//		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "ParseMultipartForm_ERRORS")
//		return
//	}
//	files := ctx.Request().MultipartForm.File["file"]
//	if len(files) == 0 {
//		ZapLog().Error( "NoFile err")
//		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "NoFind file")
//		return
//	}
//
//	s3Sdk := s3.NewS3Sdk(region, config.GConfig.Aws.AccessKeyId, config.GConfig.Aws.AccessKey, config.GConfig.Aws.AccessToken)
//	if s3Sdk == nil {
//		ZapLog().Error( "NewS3Sdk  err[return nil]")
//		this.ExceptionSerive(ctx, apibackend.BASERR_UNKNOWN_BUG.Code(), "NewS3Sdk_ERRORS")
//		return
//	}
//	defer s3Sdk.Close()
//	if timeout < 10 {
//		timeout = 10
//	}
//	results := make([]LogoFileAddr, 0)
//	for i := 0; i < len(files); i++ {
//		filename := files[i].Filename
//
//		file, err := files[i].Open()
//		if err != nil {
//			ZapLog().Error( "FileOpen  err", zap.Error(err))
//			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "FileOpen_ERRORS")
//			return
//		}
//		defer file.Close()
//		content, err := ioutil.ReadAll(file)
//		if err != nil {
//			ZapLog().Error( "FileReadAll  err", zap.Error(err))
//			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_ReadAll_ERRORS")
//			return
//		}
//		_, err = file.Seek(0, 0)
//		if err != nil {
//			ZapLog().Error( "FileSeek  err", zap.Error(err))
//			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "File_Seek_ERRORS")
//			return
//		}
//		filenameMd5 := fmt.Sprintf("%X", md5.Sum(content))
//		if len(bucketPath) != 0 {
//			bucketPath = strings.TrimRight(bucketPath, "/")
//			filenameMd5 = bucketPath + "/" + filePrefix + filenameMd5+GetFileExt(filename)
//		}else{
//			filenameMd5 = filePrefix + filenameMd5 +GetFileExt(filename)
//		}
//
//		addr, err := s3Sdk.UpLoad(bucket, filenameMd5, file, timeout, files[i].Header)
//		if err != nil {
//			ZapLog().Error( "S3 UpLoad  err", zap.Error(err))
//			this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), "S3_UpLoad_ERRORS")
//			return
//		}
//		addr = strings.TrimRight(addr, "/")
//		results = append(results, LogoFileAddr{filename, addr})
//	}
//
//	this.Response(ctx, results)
//}
//
