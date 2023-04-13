package repository

import (
	"MyGram/model"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type commentRepository struct {
}

func NewCommentRepository() *commentRepository {
	return &commentRepository{}
}

func (repo *commentRepository) FindAllByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint) ([]model.Comments, error) {
	comments := []model.Comments{}
	if err := tx.WithContext(ctx).Where("photo_id = ?", photoID).Order("id DESC").Find(&comments).Error; err != nil {
		log.Println("Error finding all comment by photo:", err)
		return comments, err
	}
	return comments, nil
}

func (repo *commentRepository) FindByPhotoId(ctx context.Context, tx *gorm.DB, photoID uint, id uint) (model.Comments, error) {
	comments := model.Comments{}
	if err := tx.WithContext(ctx).Where("photo_id = ? AND id = ?", photoID, id).Take(&comments).Error; err != nil {
		log.Println("Error finding comment by photo:", err)
		return comments, err
	}
	return comments, nil
}

func (repo *commentRepository) Create(ctx context.Context, tx *gorm.DB, req model.RequestComments, userID, photoID uint) (model.Comments, error) {
	// comment.UserID = userID
	// comment.PhotoID = photoID
	comment := model.Comments{}
	comment.Message = req.Message
	comment.PhotoID = photoID
	comment.UserID = userID
	if err := tx.WithContext(ctx).Create(&comment).Error; err != nil {
		log.Println("Error creating comment:", err)
		return comment, err
	}
	return comment, nil
}

// func (repo *commentRepository) Update(ctx context.Context, tx *gorm.DB, req model.RequestComment, id, userID, photoID uint) (model.Comment, error) {
// 	comment := model.Comment{}
// 	if err := tx.WithContext(ctx).Model(&comment).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).Updates(model.Comment{GormModel: model.GormModel{ID: id}, Message: req.Message}).Error; err != nil {
// 		log.Printf("Error updating comment: %+v data doesn't match\n", err)
// 		return comment, err
// 	}
// 	// to return result after update
// 	tx.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Take(&comment)
// 	return comment, nil
// }

func (repo *commentRepository) Update(ctx context.Context, tx *gorm.DB, req model.RequestComments, id, userID, photoID uint) (model.Comments, error) {
	comment := model.Comments{}
	if err_ := tx.WithContext(ctx).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).First(&comment).Error; err_ != nil {
		if errors.Is(err_, gorm.ErrRecordNotFound) {
			return comment, fmt.Errorf("comment with ID %d in photo ID %d by user ID %d not found", id, photoID, userID)
		}
		return comment, fmt.Errorf("failed to update photo: %v", err_)
	}

	if err := tx.WithContext(ctx).Model(&comment).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).Updates(model.Comments{GormModel: model.GormModel{ID: id}, Message: req.Message}).Error; err != nil {
		log.Printf("Error updating comment: %+v\n", err)
		return comment, fmt.Errorf("failed to update photo: %v", err)
	}

	// to return result after update
	tx.WithContext(ctx).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).Take(&comment)
	return comment, nil
}

func (repo *commentRepository) Delete(ctx context.Context, tx *gorm.DB, id, userID, photoID uint) error {
	comment := model.Comments{}
	result := tx.WithContext(ctx).Where("id = ? AND user_id = ? AND photo_id = ?", id, userID, photoID).Delete(&comment)
	if result.Error != nil {
		log.Printf("Error deleting comment: %+v\n", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		log.Printf("Error deleting comment: data not found\n")
		return gorm.ErrRecordNotFound
	}
	return nil
}
