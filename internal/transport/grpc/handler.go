package grpc

import (
	"context"
	"go-grpc-services-course/internal/rocket"
	"log"
	"net"

	rkt "github.com/TutorialEdge/tutorial-protos/rocket/v1"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Println("Could not listen to grpc")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

// Get Rocket - retrieves rocket by id and returns response
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("Get Rocket gRPC hit")
	rocket, err := h.RocketService.GetRocketByID(ctx, req.Id)
	if err != nil {
		log.Println("Failed to retrieve rocket by id")
		return &rkt.GetRocketResponse{}, err
	}
	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

// Add Rocket - Adds rocket to db
func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC hit")
	rocket, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})

	if err != nil {
		log.Println("Failed to insert rocket")
		return &rkt.AddRocketResponse{}, err
	}

	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		}}, nil
}

// Delete Rocket - deletes rocket
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Println("delete rocket gRPC endpoint hit")

	log.Println(req.Rocket)

	err := h.RocketService.DeleteRocket(ctx, req.Rocket.Id)

	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "success",
	}, nil
}
