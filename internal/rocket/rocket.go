//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket go-grpc-services-course/internal/rocket Store

package rocket

import (
	"context"
	"log"
)

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we expect our db to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rckt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service responsible for
// updating the rocket inventory
type Service struct {
	Store Store
}

// New - returns an instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID Get Rocket by id retrives the rocket based on if from the the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - inserts a new rocket into the store
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	log.Println("inside1")

	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
