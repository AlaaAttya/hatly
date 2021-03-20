package createsession

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlaaAttya/hatly/session"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	validate          *validator.Validate
	sessionRepository session.Repository
}

func NewHandler(
	validate *validator.Validate,
	sessionRepository session.Repository,
) *Handler {
	return &Handler{
		validate:          validate,
		sessionRepository: sessionRepository,
	}
}

func (h *Handler) HandleRequest(req *http.Request) error {
	fmt.Println("handling request.....")

	// marshal request body
	request := &Request{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		return errors.New(fmt.Sprintf("request decoding object failed with error %s", err))
	}
	// validate request
	err = h.validate.Struct(request)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to validate request %s", err))
	}

	// check if we have already active session, then join it
	//dbSession, err := h.sessionRepository.FindActiveSessionByCode(request.Code)
	//if err != nil {
	//	return errors.New(fmt.Sprintf("failed to "))
	//}

	// create a new user

	// persist session in the db
	id, err := h.sessionRepository.Save(session.Session{
		Code: request.Code,
	})

	if err != nil {
		return errors.New(fmt.Sprintf("failed to save session to db %s", err))
	}

	fmt.Println("new session has been create with id: " + id)

	return nil
}
