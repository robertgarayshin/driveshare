package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/robertgarayshin/driveshare/pkg/model/booking"
	"github.com/robertgarayshin/driveshare/pkg/model/car"
	"github.com/robertgarayshin/driveshare/pkg/model/user"
	"net/http"

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

func (h *GatewayHandler) RegisterRoutes() *gin.Engine {
	router := gin.New()
	users := router.Group("/users")
	{
		users.GET(":id", h.GetUser)
	}

	return router
}

func (h *GatewayHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userReq := &user.GetUserRequest{Id: id}
	userResp, err := h.userClient.GetUser(ctx, userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}
