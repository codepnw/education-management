package handlers

import (
	"fmt"
	"net/http"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/usecases"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	usecase *usecases.CourseUsecase
}

func NewCourseHandler(usecase *usecases.CourseUsecase) *CourseHandler {
	return &CourseHandler{usecase: usecase}
}

func (ch *CourseHandler) CreateCourse(c *gin.Context) {
	course := new(entities.Course)
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := ch.usecase.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	course.ID = id
	c.JSON(http.StatusCreated, id)
}

func (ch *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := ch.usecase.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (ch *CourseHandler) GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	course, err  := ch.usecase.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (ch *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	course := new(entities.Course)

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ch.usecase.UpdateCourse(id, course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meesage": fmt.Sprintf("course id %s updated", id)})
}

func (ch *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	if err := ch.usecase.DeleteCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": fmt.Sprintf("course id %s deleted", id)})
}
