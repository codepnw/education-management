package handlers

import (
	"fmt"
	"net/http"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/usecases"
	"github.com/gin-gonic/gin"
)

type ClassroomHandler struct {
	usecase *usecases.ClassroomUsecase
}

func NewClassroomHandler(usecase *usecases.ClassroomUsecase) *ClassroomHandler {
	return &ClassroomHandler{usecase: usecase}
}

func (ch *ClassroomHandler) CreateClassroom(c *gin.Context) {
	var classroom *entities.Classroom

	if err := c.ShouldBindJSON(&classroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := ch.usecase.CreateClassroom(classroom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	classroom.ID = id
	c.JSON(http.StatusCreated, classroom)
}

func (ch *ClassroomHandler) GetAllClassroom(c *gin.Context) {
	classrooms, err := ch.usecase.GetAllClassroom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classrooms)
}

func (ch *ClassroomHandler) GetClassroomByID(c *gin.Context) {
	id := c.Param("id")
	classroom, err := ch.usecase.GetClassroomByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classroom)
}

func (ch *ClassroomHandler) UpdateClassroom(c *gin.Context) {
	id := c.Param("id")

	var classroom *entities.Classroom
	if err := c.ShouldBindJSON(&classroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ch.usecase.UpdateClassroom(id, classroom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "classroom updated succesfully"})
}

func (ch *ClassroomHandler) DeleteClassroom(c *gin.Context) {
	id := c.Param("id")
	err := ch.usecase.DeleteClassroom(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": fmt.Sprintf("classroom id %s deleted", id)})
}
