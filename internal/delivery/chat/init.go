package chat

import (
	"github.com/gin-gonic/gin"

	errorCommon "github.com/aziemp66/freya-be/common/error"
	httpCommon "github.com/aziemp66/freya-be/common/http"

	chatUsecase "github.com/aziemp66/freya-be/internal/usecase/chat"
)

type ChatDeliveryImplementation struct {
	chatUsecase chatUsecase.Usecase
}

func NewChatDeliveryImplementation(router *gin.RouterGroup, chatUsecase chatUsecase.Usecase) *ChatDeliveryImplementation {
	chatDelivery := &ChatDeliveryImplementation{
		chatUsecase: chatUsecase,
	}

	router.POST("/appointment", chatDelivery.CreateAppointment)
	router.GET("/appointment/:id", chatDelivery.GetAppointmentByID)
	router.GET("/appointment", chatDelivery.GetAllAppointmentByUserID)
	router.PUT("/appointment/:id", chatDelivery.UpdateAppointmentStatus)

	return chatDelivery
}

func (d *ChatDeliveryImplementation) CreateAppointment(c *gin.Context) {
	var appointmentRequest httpCommon.AddAppointment

	if err := c.ShouldBindJSON(&appointmentRequest); err != nil {
		c.Error(err)
		return
	}

	userId := c.GetString("user_id")

	err := d.chatUsecase.InsertAppointment(c, appointmentRequest.PsychologistId, userId, appointmentRequest.Date)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Create appointment success",
	})
}

func (d *ChatDeliveryImplementation) GetAppointmentByID(c *gin.Context) {
	appointmentId := c.Param("id")

	appointment, err := d.chatUsecase.FindAppointmentByID(c, appointmentId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Get appointment success",
		Value:   appointment,
	})
}

func (d *ChatDeliveryImplementation) GetAllAppointmentByUserID(c *gin.Context) {
	userId := c.GetString("user_id")

	userRole := c.GetString("user_role")

	var appointments []httpCommon.Appointment
	var err error
	if userRole == "psychologist" {
		appointments, err = d.chatUsecase.FindAppointmentByPsychologistID(c, userId)
	} else if userRole == "base" {
		appointments, err = d.chatUsecase.FindAppointmentByUserID(c, userId)
	} else {
		c.Error(errorCommon.NewForbiddenError("You are not allowed to access this resource"))
		return
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Get all appointment success",
		Value:   appointments,
	})
}

func (d *ChatDeliveryImplementation) UpdateAppointmentStatus(c *gin.Context) {
	appointmentId := c.Param("id")

	var appointmentRequest httpCommon.AppointmentStatusUpdate

	if err := c.ShouldBindJSON(&appointmentRequest); err != nil {
		c.Error(err)
		return
	}

	appointment, err := d.chatUsecase.FindAppointmentByID(c, appointmentId)

	if err != nil {
		c.Error(err)
		return
	}

	err = d.chatUsecase.UpdateAppointmentStatus(c, appointmentId, appointmentRequest.Status)

	if err != nil {
		c.Error(err)
		return
	}

	if appointmentRequest.Status == httpCommon.APPOINTMENTACCEPTED {
		err = d.chatUsecase.InsertChatroom(c, appointmentId, appointment.PsychologistId, appointment.UserId)

		if err != nil {
			c.Error(err)
			return
		}
	} else if appointmentRequest.Status == httpCommon.APPOINTMENTCANCELED {
		err = d.chatUsecase.DeleteChatroom(c, appointmentId)

		if err != nil {
			c.Error(err)
			return
		}
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Update appointment success",
	})
}
