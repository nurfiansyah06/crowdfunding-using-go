package handler

import (
	"crowfunding/campaign"
	"crowfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}


//api/v1/campaigns

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get Campaigns",http.StatusBadRequest,"Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get Campaigns",http.StatusOK,"Success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}