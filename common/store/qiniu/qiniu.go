package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func getToken(accessKey, secretKey, bucket string) string {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	// 获取上传凭证
	return putPolicy.UploadToken(mac)
}
func UploadToQiNiu(ctx context.Context, accessKey, secretKey string, data []byte, filetype, bucket string) (url string, err error) {
	upToken := getToken(accessKey, secretKey, bucket)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	rd := bytes.NewReader(data)

	key := fmt.Sprintf("%s.%s", uuid.New().String(), filetype)
	if err = formUploader.Put(
		ctx,
		&ret, upToken,
		key, rd,
		rd.Size(), &putExtra,
	); err != nil {
		return
	}
	url = ret.Key // 返回上传后的文件访问路径
	return
}
