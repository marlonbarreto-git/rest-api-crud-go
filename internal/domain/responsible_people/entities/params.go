package entities

type (
	ResponsiblesPageParams struct {
		Size int `form:"size" binding:"gte=1" default:"1"`
		Page int `form:"page" binding:"gte=1" default:"10"`
	}

	ResponsibleParams struct {
		Id int `uri:"id" binding:"required,gt=0"`
	}
)
