package controller

import (
	"net/http"
	"projek-pertama/model"
	"projek-pertama/repository"
	"projek-pertama/util"
	"time"

	"github.com/gin-gonic/gin"
)

type personController struct {
	personRepository repository.IpersonRepository
}

func NewPersonController(personRepository repository.IpersonRepository) *personController {

	return &personController{
		personRepository: personRepository,
	}
}

func (pc *personController) Create(ctx *gin.Context) {
	var newPerson model.Person

	err := ctx.ShouldBindJSON(&newPerson)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdPerson, err := pc.personRepository.Create(newPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdPerson, ""))
}

// GetAll Person godoc
// @Summary Get All Person
// @Schemes
// @Description get all Person
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} []model.Person
// @Router /person [get]
func (pc *personController) GetAll(ctx *gin.Context) {
	time.Sleep(5 * time.Second)

	persons := make([]model.Person, 0)

	persons, err := pc.personRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, persons, ""))
}

func (pc *personController) Delete(ctx *gin.Context) {

	idString := ctx.Param("id")
	err := pc.personRepository.Delete(idString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, ""))
}
