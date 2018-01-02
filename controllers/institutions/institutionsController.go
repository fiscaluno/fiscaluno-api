package institutions

import (
	institution "github.com/fiscaluno/fiscaluno-api/models/institution"
	"github.com/kataras/iris"
)

func AllInstitutions(ctx iris.Context) {
	ctx.JSON(institution.All())
}
