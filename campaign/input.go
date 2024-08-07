package campaign

import "crowdfunding/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreatedCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type UploadCampaignImage struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
