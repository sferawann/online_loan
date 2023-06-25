package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/usecase"
)

type PaymentCon struct {
	PaymentUsecase usecase.PaymentUsecase
}

func NewPaymentController(PaymentUsecase usecase.PaymentUsecase) *PaymentCon {
	return &PaymentCon{PaymentUsecase: PaymentUsecase}
}

func (c *PaymentCon) Insert(ctx *gin.Context) {
	insertPayment := model.Payment{}
	if err := ctx.ShouldBind(&insertPayment); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPayment, err := c.PaymentUsecase.Save(insertPayment)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Payment": newPayment})
}

func (c *PaymentCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.Payment{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPayment, err := c.PaymentUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Payment": updatedPayment})
}

func (c *PaymentCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.PaymentUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted Payment!"})
}

func (c *PaymentCon) FindAll(ctx *gin.Context) {
	Payments, err := c.PaymentUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Payments": Payments})
}

func (c *PaymentCon) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Payment, err := c.PaymentUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Payment": Payment})
}