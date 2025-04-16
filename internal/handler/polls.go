package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"quick-poll/internal/models"
	"quick-poll/pkg/errors"
	"quick-poll/pkg/responses"
)

func (h *Handler) CreatePoll(c echo.Context) error {
	var req models.PollRequest
	if err := c.Bind(&req); err != nil {
		h.s.Logger.Errorf("failed to parse request body: %v", err)
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse("failed parsing body"))
	}

	poll, err := h.s.CreatePoll(c.Request().Context(), req.Question, req.Options)
	if err != nil {
		h.s.Logger.Errorf("failed to create poll (question=%q, options=%v): %v", req.Question, req.Options, err)
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(fmt.Sprintf("failed creating poll: %v", err)))
	}

	return c.JSON(http.StatusCreated, responses.NewSuccessResponse(poll))
}

func (h *Handler) Vote(c echo.Context) error {
	pollID := c.Param("id")
	option := c.QueryParam("option")

	h.s.Logger.Infof("received vote request (pollID=%s, option=%s)", pollID, option)
	if err := h.s.Vote(c.Request().Context(), pollID, option); err != nil {
		switch err {
		case errors.ErrPollNotFound:
			h.s.Logger.Errorf("poll not found (pollID=%s)", pollID)
			return c.JSON(http.StatusBadRequest, responses.NewErrorResponse("poll not found"))
		case errors.ErrInvalidOption:
			h.s.Logger.Errorf("invalid option (pollID=%s, option=%s)", pollID, option)
			return c.JSON(http.StatusBadRequest, responses.NewErrorResponse("invalid option"))
		default:
			h.s.Logger.Errorf("failed to process vote (pollID=%s, option=%s): %v", pollID, option, err)
			return c.JSON(http.StatusBadRequest, responses.NewErrorResponse("failed voting"))
		}
	}

	h.s.Logger.Infof("vote sent successfully (pollID=%s, option=%s)", pollID, option)

	return c.JSON(http.StatusOK, responses.NewSuccessResponse("the vote has been sent"))
}

func (h *Handler) GetResults(c echo.Context) error {
	pollID := c.Param("id")
	h.s.Logger.Infof("fetching results for poll (pollID=%s)", pollID)

	poll, err := h.s.GetResults(c.Request().Context(), pollID)
	if err != nil {
		h.s.Logger.Errorf("poll not found (pollID=%s)", pollID)
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse("poll not found"))
	}
	h.s.Logger.Infof("results fetched successfully (pollID=%s)", pollID)

	return c.JSON(http.StatusOK, responses.NewSuccessResponse(poll))
}
