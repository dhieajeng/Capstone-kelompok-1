package service

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type eventService struct {
	eventReposioty repository.EventRepository
}

func (s *eventService) GetAllEvent(c echo.Context) (*[]entity.Event, error) {
	events, err := s.eventReposioty.GetAll(c.Request().Context())
	if err != nil {
		return nil, err
	}

	return &events, nil
}

func (s *eventService) GetAllEventWithPaginateAndFilter(c echo.Context,
	paginate *binder.PaginateRequest,
	filter *binder.FilterRequest,
	sort *binder.SortRequest) (*map[string]interface{}, error) {
	events, totalItems, err := s.eventReposioty.GetAllWithPaginateAndFilter(c.Request().Context(), *paginate, *filter, *sort)
	if err != nil {
		return nil, err
	}

	totalPages := int((totalItems + int64(*paginate.Limit) - 1) / int64(*paginate.Limit))

	data := util.NewPagination(*paginate.Limit, *paginate.Page, int(totalItems), totalPages, events).Response()

	return &data, nil
}

func (s *eventService) CreateEvent(c echo.Context, event *entity.Event) (*entity.Event, error) {
	coverName, _ := s.handleFileUpload(c, "cover", "storage/event/covers")
	event.Cover = coverName

	logoName, _ := s.handleFileUpload(c, "organizer_logo", "storage/event/logos")
	event.OrganizerLogo = logoName

	eventRes, err := s.eventReposioty.Create(c.Request().Context(), event)
	if err != nil {
		return nil, err
	}

	return eventRes, nil
}

func (s *eventService) FindEventById(c echo.Context, id uuid.UUID) (*entity.Event, error) {
	return s.eventReposioty.FindById(c.Request().Context(), id)
}

func (s *eventService) FindEventDetailById(c echo.Context, id uuid.UUID) (*entity.Event, error) {
	return s.eventReposioty.FindWithDetailById(c.Request().Context(), id)
}

func (s *eventService) UpdateEvent(c echo.Context, event *entity.Event) (*entity.Event, error) {

	currentEvent, err := s.eventReposioty.FindById(c.Request().Context(), event.ID)
	if err != nil {
		return nil, err
	}

	var cover *string = currentEvent.Cover
	var organizerLogo *string = currentEvent.OrganizerLogo

	if cover != nil {
		s.handleFileDelete("storage/event/covers", *cover)
	}
	coverName, _ := s.handleFileUpload(c, "cover", "storage/event/covers")
	cover = coverName

	if organizerLogo != nil {
		s.handleFileDelete("storage/event/logos", *organizerLogo)
	}
	logoName, _ := s.handleFileUpload(c, "organizer_logo", "storage/event/logos")
	organizerLogo = logoName

	event.Cover = cover
	event.OrganizerLogo = organizerLogo
	eventRes, err := s.eventReposioty.Update(c.Request().Context(), event)
	if err != nil {
		return nil, err
	}
	return eventRes, nil
}

func (s *eventService) DeleteEvent(c echo.Context, id uuid.UUID) error {
	event, err := s.eventReposioty.FindById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	err = s.eventReposioty.Delete(c.Request().Context(), event)
	if err != nil {
		return err
	}
	return nil
}

func (s *eventService) handleFileUpload(c echo.Context, key, path string) (*string, error) {
	file, err := c.FormFile(key)
	if err != nil {
		return nil, err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s-%s", strconv.FormatInt(time.Now().Unix(), 10), file.Filename)
	fullPath := filepath.Join(currentDir, path, fileName)

	if err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return nil, err
	}

	fmt.Println(fullPath)

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return &fileName, nil
}

func (s *eventService) handleFileDelete(path, fileName string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fullPath := filepath.Join(currentDir, path, fileName)

	err = os.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}

type EventService interface {
	GetAllEvent(c echo.Context) (*[]entity.Event, error)
	GetAllEventWithPaginateAndFilter(c echo.Context, paginate *binder.PaginateRequest, filter *binder.FilterRequest, sort *binder.SortRequest) (*map[string]interface{}, error)
	CreateEvent(c echo.Context, event *entity.Event) (*entity.Event, error)
	FindEventById(c echo.Context, id uuid.UUID) (*entity.Event, error)
	FindEventDetailById(c echo.Context, id uuid.UUID) (*entity.Event, error)
	UpdateEvent(c echo.Context, event *entity.Event) (*entity.Event, error)
	DeleteEvent(c echo.Context, id uuid.UUID) error
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{
		eventReposioty: eventRepository,
	}
}
