package entities

type (
	ResponsiblesPageParams struct {
		Size int `form:"size" binding:"gte=1" default:"1"`
		Page int `form:"page" binding:"gte=1" default:"10"`
	}

	ResponsibleParams struct {
		ResponsibleID int `uri:"responsibleId" binding:"required,gt=0"`
		PersonID      int `uri:"personId" binding:"omitempty,gt=0"`
	}
)
