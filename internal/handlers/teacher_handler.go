package handlers

import (
	"fmt"
	"net/http"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/usecases"
	"github.com/codepnw/education/internal/utils"
	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	usecase *usecases.TeacherUsecase
}

func NewTeacherHandler(usecase *usecases.TeacherUsecase) *TeacherHandler {
	return &TeacherHandler{usecase: usecase}
}

func (th *TeacherHandler) CreateTeacher(c *gin.Context) {
	teacher := new(entities.Teacher)

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dob, err := utils.IsValidDate(teacher.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teacher.DOB = dob

	if !utils.IsValidPhone(teacher.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid phone format: 10 numbers"})
		return
	}

	id, err := th.usecase.CreateTeacher(teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	teacher.ID = id
	c.JSON(http.StatusCreated, teacher)
}

func (th *TeacherHandler) GetAllTeachers(c *gin.Context) {
	teachers, err := th.usecase.GetAllTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

func (th *TeacherHandler) GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	teacher, err := th.usecase.GetTeacherByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (th *TeacherHandler) UpdateTeacher(c *gin.Context) {
	id := c.Param("id")

	teacher := new(entities.Teacher)
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.usecase.UpdateTeacher(id, teacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "teacher updated succesfully"})
}

func (th *TeacherHandler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")

	if err := th.usecase.DeleteTeacher(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("teacher id %v deleted", id)})
}
