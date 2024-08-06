package handlers

import (
	"fmt"
	"net/http"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/usecases"
	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	usecase *usecases.ScheduleUsecase
}

func NewScheduleHandler(usecase *usecases.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{usecase: usecase}
}

func (sh *ScheduleHandler) CreateSchedule(c *gin.Context) {
	schedule := new(entities.Schedule)

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := sh.usecase.CreateSchedule(schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	schedule.ID = id
	c.JSON(http.StatusCreated, schedule)
}

func (sh *ScheduleHandler) GetAllSchedule(c *gin.Context) {
	schedule, err := sh.usecase.GetAllSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

func (sh *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	id := c.Param("id")
	schedule, err := sh.usecase.GetScheduleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

func (sh *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	req := new(entities.Schedule)

	if err := sh.usecase.UpdateSchedule(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("schedule id %s updated", id)})
}

func (sh *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	if err := sh.usecase.DeleteSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("schedule id %s deleted", id)})
}
