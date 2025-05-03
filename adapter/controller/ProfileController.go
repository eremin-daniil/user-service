package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	requestModel "user_service/adapter/controller/request"
	responseModel "user_service/adapter/controller/response"
	"user_service/boundary/domain/usecase"
	"user_service/infrastructure/constant"
	loggerInterface "user_service/infrastructure/logger/interface"
)

type ProfileController struct {
	logger  loggerInterface.Logger
	useCase usecase.UserUseCaseInterface
}

func NewProfileController(logger loggerInterface.Logger, useCase usecase.UserUseCaseInterface) *ProfileController {
	return &ProfileController{logger: logger, useCase: useCase}
}

func (c *ProfileController) CreateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constant.UserIDCtxKey).(string)
	requestData := requestModel.NewEmptyCreateProfileRq()
	err := c.fillRequestData(w, r, requestData)
	if err != nil {
		return
	}
	id, err := c.useCase.CreateProfile(r.Context(), userID, requestData.ToDTO())
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.jsonResponse(w, id, http.StatusOK)
}

func (c *ProfileController) GetProfileCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constant.UserIDCtxKey).(string)
	dto, err := c.useCase.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := responseModel.ProfileCurrentUserRsFromDTO(dto)
	c.jsonResponse(w, response, http.StatusOK)
}

func (c *ProfileController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constant.UserIDCtxKey).(string)
	requestData := requestModel.NewEmptyUpdateProfileRq()
	err := c.fillRequestData(w, r, requestData)
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = c.useCase.UpdateProfile(r.Context(), userID, requestData.ToDTO())
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *ProfileController) fillRequestData(w http.ResponseWriter, r *http.Request, requestData any) error {
	arr, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	err = json.NewDecoder(bytes.NewReader(arr)).Decode(requestData)
	if err != nil {
		errorMassage := fmt.Sprintf("invalid request: %s", err)
		c.logger.Info(r.Context(), errorMassage)
		http.Error(w, errorMassage, http.StatusBadRequest)
		return err
	}
	return nil
}

func (c *ProfileController) jsonResponse(w http.ResponseWriter, result interface{}, statusCode int) {
	body, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
