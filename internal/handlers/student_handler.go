package handlers

import (
	"fmt"
	"net/http"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/usecases"
	"github.com/codepnw/education/internal/utils"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	usecase *usecases.StudentUsecase
}

func NewStudentHandler(su *usecases.StudentUsecase) *StudentHandler {
	return &StudentHandler{usecase: su}
}

func (sh *StudentHandler) CreateStudent(c *gin.Context) {
	var student *entities.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dob, err := utils.IsValidDate(student.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student.DOB = dob

	if !utils.IsValidPhone(student.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid phone format: 10 numbers"})
		return
	}

	id, err := sh.usecase.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	student.ID = id
	c.JSON(http.StatusCreated, student)
}

func (sh *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := sh.usecase.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (sh *StudentHandler) GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	teacher, err := sh.usecase.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (sh *StudentHandler) UpdateStudentByID(c *gin.Context) {
	id := c.Param("id")

	var student *entities.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sh.usecase.UpdateStudentByID(id, student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student updated succesfully"})
}

func (sh *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := sh.usecase.DeleteStudent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("student id %v deleted", id)})
}
