package services

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type socialMediaService struct {
	db              *gorm.DB
	repoSocialMedia repository.SocialMediaRepository
}

func NewSocialMediaService(db *gorm.DB, repoSocialMedia repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{
		db:              db,
		repoSocialMedia: repoSocialMedia,
	}
}

func (service *socialMediaService) FindAllByUserId(ctx context.Context, userID uint) ([]model.SocialMedia, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if socialMedias, err := service.repoSocialMedia.FindAllByUserId(ctx, tx, userID); err != nil {
		log.Println("Data not found")
		return socialMedias, err
	} else {
		return socialMedias, nil
	}
}

func (service *socialMediaService) FindByUserId(ctx context.Context, userID uint, id uint) (model.SocialMedia, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if socialMedia, err := service.repoSocialMedia.FindByUserId(ctx, tx, userID, id); err != nil {
		log.Println("Data not found")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}

func (service *socialMediaService) Create(ctx context.Context, req model.RequestSocialMedia, userID uint) (model.SocialMedia, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if socialMedia, err := service.repoSocialMedia.Create(ctx, tx, req, userID); err != nil {
		log.Println("Failed to create socialMedia")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}

func (service *socialMediaService) Update(ctx context.Context, req model.RequestSocialMedia, id, userID uint) (model.SocialMedia, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if socialMedia, err := service.repoSocialMedia.Update(ctx, tx, req, id, userID); err != nil {
		log.Println("Failed to update socialMedia")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}

func (service *socialMediaService) Delete(ctx context.Context, id, userID uint) (model.SocialMedia, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if socialMedia, err := service.repoSocialMedia.Delete(ctx, tx, id, userID); err != nil {
		log.Println("Failed to delete socialMedia")
		return socialMedia, err
	} else {
		return socialMedia, nil
	}
}
