package restaurantmodel

import (
	"errors"
	"food_delivery/common"
	"food_delivery/modules/user/usermodel"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel                    `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Owner_Id        int                `json:"-" gorm:"column:owner_id;"`
	FakeOwnerId 	*common.UID		   `json:"owner_id"`
	Addr            string             `json:"address" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	Owner           *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikedCount      int                `json:"liked_count" gorm:"liked_count"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	UserId          int            `json:"owner_id" gorm:"column:owner_id;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
	User            *usermodel.User `json:"user" gorm:"preload:false;foreignKey:UserId;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {

	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}

func (data *RestaurantCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *Restaurant) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeRestaurant)
	
	if u := data.Owner; u != nil {
		u.Mask(common.DbTypeUser)
	}

	fakeId := common.NewUID(uint32(data.Owner_Id), common.DbTypeUser, 1)
	data.FakeOwnerId = &fakeId
}

