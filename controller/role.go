package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/usecase"
)

type RoleCon struct {
	RoleUsecase usecase.RoleUsecase
}

func NewRoleController(RoleUsecase usecase.RoleUsecase) *RoleCon {
	return &RoleCon{RoleUsecase: RoleUsecase}
}

func (c *RoleCon) Insert(ctx *gin.Context) {
	insertRole := model.Role{}
	if err := ctx.ShouldBind(&insertRole); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newRole, err := c.RoleUsecase.Save(insertRole)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Role": newRole})
}

func (c *RoleCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.Role{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRole, err := c.RoleUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Role": updatedRole})
}

func (c *RoleCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.RoleUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted Role!"})
}

func (c *RoleCon) FindAll(ctx *gin.Context) {
	Roles, err := c.RoleUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Roles": Roles})
}

func (c *RoleCon) FindById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Role, err := c.RoleUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Role": Role})
}

func (c *RoleCon) FindByName(ctx *gin.Context) {
	RolenameParam := ctx.Param("name")

	Role, err := c.RoleUsecase.FindByName(RolenameParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Role": Role})
}
