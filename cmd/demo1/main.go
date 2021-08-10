// 新建一个downloader的package包，里面有一个对外发布的BatchDownloadVideoInfo函数
// 这个函数可以返回一个视频的相关信息
// www.bilibili.com网站的视频信息查询接口如下https://api.bilibili.com/x/web-interface/view?bvid=BV1mP4y147dc
package main

import (
	"fmt"
	"gitee.com/ethancheng/regular_demo/internal/pkg/downloader"
)

func main() {
	request := downloader.InfoRequest{BvIds: []string{"BV1mP4y147dc", "BV1Ff4y187q9"}}
	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}
	for _, i := range response.Infos {
		fmt.Println("-----------------")
		fmt.Printf("B站视频: 标题: %s 的描述是: %s \n", i.Date.Title, i.Date.Desc)
		fmt.Println("-----------------")
	}

}
