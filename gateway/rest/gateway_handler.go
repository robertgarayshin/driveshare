package rest

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/robertgarayshin/driveshare/pkg/model/booking"
	"github.com/robertgarayshin/driveshare/pkg/model/car"
	"github.com/robertgarayshin/driveshare/pkg/model/user"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/dgrijalva/jwt-go"
)

type CarParameters struct {
	Brand          string  `json:"brand"`
	Model          string  `json:"model"`
	BodyType       string  `json:"bodyType"`
	Drive          string  `json:"drive"`
	EngineType     string  `json:"engineType"`
	EngineCapacity float32 `json:"engineCapacity"`
	ReleaseYear    int     `json:"releaseYear"`
	Seats          int     `json:"seats"`
	Mileage        int     `json:"mileage"`
	Transmission   string  `json:"transmission"`
	WheelPosition  string  `json:"wheelPosition"`
}

type Price struct {
	BasePerDay    float64 `json:"basePerDay"`
	MinRentPeriod int     `json:"minRentPeriod"`
}

type RentTerms struct {
	Deposit             float64 `json:"deposit"`
	KilPerDay           int     `json:"kilPerDay"`
	AdditionalPayment   float64 `json:"additionalPayment"`
	MinDriverExperience int     `json:"minDriverExperience"`
	MinAge              int     `json:"minAge"`
}

type Insurance struct {
	КАСКО string `json:"КАСКО"`
	ОСАГО string `json:"ОСАГО"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type CarModel struct {
	Id          int           `json:"id"`
	Category    string        `json:"category"`
	Owner       int           `json:"owner"`
	MainPhoto   string        `json:"mainPhono"`
	Photo       []string      `json:"photo"`
	Produced    string        `json:"produced"`
	Status      string        `json:"status"`
	Rating      float64       `json:"rating"`
	Description string        `json:"description"`
	Parameters  CarParameters `json:"parameters"`
	Price       Price         `json:"price"`
	RentTerms   RentTerms     `json:"rentTerms"`
	Address     Address       `json:"address"`
	Insurance   Insurance     `json:"insurance"`
}

type UserModel struct {
	Id                int    `json:"id" form:"id"`
	Email             string `json:"email" form:"email"`
	Name              string `json:"name" form:"name"`
	Password          string `json:"password" form:"password"`
	Surname           string `json:"surname" form:"surname"`
	UserRole          string `json:"user_role" form:"user_role"`
	ConfirmationToken string `json:"confirmation_token" form:"confirmation_token"`
	Phone             string `json:"phone" form:"phone"`
	Nationality       string `json:"nationality" form:"nationality"`
	DrivingExperience int    `json:"driving_experience" form:"driving_experience"`
	BirthDate         string `json:"birthday_date" form:"birthday_date"`
}

type Review struct {
	Id        string    `json:"id"`
	CarId     int       `json:"carId"`
	Rating    float32   `json:"rating"`
	Comment   string    `json:"comment"`
	User      UserModel `json:"user"`
	CreatedAt string    `json:"createdAt"`
}

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
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
	users := router.Group("/viewer")
	{
		users.GET(":id", h.GetUser)
		users.OPTIONS("", h.EditProfilePreflight)
		users.PUT("", h.EditProfile)
	}
	router.GET("user/:id", h.GetUser)
	router.GET("car/categories", h.GetCategories)
	cars := router.Group("car")
	{
		cars.GET("", h.GetCars)
		cars.GET(":id", h.GetCar)
	}
	router.OPTIONS("auth/sign-in", h.SignInPreflight)

	router.POST("auth/sign-in", h.SignIn)
	router.GET("review/:id", h.GetReview)

	return router
}

func (h *GatewayHandler) GetCars(ctx *gin.Context) {
	size := ctx.Query("size")
	fmt.Println(size)
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	ctx.JSON(http.StatusOK,
		[]CarModel{
			{
				Id:        1,
				Category:  "car",
				Owner:     1,
				MainPhoto: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
				Photo: []string{
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
				},
				Produced:    "2002",
				Status:      "booked",
				Rating:      1.0,
				Description: "test",
				Parameters: CarParameters{
					Brand:          "test",
					Model:          "test",
					BodyType:       "test",
					Drive:          "test",
					EngineType:     "test",
					EngineCapacity: 0,
					ReleaseYear:    0,
					Seats:          0,
					Mileage:        0,
					Transmission:   "test",
					WheelPosition:  "left",
				},
				Price: Price{
					BasePerDay:    1.0,
					MinRentPeriod: 0,
				},
				RentTerms: RentTerms{
					Deposit:             1.0,
					KilPerDay:           0,
					AdditionalPayment:   1.0,
					MinDriverExperience: 0,
					MinAge:              0,
				},
				Address: Address{
					City:   "test",
					Street: "test",
				},
				Insurance: Insurance{
					КАСКО: "test",
					ОСАГО: "test",
				},
			},
			{
				Id:        2,
				Category:  "car",
				Owner:     1,
				MainPhoto: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
				Photo: []string{
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
				},
				Produced:    "2002",
				Status:      "booked",
				Rating:      1.0,
				Description: "test",
				Parameters: CarParameters{
					Brand:          "test",
					Model:          "test",
					BodyType:       "test",
					Drive:          "test",
					EngineType:     "test",
					EngineCapacity: 0,
					ReleaseYear:    0,
					Seats:          0,
					Mileage:        0,
					Transmission:   "test",
					WheelPosition:  "left",
				},
				Price: Price{
					BasePerDay:    1.0,
					MinRentPeriod: 0,
				},
				RentTerms: RentTerms{
					Deposit:             1.0,
					KilPerDay:           0,
					AdditionalPayment:   1.0,
					MinDriverExperience: 0,
					MinAge:              0,
				},
				Address: Address{
					City:   "test",
					Street: "test",
				},
				Insurance: Insurance{
					КАСКО: "test",
					ОСАГО: "test",
				},
			},
			{
				Id:        3,
				Category:  "car",
				Owner:     1,
				MainPhoto: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
				Photo: []string{
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
				},
				Produced:    "2002",
				Status:      "booked",
				Rating:      1.0,
				Description: "test",
				Parameters: CarParameters{
					Brand:          "test",
					Model:          "test",
					BodyType:       "test",
					Drive:          "test",
					EngineType:     "test",
					EngineCapacity: 0,
					ReleaseYear:    0,
					Seats:          0,
					Mileage:        0,
					Transmission:   "test",
					WheelPosition:  "left",
				},
				Price: Price{
					BasePerDay:    1.0,
					MinRentPeriod: 0,
				},
				RentTerms: RentTerms{
					Deposit:             1.0,
					KilPerDay:           0,
					AdditionalPayment:   1.0,
					MinDriverExperience: 0,
					MinAge:              0,
				},
				Address: Address{
					City:   "test",
					Street: "test",
				},
				Insurance: Insurance{
					КАСКО: "test",
					ОСАГО: "test",
				},
			},
		})

}

func (h *GatewayHandler) GetCar(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	ctx.JSON(http.StatusOK, CarModel{
		Id:        1,
		Category:  "car",
		Owner:     1,
		MainPhoto: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
		Photo: []string{
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4Sqy_Jx2ZJVmI60I_3fvj7TcNsCGFWuRPcA&s",
		},
		Produced:    "2002",
		Status:      "booked",
		Rating:      1.8,
		Description: "test",
		Parameters: CarParameters{
			Brand:          "test",
			Model:          "test",
			BodyType:       "test",
			Drive:          "test",
			EngineType:     "test",
			EngineCapacity: 0,
			ReleaseYear:    0,
			Seats:          0,
			Mileage:        0,
			Transmission:   "test",
			WheelPosition:  "left",
		},
		Price: Price{
			BasePerDay:    1.0,
			MinRentPeriod: 0,
		},
		RentTerms: RentTerms{
			Deposit:             1.0,
			KilPerDay:           0,
			AdditionalPayment:   1.0,
			MinDriverExperience: 0,
			MinAge:              0,
		},
		Address: Address{
			City:   "test",
			Street: "test",
		},
		Insurance: Insurance{
			КАСКО: "test",
			ОСАГО: "test",
		},
	})
}

func (h *GatewayHandler) GetCategories(ctx *gin.Context) {
	type cat struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Photo    string `json:"photo"`
		MinPrice int    `json:"minPrice"`
		MaxPrice int    `json:"maxPrice"`
	}
	cats := []cat{
		{
			ID:       "1",
			Name:     "test",
			Photo:    "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Google_Ads_logo.svg/1200px-Google_Ads_logo.svg.png",
			MinPrice: 500,
			MaxPrice: 1000,
		},
		{
			ID:       "2",
			Name:     "test",
			MinPrice: 500,
			MaxPrice: 1000,
		},
		{
			ID:       "3",
			Name:     "test",
			Photo:    "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Google_Ads_logo.svg/1200px-Google_Ads_logo.svg.png",
			MinPrice: 500,
			MaxPrice: 1000,
		},
		{
			ID:       "4",
			Name:     "test",
			Photo:    "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Google_Ads_logo.svg/1200px-Google_Ads_logo.svg.png",
			MinPrice: 500,
			MaxPrice: 1000,
		},
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	ctx.JSON(http.StatusOK, cats)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *GatewayHandler) SignInPreflight(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "*")

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *GatewayHandler) SignIn(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "*")
	type signInInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input signInInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	token, _ := GenerateToken()
	type confirm struct {
		ConfirmToken string `json:"confirmationToken"`
		UserId       int    `json:"id"`
	}
	ctx.JSON(http.StatusOK, confirm{ConfirmToken: token, UserId: 1})
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		1,
	})

	return token.SignedString([]byte(signingKey))
}

func (h *GatewayHandler) GetUser(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	ctx.JSON(http.StatusOK, UserModel{
		Id:                1,
		Email:             "test@gmail.com",
		Name:              "test",
		Password:          "test",
		Surname:           "test",
		UserRole:          "test",
		ConfirmationToken: "test",
	})
}

func (h *GatewayHandler) EditProfilePreflight(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "*")

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *GatewayHandler) EditProfile(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	usr := &UserModel{}

	err := ctx.Bind(usr)
	if err != nil {
		fmt.Println(err)
	}

	spew.Dump(usr)
}

func (h *GatewayHandler) GetReview(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	ctx.JSON(http.StatusOK, []Review{
		{Id: "1",
			CarId:   1,
			Rating:  5,
			Comment: "tesrttttt",
			User: UserModel{

				Id:                1,
				Email:             "test@gmail.com",
				Name:              "test",
				Password:          "test",
				Surname:           "test",
				UserRole:          "test",
				ConfirmationToken: "test",
			},
			CreatedAt: "10 окт 2022",
		}})
}
