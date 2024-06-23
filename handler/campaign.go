package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// logic

// tangkap parameter di handler
// handler ke service
// service yang menentukan repository mana yang di-call
// repository : GetAll, getByUserId
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// v1/campaign
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	log.Println(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
