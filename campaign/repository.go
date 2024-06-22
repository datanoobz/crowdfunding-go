package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(UserID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("campaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserId(UserID int) ([]Campaign, error) {
	var campaigns []Campaign
	// "campaign_images.is_primary = 1"
	err := r.db.Where("user_id=?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
