package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/usecase"
)

type PaymentMethodCon struct {
	PaymentMethodUsecase usecase.PaymentMethodUsecase
}

func NewPaymentMethodController(PaymentMethodUsecase usecase.PaymentMethodUsecase) *PaymentMethodCon {
	return &PaymentMethodCon{PaymentMethodUsecase: PaymentMethodUsecase}
}

func (c *PaymentMethodCon) Insert(ctx *gin.Context) {
	insertPaymentMethod := model.PaymentMethod{}
	if err := ctx.ShouldBind(&insertPaymentMethod); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPaymentMethod, err := c.PaymentMethodUsecase.Save(insertPaymentMethod)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"PaymentMethod": newPaymentMethod})
}

func (c *PaymentMethodCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.PaymentMethod{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPaymentMethod, err := c.PaymentMethodUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"PaymentMethod": updatedPaymentMethod})
}

func (c *PaymentMethodCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.PaymentMethodUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted PaymentMethod!"})
}

func (c *PaymentMethodCon) FindAll(ctx *gin.Context) {
	PaymentMethods, err := c.PaymentMethodUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"PaymentMethods": PaymentMethods})
}

func (c *PaymentMethodCon) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PaymentMethod, err := c.PaymentMethodUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"PaymentMethod": PaymentMethod})
}

func (c *PaymentMethodCon) FindByName(ctx *gin.Context) {
	PaymentMethodnameParam := ctx.Param("name")

	PaymentMethod, err := c.PaymentMethodUsecase.FindByName(PaymentMethodnameParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"PaymentMethod": PaymentMethod})
}
