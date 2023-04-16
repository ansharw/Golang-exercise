package repository

import (
	"MyGram/model"
	"context"
	"errors"
	"fmt"
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
	tx.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&socialMedias)
	if len(socialMedias) == 0 {
		log.Println("Error finding all social media")
		return socialMedias, errors.New("no records found")
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
	if err := tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&socialMedia).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return socialMedia, fmt.Errorf("social media with ID %d and user ID %d not found", id, userID)
		}
		return socialMedia, fmt.Errorf("failed to update social media: %v", err)
	}

	if err := tx.WithContext(ctx).Model(&socialMedia).Where("id = ? AND user_id = ?", id, userID).Updates(model.SocialMedia{GormModel: model.GormModel{ID: id}, Name: req.Name, SocialMediaURL: req.SocialMediaURL}).Error; err != nil {
		return socialMedia, fmt.Errorf("failed to update social media: %v", err)
	}
	return socialMedia, nil
}

func (repo *socialMediaRepository) Delete(ctx context.Context, tx *gorm.DB, id, userID uint) error {
	socialMedia := model.SocialMedia{}
	result := tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&socialMedia)
	if result.Error != nil {
		log.Printf("Error deleting social media: %+v\n", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		log.Printf("Error deleting social media: data not found\n")
		return gorm.ErrRecordNotFound
	}
	return nil
}
