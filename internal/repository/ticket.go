package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

type ticketRepository struct {
	db *gorm.DB
}

func (r *ticketRepository) Creates(c context.Context, tickets *[]entity.Ticket) (*[]entity.Ticket, error) {
	if err := r.db.WithContext(c).Create(tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *ticketRepository) FindByNoTicket(c context.Context, noTicket string) (*entity.Ticket, error) {
	ticket := new(entity.Ticket)
	if err := r.db.WithContext(c).Where("no_ticket = ?", noTicket).Take(ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *ticketRepository) Edit(c context.Context, ticket *entity.Ticket) (*entity.Ticket, error) {
	query := r.db.WithContext(c).Model(&ticket).Where("id = ?", ticket.ID)
	query = query.Update("is_valid", ticket.IsValid)

	if err := query.Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

type TicketRepository interface {
	Creates(c context.Context, tickets *[]entity.Ticket) (*[]entity.Ticket, error)
	FindByNoTicket(c context.Context, noTicket string) (*entity.Ticket, error)
	Edit(c context.Context, ticket *entity.Ticket) (*entity.Ticket, error)
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}
