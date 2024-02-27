package utils

import (
	"Innovation/define"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"path"
)

func GetUUID() string {
	return uuid.NewV4().String()
}

// CosUpload 文件上传到腾讯云
func CosUpload(r *gin.Context) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, err := r.FormFile("data")
	fileHeader := file.Filename
	//file, fileHeader, err := r.FormFile("data")
	// Ext取后缀 扩展名 extension
	key := define.TencentFilePrefix + GetUUID() + path.Ext(fileHeader)

	if err != nil {
		panic(err)
	}

	open, err := file.Open()
	if err != nil {
		panic(err)
	}

	_, err = client.Object.Put(
		context.Background(), key, open, nil,
	)

	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}

// FileUploadToJpg 上传图片
func FileUploadToJpg(data *[]byte) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	// Ext取后缀 扩展名 extension
	key := define.TencentFilePrefix + GetUUID() + ".jpg"

	_, err := client.Object.Put(
		context.Background(), key, bytes.NewReader(*data), nil,
	)

	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}

func FileUpload(data *[]byte, extension string) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	// Ext取后缀 扩展名 extension
	key := define.TencentFilePrefix + GetUUID() + extension

	_, err := client.Object.Put(
		context.Background(), key, bytes.NewReader(*data), nil,
	)

	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}
