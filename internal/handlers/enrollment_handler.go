package handlers

import (
	"fmt"
	"net/http"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/usecases"
	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	usecase *usecases.EnrollmentUsecase
}

func NewEnrollmentHandler(usecase *usecases.EnrollmentUsecase) *EnrollmentHandler {
	return &EnrollmentHandler{usecase: usecase}
}

func (eh *EnrollmentHandler) CreateEnroll(c *gin.Context) {
	req := new(entities.Enrollment)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := eh.usecase.CreateEnroll(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.ID = id
	c.JSON(http.StatusCreated, req)
}

func (eh *EnrollmentHandler) GetAllEnroll(c *gin.Context) {
	enrollments, err := eh.usecase.GetAllEnroll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollments)
}

func (eh *EnrollmentHandler) GetEnrollByID(c *gin.Context) {
	id := c.Param("id")
	enrollment, err := eh.usecase.GetEnrollByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}

func (eh *EnrollmentHandler) UpdateEnroll(c *gin.Context) {
	id := c.Param("id")
	req := new(entities.Enrollment)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := eh.usecase.UpdateEnroll(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("enrollment id %s updated", id)})
}

func (eh *EnrollmentHandler) DeleteEnroll(c *gin.Context) {
	id := c.Param("id")
	err := eh.usecase.DeleteEnroll(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": fmt.Sprintf("enrollment id %s deleted", id )})
}
