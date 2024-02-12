package services

type ServiceContainer interface {
	ImageService() ImageService
}

type serviceContainer struct {
	imageService ImageService
}

func NewServiceContainer() ServiceContainer {
	return &serviceContainer{NewImageService()}
}

func (c *serviceContainer) ImageService() ImageService {
	return c.imageService
}
