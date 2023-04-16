package repository

import (
	"MyGram/model"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type photoRepository struct {
}

func NewPhotoRepository() *photoRepository {
	return &photoRepository{}
}

func (repo *photoRepository) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]model.Photo, error) {
	photos := []model.Photo{}
	tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&photos)
	if len(photos) == 0 {
		log.Println("Error finding all photo")
		return photos, errors.New("no records found")
	}
	return photos, nil
}

func (repo *photoRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (model.Photo, error) {
	photos := model.Photo{}
	if err := tx.WithContext(ctx).Where("user_id = ? AND id = ?", userID, id).Take(&photos).Error; err != nil {
		log.Println("Error finding photo:", err)
		return photos, err
	}
	return photos, nil
}

func (repo *photoRepository) Create(ctx context.Context, tx *gorm.DB, req model.RequestPhoto, userID uint) (model.Photo, error) {
	// photo.UserID = userID
	photo := model.Photo{}
	photo.Title = req.Title
	photo.PhotoURL = req.PhotoURL
	photo.Caption = req.Caption
	photo.UserID = userID
	if err := tx.WithContext(ctx).Create(&photo).Error; err != nil {
		log.Println("Error creating photo:", err)
		return photo, err
	}
	return photo, nil
}

func (repo *photoRepository) Update(ctx context.Context, tx *gorm.DB, req model.RequestPhoto, id, userID uint) (model.Photo, error) {
	photo := model.Photo{}
	if err := tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&photo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photo, fmt.Errorf("photo with ID %d and user ID %d not found", id, userID)
		}
		return photo, fmt.Errorf("failed to update photo: %v", err)
	}

	if err := tx.WithContext(ctx).Model(&photo).Where("id = ? AND user_id = ?", id, userID).Updates(model.Photo{GormModel: model.GormModel{ID: id}, Title: req.Title, Caption: req.Caption, PhotoURL: req.PhotoURL}).Error; err != nil {
		return photo, fmt.Errorf("failed to update photo: %v", err)
	}

	return photo, nil
}

func (repo *photoRepository) Delete(ctx context.Context, tx *gorm.DB, id, userID uint) error {
	photo := model.Photo{}
	result := tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&photo)
	if result.Error != nil {
		log.Printf("Error deleting social media: %+v\n", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		log.Printf("Error deleting social media: data not found\n")
		return gorm.ErrRecordNotFound
	}
	return nil
}
