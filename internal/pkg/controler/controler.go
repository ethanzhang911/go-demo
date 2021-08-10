package controllers

import (
	"gitee.com/ethancheng/regular_demo/internal/pkg/model"
	"github.com/gin-gonic/gin"
	"sync"
)

type idGeneraget struct {
	mtx     sync.Mutex
	counter int64
}

var g idGeneraget

func (g *idGeneraget) getNextID() int64 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter++
	return g.counter
}

type VideoControler interface {
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
}

type controler struct {
	videos []model.Video
}

func NewController() VideoControler {
	return &controler{videos: make([]model.Video, 0)}
}

func (a *controler) GetAll(c *gin.Context) {
	c.JSON(200, a.videos)
}

func (a *controler) Update(c *gin.Context) {
	var videoToUpdate model.Video
	err := c.ShouldBindUri(&videoToUpdate)
	if err != nil {
		c.String(400, "bad request %v", err)
		return
	}
	err = c.ShouldBindJSON(&videoToUpdate)
	if err != nil {
		c.String(400, "bad request %v", err)
		return
	}
	for id, video := range a.videos {
		if video.Id == videoToUpdate.Id {
			a.videos[id] = videoToUpdate
		}
		c.String(200, " ID :%d has been updated", videoToUpdate.Id)
		return
	}
	c.String(400, "id not found %d", videoToUpdate.Id)
}

func (a *controler) Delete(c *gin.Context) {
	panic("implement me")
}

func (a *controler) Create(c *gin.Context) {
	video := model.Video{Id: g.getNextID()}
	if err := c.ShouldBindJSON(&video); err != nil {
		c.String(400, "bad request %v", err)
		return
	}
	a.videos = append(a.videos, video)
	c.String(200, "new video has been created,new id is %d,new name is %s,new age is %d", video.Id, video.Name, video.Age)
}
