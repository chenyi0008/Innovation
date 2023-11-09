package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/sms/bytes"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"mime/multipart"
)

const (
	AccessKey = "JoSuSYH6BRs-90V7ySNfFHxZz20JJZgXAfMciKQR"
	SecretKey = "F7RkSFCVsmNqq6QFzwU15OTregxxQA2Bj5tR2tfN"
	Bucket    = "innovation-chenyi"
)

var (
	cfg           storage.Config
	mac           *qbox.Mac
	bucketManager *storage.BucketManager
)

func QiniuConfigInit() {
	cfg = storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	mac = qbox.NewMac(AccessKey, SecretKey)
	bucketManager = storage.NewBucketManager(mac, &cfg)
}

func QiniuUpload(key string, file []byte) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}

	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:file": "video",
		},
	}

	dataLen := int64(len(file))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(file), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

func ConvertFileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容到字节数组
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func QiniuDownload(key string) string {
	//// 生成私有下载链接
	//url := storage.MakePrivateURL(mac, "s3uf295e4.hn-bkt.clouddn.com", key, 3600)
	//// 生成共有下载链接
	url := storage.MakePublicURL("s3uf295e4.hn-bkt.clouddn.com", key)
	return url

}
