package service

import (
	"github.com/superiqbal7/golang-gin-crud-rest-api/entity"
	"github.com/superiqbal7/golang-gin-crud-rest-api/repository"
)

type VideoService interface {
	Save(entity.Video) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	repository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (service *videoService) Save(video entity.Video) error {
	service.repository.Save(video)
	return nil
}

func (service *videoService) Update(video entity.Video) error {
	service.repository.Update(video)
	return nil
}

func (service *videoService) Delete(video entity.Video) error {
	service.repository.Delete(video)
	return nil
}

func (service *videoService) FindAll() []entity.Video {
	return service.repository.FindAll()
}

// package service

// import "github.com/superiqbal7/golang-gin-crud-rest-api/entity"

// type VideoService interface {
// 	Save(entity.Video) entity.Video
// 	FindAll() []entity.Video
// }

// type videoService struct {
// 	videos []entity.Video
// }

// func New() VideoService {
// 	return &videoService{}
// }

// func (service *videoService) Save(video entity.Video) entity.Video {
// 	service.videos = append(service.videos, video)
// 	return video
// }

// func (service *videoService) FindAll() []entity.Video {
// 	return service.videos
// }
