package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/usecase"
)

type AcceptStatusCon struct {
	AcceptStatusUsecase usecase.AcceptStatusUsecase
}

func NewAcceptStatusController(AcceptStatusUsecase usecase.AcceptStatusUsecase) *AcceptStatusCon {
	return &AcceptStatusCon{AcceptStatusUsecase: AcceptStatusUsecase}
}

func (c *AcceptStatusCon) Insert(ctx *gin.Context) {
	insertAcceptStatus := model.AcceptStatus{}
	if err := ctx.ShouldBind(&insertAcceptStatus); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newAcceptStatus, err := c.AcceptStatusUsecase.Save(insertAcceptStatus)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"AcceptStatus": newAcceptStatus})
}

func (c *AcceptStatusCon) FindAll(ctx *gin.Context) {
	AcceptStatuss, err := c.AcceptStatusUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"AcceptStatuss": AcceptStatuss})
}

func (c *AcceptStatusCon) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	AcceptStatus, err := c.AcceptStatusUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"AcceptStatus": AcceptStatus})
}

func (c *AcceptStatusCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.AcceptStatus{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAcceptStatus, err := c.AcceptStatusUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"AcceptStatus": updatedAcceptStatus})
}

func (c *AcceptStatusCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.AcceptStatusUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted AcceptStatus!"})
}
