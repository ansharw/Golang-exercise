package repository

import (
	"MyGram/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type commentRepository struct {
}

func NewCommentRepository() *commentRepository {
	return &commentRepository{}
}

func (repo *commentRepository) FindAllByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint) ([]model.Comment, error) {
	comments := []model.Comment{}
	if err := tx.WithContext(ctx).Where("photo_id = ?", photoID).Order("id DESC").Find(&comments).Error; err != nil {
		log.Println("Error finding all comment by photo:", err)
		return comments, err
	}
	return comments, nil
}

func (repo *commentRepository) FindByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint, id uint) (model.Comment, error) {
	comments := model.Comment{}
	if err := tx.WithContext(ctx).Where("photo_id = ? AND id = ?", photoID, id).Take(&comments).Error; err != nil {
		log.Println("Error finding comment by photo:", err)
		return comments, err
	}
	return comments, nil
}

func (repo *commentRepository) Create(ctx context.Context, tx *gorm.DB, comment model.Comment, userID, photoID uint) (model.Comment, error) {
	comment.UserID = userID
	comment.PhotoID = photoID
	if err := tx.WithContext(ctx).Create(&comment).Error; err != nil {
		log.Println("Error creating comment:", err)
		return comment, err
	}
	return comment, nil
}

func (repo *commentRepository) Update(ctx context.Context, tx *gorm.DB, comment model.Comment, id, userID, photoID uint) (model.Comment, error) {
	if err := tx.WithContext(ctx).Model(&comment).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).Updates(model.Comment{GormModel: model.GormModel{ID: id}, Message: comment.Message}).Error; err != nil {
		log.Printf("Error updating comment: %+v data doesn't match\n", err)
		return comment, err
	}
	// to return result after update
	tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Take(&comment)
	return comment, nil
}

func (repo *commentRepository) Delete(ctx context.Context, tx *gorm.DB, id, userID, photoID uint) (model.Comment, error) {
	comment := model.Comment{}
	if err := tx.WithContext(ctx).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).Take(&comment).Error; err != nil {
		log.Printf("Error deleting comment: %+v data doesn't match\n", err)
		return comment, err
	}
	return comment, nil
}
