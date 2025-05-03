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
	logger "user_service/infrastructure/logger/interface"
)

type UserController struct {
	logger  logger.Logger
	useCase usecase.UserUseCaseInterface
}

func NewUserController(logger logger.Logger, useCase usecase.UserUseCaseInterface) *UserController {
	return &UserController{logger: logger, useCase: useCase}
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	requestData := requestModel.NewEmptyRegisterUserRq()
	err := c.fillRequestData(w, r, requestData)
	if err != nil {
		return
	}
	id, err := c.useCase.RegisterUser(r.Context(), requestData.ToDTO())
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := responseModel.RegisterUserRsFromID(id)
	c.jsonResponse(w, response, http.StatusOK)
}

func (c *UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constant.UserIDCtxKey).(string)
	dto, err := c.useCase.GetUserByID(r.Context(), userID)
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := responseModel.CurrentUserRsFromDTO(dto)
	c.jsonResponse(w, response, http.StatusOK)
}

func (c *UserController) GetByLogin(w http.ResponseWriter, r *http.Request) {
	pathVariables := r.Context().Value(constant.PathVariablesCtxKey).(map[string]string)
	login := pathVariables["login"]
	dto, err := c.useCase.GetUserByLogin(r.Context(), login)
	if err != nil {
		c.logger.Error(r.Context(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := responseModel.OtherUserRsFromDTO(dto)
	c.jsonResponse(w, response, http.StatusOK)
}

func (c *UserController) fillRequestData(w http.ResponseWriter, r *http.Request, requestData any) error {
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

func (c *UserController) jsonResponse(w http.ResponseWriter, result interface{}, statusCode int) {
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
