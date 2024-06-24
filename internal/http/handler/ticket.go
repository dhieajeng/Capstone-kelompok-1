package handler

import (
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TicketHandler struct {
	ticketService      service.TicketService
	transactionService service.TransactionService
}

func (h *TicketHandler) UseTicket(c echo.Context) error {
	req := new(binder.UseTicketRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	ticket, err := h.ticketService.FindByNoTicket(c, req.NoTicket)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	if !ticket.IsValid {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, "ticket sudah pernah digunakan"))
	}

	transaction, err := h.transactionService.FindTransactionById(c, ticket.TransactionID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	if transaction.Status != "paid" {
		return c.JSON(http.StatusPaymentRequired, response.Error(http.StatusPaymentRequired, false, "transaksi belum dibayar"))
	}

	ticketDTO := entity.UsedTicket(ticket.ID)
	ticketResponse, err := h.ticketService.EditTicket(c, ticketDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "tiket telah digunakan", ticketResponse))
}

func NewTicketHandler(ticketService service.TicketService, transactionService service.TransactionService) TicketHandler {
	return TicketHandler{
		ticketService:      ticketService,
		transactionService: transactionService,
	}
}
