package cos

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	cos "github.com/tencentyun/cos-go-sdk-v5"
	"github.com/wen-flower/easy-douyin/pkg/constant"
)

var client *cos.Client

func Init() {
	u, _ := url.Parse(constant.BucketURL)
	su, _ := url.Parse("https://cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("TENCENT_COS_SECRET_ID"),  // 用户的 SecretId
			SecretKey: os.Getenv("TENCENT_COS_SECRET_KEY"), // 用户的 SecretKey
		},
	})
}

func Put(name string, reader io.Reader, uopts *cos.ObjectPutOptions) error {
	_, err := client.Object.Put(context.Background(), name, reader, nil)
	return err
}
