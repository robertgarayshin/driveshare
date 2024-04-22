package presenter

import (
	"context"
	"github.com/robertgarayshin/driveshare/proto/cars"
)

type CarServer struct {
	cars.CarsServer
}

func (s *CarServer) CreateCar(c context.Context, car *cars.Car) (*cars.CarResponse, error) {

}

func (s *CarServer) EditCar(c context.Context, car *cars.Car) (*cars.CarResponse, error) {

}

func (s *CarServer) DeleteCar(c context.Context, car *cars.Car) (*cars.CarResponse, error) {

}

func (s *CarServer) GetAllCars(c context.Context, req *cars.CarRequest) (*cars.CarResponse, error) {

}

func (s *CarServer) GetCarById(c context.Context, req *cars.Car) (*cars.CarResponse, error) {

}
