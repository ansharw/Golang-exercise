package repository

import (
	"MyGram/model"
	"context"
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
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&photos).Error; err != nil {
		log.Println("Error finding all photo:", err)
		return photos, err
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

func (repo *photoRepository) Create(ctx context.Context, tx *gorm.DB, photo model.Photo, userID uint) (model.Photo, error) {
	photo.UserID = userID
	if err := tx.WithContext(ctx).Create(&photo).Error; err != nil {
		log.Println("Error creating photo:", err)
		return photo, err
	}
	return photo, nil
}

func (repo *photoRepository) Update(ctx context.Context, tx *gorm.DB, photo model.Photo, id, userID uint) (model.Photo, error) {
	if err := tx.WithContext(ctx).Model(&photo).Where("id = ? AND user_id = ?", id, userID).Updates(model.Photo{GormModel: model.GormModel{ID: id}, Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL}).Error; err != nil {
		log.Printf("Error updating photo: %+v data doesn't match\n", err)
		return photo, err
	}
	// to return result after update
	tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Take(&photo)
	return photo, nil
}

func (repo *photoRepository) Delete(ctx context.Context, tx *gorm.DB, id, userID uint) (model.Photo, error) {
	photo := model.Photo{}
	if err := tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Take(&photo).Error; err != nil {
		log.Printf("Error deleting photo: %+v data doesn't match\n", err)
		return photo, err
	}
	return photo, nil
}
