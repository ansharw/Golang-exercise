package repository

import (
	"MyGram/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type socialMediaRepository struct {
}

func NewSocialMediaRepository() *socialMediaRepository {
	return &socialMediaRepository{}
}

func (repo *socialMediaRepository) FindAllByUserId(ctx context.Context, tx *gorm.DB, userID uint) ([]model.SocialMedia, error) {
	socialMedias := []model.SocialMedia{}
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&socialMedias).Error; err != nil {
		log.Println("Error finding all social media:", err)
		return socialMedias, err
	}
	return socialMedias, nil
}

func (repo *socialMediaRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userID uint, id uint) (model.SocialMedia, error) {
	socialMedias := model.SocialMedia{}
	if err := tx.WithContext(ctx).Where("user_id = ? AND id = ?", userID, id).Take(&socialMedias).Error; err != nil {
		log.Println("Error finding social media:", err)
		return socialMedias, err
	}
	return socialMedias, nil
}

func (repo *socialMediaRepository) Create(ctx context.Context, tx *gorm.DB, req model.RequestSocialMedia, userID uint) (model.SocialMedia, error) {
	socialMedia := model.SocialMedia{}
	socialMedia.Name = req.Name
	socialMedia.SocialMediaURL = req.SocialMediaURL
	socialMedia.UserID = userID
	if err := tx.WithContext(ctx).Create(&socialMedia).Error; err != nil {
		log.Println("Error creating social media:", err)
		return socialMedia, err
	}
	return socialMedia, nil
}

func (repo *socialMediaRepository) Update(ctx context.Context, tx *gorm.DB, req model.RequestSocialMedia, id, userID uint) (model.SocialMedia, error) {
	socialMedia := model.SocialMedia{}
	if err := tx.WithContext(ctx).Model(&socialMedia).Where("id = ? AND user_id = ?", id, userID).Updates(model.SocialMedia{GormModel: model.GormModel{ID: id}, Name: req.Name, SocialMediaURL: req.SocialMediaURL}).Error; err != nil {
		log.Printf("Error updating social media: %+v data doesn't match\n", err)
		return socialMedia, err
	}
	// to return result after update
	tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Take(&socialMedia)
	return socialMedia, nil
}

func (repo *socialMediaRepository) Delete(ctx context.Context, tx *gorm.DB, id, userID uint) (model.SocialMedia, error) {
	socialMedia := model.SocialMedia{}
	if err := tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Take(&socialMedia).Error; err != nil {
		log.Printf("Error deleting social media: %+v data doesn't match\n", err)
		return socialMedia, err
	}
	return socialMedia, nil
}
