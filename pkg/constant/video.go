package constant

import (
	"strconv"
)

const BucketURL = "https://simple-douyin-1256498121.cos.ap-shanghai.myqcloud.com"

const VideoPath = "video/"
const VideoCoverPath = "video/cover/"

const FullVideoPath = BucketURL + "/" + VideoPath
const FullVideoCoverPath = BucketURL + "/" + VideoCoverPath

// ParseVideoCoverUrl 通过视频ID 解析视频封面 URL
func ParseVideoCoverUrl(vid int64) string {
	return FullVideoCoverPath + strconv.FormatInt(vid, 10)
}

// ParseVideoUrl 通过视频ID 解析视频 URL
func ParseVideoUrl(vid int64) string {
	return FullVideoPath + strconv.FormatInt(vid, 10)
}
