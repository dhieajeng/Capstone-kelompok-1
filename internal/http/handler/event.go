package handler

import (
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type EventHandler struct {
	eventService service.EventService
}

func (h *EventHandler) GetAllEvent(c echo.Context) error {
	paginateReq := new(binder.PaginateRequest)
	if err := c.Bind(paginateReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	page := h.getDefaultInt(paginateReq.Page, 1)
	limit := h.getDefaultInt(paginateReq.Limit, 10)
	paginate := &binder.PaginateRequest{
		Page:  &page,
		Limit: &limit,
	}

	filterReq := new(binder.FilterRequest)
	if err := c.Bind(filterReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	filter := &binder.FilterRequest{
		Keyword:  filterReq.Keyword,
		Location: filterReq.Location,
		Topic:    filterReq.Topic,
		Category: filterReq.Category,
		Time:     filterReq.Time,
		IsPaid:   filterReq.IsPaid,
	}

	sortReq := new(binder.SortRequest)
	if err := c.Bind(sortReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	sort := &binder.SortRequest{
		Sort: sortReq.Sort,
	}

	events, err := h.eventService.GetAllEventWithPaginateAndFilter(c, paginate, filter, sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK,
		true,
		"sukses menampilkan semua data event",
		events))
}

func (h *EventHandler) GetDetailEvent(c echo.Context) error {
	id := c.Param("id")
	eventId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	event, err := h.eventService.FindEventById(c, eventId)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error(http.StatusNotFound, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "sukses menampilkan detail event", event))
}

func (h *EventHandler) GetDetailEventWithTicket(c echo.Context) error {
	id := c.Param("id")
	eventId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	event, err := h.eventService.FindEventDetailById(c, eventId)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error(http.StatusNotFound, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "sukses menampilkan detail event", event))
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	req := new(binder.CreateEventRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err),
		))
	}

	dataUser, _ := c.Get("user").(*jwt.Token)
	userClaims := dataUser.Claims.(*jwt_token.JwtCustomClaims)

	layout := "2006-01-02 15:04:05"
	start, err := time.Parse(layout, req.Start)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(
			http.StatusUnprocessableEntity,
			false,
			err.Error(),
		))
	}
	end, err := time.Parse(layout, req.End)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(
			http.StatusUnprocessableEntity,
			false,
			err.Error(),
		))
	}

	eventParams := entity.NewEventParams{
		Name:             req.Name,
		UserID:           uuid.MustParse(userClaims.ID),
		LocationID:       req.LocationID,
		CategoryID:       req.CategoryID,
		TopicID:          req.TopicID,
		Start:            start,
		End:              end,
		Address:          req.Address,
		AddressLink:      req.AddressLink,
		Organizer:        req.Organizer,
		OrganizerLogo:    req.OrganizerLogo,
		Cover:            req.Cover,
		Description:      req.Description,
		TermAndCondition: req.TermAndCondition,
		IsPaid:           req.IsPaid,
		IsPublic:         req.IsPublic,
		IsApproved:       true,
	}
	eventDTO := entity.NewEvent(eventParams)

	event, err := h.eventService.CreateEvent(c, eventDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "berhasil menambahkan event", event))

}

func (h *EventHandler) EditEvent(c echo.Context) error {
	req := new(binder.EditEventRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err),
		))
	}

	layout := "2006-01-02 15:04:05"
	start, err := time.Parse(layout, req.Start)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(
			http.StatusUnprocessableEntity,
			false,
			err.Error(),
		))
	}
	end, err := time.Parse(layout, req.End)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(
			http.StatusUnprocessableEntity,
			false,
			err.Error(),
		))
	}

	eventParams := entity.EditEventParams{
		ID:               req.ID,
		Name:             req.Name,
		LocationID:       req.LocationID,
		CategoryID:       req.CategoryID,
		TopicID:          req.TopicID,
		Start:            start,
		End:              end,
		Address:          req.Address,
		AddressLink:      req.AddressLink,
		Organizer:        req.Organizer,
		Description:      req.Description,
		TermAndCondition: req.TermAndCondition,
		IsPaid:           req.IsPaid,
		IsPublic:         req.IsPublic,
	}
	eventDTO := entity.EditEvent(eventParams)

	event, err := h.eventService.UpdateEvent(c, eventDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "berhasil edit event", event))

}

func (h *EventHandler) DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	eventId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	err = h.eventService.DeleteEvent(c, eventId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "berhasil menghapus event", nil))
}

func (h *EventHandler) getDefaultInt(value *int, defaultValue int) int {
	if value != nil {
		return *value
	}
	return defaultValue
}

func NewEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{eventService: eventService}
}
