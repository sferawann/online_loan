package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/usecase"
)

type TraCon struct {
	TraUsecase usecase.TraUsecase
}

func NewTraController(TraUsecase usecase.TraUsecase) *TraCon {
	return &TraCon{TraUsecase: TraUsecase}
}

func (c *TraCon) Insert(ctx *gin.Context) {
	insertTra := model.Transaction{}
	if err := ctx.ShouldBind(&insertTra); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newTra, err := c.TraUsecase.Save(insertTra)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transaction": newTra})
}

func (c *TraCon) FindAll(ctx *gin.Context) {
	Tras, err := c.TraUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transaction": Tras})
}

func (c *TraCon) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Tra, err := c.TraUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transaction": Tra})
}
