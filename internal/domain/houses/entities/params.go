package entities

type (
	HousesPageParams struct {
		Size int `form:"size" binding:"gte=1" default:"1"`
		Page int `form:"page" binding:"gte=1" default:"10"`
	}

	HouseParams struct {
		Id string `uri:"id" binding:"required"`
	}
)
