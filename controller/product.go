package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/usecase"
)

type ProductCon struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(ProductUsecase usecase.ProductUsecase) *ProductCon {
	return &ProductCon{ProductUsecase: ProductUsecase}
}

func (c *ProductCon) Insert(ctx *gin.Context) {
	insertProduct := model.Product{}
	if err := ctx.ShouldBind(&insertProduct); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newProduct, err := c.ProductUsecase.Save(insertProduct)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Product": newProduct})
}

func (c *ProductCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.Product{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := c.ProductUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Product": updatedProduct})
}

func (c *ProductCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.ProductUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted Product!"})
}

func (c *ProductCon) FindAll(ctx *gin.Context) {
	Products, err := c.ProductUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Products": Products})
}

func (c *ProductCon) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product, err := c.ProductUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Product": Product})
}

func (c *ProductCon) FindByName(ctx *gin.Context) {
	ProductnameParam := ctx.Param("name")

	Product, err := c.ProductUsecase.FindByName(ProductnameParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Product": Product})
}
