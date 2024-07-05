package handler

import pb "root/genproto"

type Handler struct {
	User pb.UserServiceClient
}

func NewHandler(us pb.UserServiceClient) *Handler {
	return &Handler{us}
}
