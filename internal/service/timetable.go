package service

import (
	"errors"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type timetableService struct {
	timetableRepo repository.TimetableRepository
}

func (s *timetableService) FindById(c echo.Context, id uuid.UUID) (*entity.Timetable, error) {
	timetable, err := s.timetableRepo.FindById(c.Request().Context(), &id)
	if err != nil {
		return nil, err
	}

	return timetable, err
}

func (s *timetableService) FindByIds(c echo.Context, ids []uuid.UUID) ([]entity.Timetable, error) {
	timetable, err := s.timetableRepo.FindByIds(c.Request().Context(), ids)
	if err != nil {
		return nil, err
	}

	return timetable, err
}

func (s *timetableService) UpdateTicketStock(c echo.Context, ticketCounts map[uuid.UUID]int32, isDecrease bool) error {
	ids := make([]uuid.UUID, 0, len(ticketCounts))
	for id := range ticketCounts {
		ids = append(ids, id)
	}

	timetables, err := s.timetableRepo.FindByIds(c.Request().Context(), ids)
	if err != nil {
		return err
	}

	for _, timetable := range timetables {
		var count int32 = ticketCounts[timetable.ID]
		if timetable.Stock < count && isDecrease {
			return errors.New("not enough stock")
		}

		if isDecrease {
			timetable.Stock -= count
		} else {
			timetable.Stock += count
		}
		if err := s.timetableRepo.UpdateStock(c.Request().Context(), &timetable); err != nil {
			return err
		}
	}

	return nil
}

type TimetableService interface {
	FindById(c echo.Context, id uuid.UUID) (*entity.Timetable, error)
	FindByIds(c echo.Context, ids []uuid.UUID) ([]entity.Timetable, error)
	UpdateTicketStock(c echo.Context, ticketCounts map[uuid.UUID]int32, isDecrease bool) error
}

func NewTimetableService(timetableRepo repository.TimetableRepository) TimetableService {
	return &timetableService{timetableRepo: timetableRepo}
}
