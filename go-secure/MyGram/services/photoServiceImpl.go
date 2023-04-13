package services

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type photoService struct {
	db        *gorm.DB
	repoPhoto repository.PhotoRepository
}

func NewPhotoService(db *gorm.DB, repoPhoto repository.PhotoRepository) *photoService {
	return &photoService{
		db:        db,
		repoPhoto: repoPhoto,
	}
}

func (service *photoService) FindAllByUserId(ctx context.Context, userID uint) ([]model.Photo, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photos, err := service.repoPhoto.FindAllByUserId(ctx, tx, userID); err != nil {
		log.Println("Data not found")
		return photos, err
	} else {
		return photos, nil
	}
}

func (service *photoService) FindByUserId(ctx context.Context, userID uint, id uint) (model.Photo, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photo, err := service.repoPhoto.FindByUserId(ctx, tx, userID, id); err != nil {
		log.Println("Data not found")
		return photo, err
	} else {
		return photo, nil
	}
}

func (service *photoService) Create(ctx context.Context, req model.RequestPhoto, userID uint) (model.Photo, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photo, err := service.repoPhoto.Create(ctx, tx, req, userID); err != nil {
		log.Println("Failed to create photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (service *photoService) Update(ctx context.Context, req model.RequestPhoto, id, userID uint) (model.Photo, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photo, err := service.repoPhoto.Update(ctx, tx, req, id, userID); err != nil {
		log.Println("Failed to update photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (service *photoService) Delete(ctx context.Context, id, userID uint) error {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if err := service.repoPhoto.Delete(ctx, tx, id, userID); err != nil {
		log.Println("Failed to delete photo")
		return err
	} else {
		return nil
	}
}
