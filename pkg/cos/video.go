package cos

import (
	"context"
	"io"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/wen-flower/easy-douyin/pkg/constant"
)

func VideoUpload(name string, reader io.Reader) error {
	err := Put(constant.VideoPath+name, reader, nil)
	if err != nil {
		return err
	}
	opt := &cos.GetSnapshotOptions{
		Time: 1,
	}
	// 截取封面
	resp, err := client.CI.GetSnapshot(context.Background(), constant.VideoPath+name, opt)
	if err != nil {
		return err
	}

	err = Put(constant.VideoCoverPath+name, resp.Body, nil)
	if err != nil {
		return err
	}

	return err
}
