package downloader

import (
	"encoding/json"
	"io"
	"net/http"
)

const CommonUrl = "https://api.bilibili.com/x/web-interface/view?bvid="

// InfoRequest 这里的BvIds必须为大写，方便在main文件外使用，传参使用
type InfoRequest struct {
	BvIds []string
}

// VideoInfo 特别注意这里的json的字段，哪怕这里的`date`变为`data`都会让程序出错
// 这里要结合API接口返回的json来复核
type VideoInfo struct {
	Code int `json:"code"`
	Date struct {
		BvId  string `json:"bvid"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"data"`
}

type InfoResponse struct {
	Infos []VideoInfo
}

func BatchDownloadVideoInfo(request InfoRequest) (InfoResponse, error) {
	var videoInfo VideoInfo
	var response InfoResponse
	for _, bvid := range request.BvIds {
		res, err := http.Get(CommonUrl + bvid)
		if err != nil {
			return InfoResponse{}, err
		}
		resBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return InfoResponse{}, err
		}
		err = json.Unmarshal(resBytes, &videoInfo)
		if err != nil {
			return InfoResponse{}, err
		}
		if err := res.Body.Close(); err != nil {
			return InfoResponse{}, err
		}
		response.Infos = append(response.Infos, videoInfo)

	}
	return response, nil
}
