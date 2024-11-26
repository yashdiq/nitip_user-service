package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"user-service/pkg/model"
	"user-service/pkg/pb"
	"user-service/pkg/repository/interfaces"

	"github.com/google/uuid"
)

type UserHandler struct {
	UserRepository interfaces.UserRepository
	pb.UserServiceServer
}

func NewUserHandler(userRepository interfaces.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

func (h *UserHandler) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	log.Printf("user: %+v\n", req)

	instance := model.User{
		Id: uuid.NewString(),
		Username: req.Username,
		Email:    req.Email,
		Phone: req.Phone,
	}
	
	_, err := h.UserRepository.Save(instance)
	
	if err != nil {
		fmt.Println("error: ", err)

		return &pb.SaveResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error",
		}, err
	}
	
	return &pb.SaveResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

