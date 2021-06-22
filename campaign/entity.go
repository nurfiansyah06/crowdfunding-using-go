package campaign

import (
	"crowfunding/user"
	"time"
)


type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	BackerCount      int
	Perks            string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt		 time.Time
	CampaignImages   []CampaignImage
	User 			 user.User
}

type CampaignImage struct {
	ID 				int
	CampaignID 		int
	FileName 		string
	IsPrimary 		int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}