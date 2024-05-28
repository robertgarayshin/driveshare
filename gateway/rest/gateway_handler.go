package rest

import (
	"booking-service/proto/booking"
	"car-service/proto/car"
	"context"
	"encoding/json"
	"net/http"
	"user-service/proto/user"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type GatewayHandler struct {
	userClient    user.UserServiceClient
	carClient     car.CarServiceClient
	bookingClient booking.BookingServiceClient
}

func NewGatewayHandler(userConn, carConn, bookingConn *grpc.ClientConn) *GatewayHandler {
	return &GatewayHandler{
		userClient:    user.NewUserServiceClient(userConn),
		carClient:     car.NewCarServiceClient(carConn),
		bookingClient: booking.NewBookingServiceClient(bookingConn),
	}
}

func (h *GatewayHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	userReq := &user.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	userResp, err := h.userClient.CreateUser(context.Background(), userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userResp)
}

func (h *GatewayHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userReq := &user.GetUserRequest{Id: id}
	userResp, err := h.userClient.GetUser(context.Background(), userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userResp)
}

// Implement similar handlers for Car and Booking
