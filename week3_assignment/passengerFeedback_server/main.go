package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/TranThienChi/grabvn-golang-bootcamp/week3_assignment/passengerFeedback"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

// HandlerInterface abstraction passenger feedback service
type HandlerInterface interface {
	addPassengerFeedback(feedback *pb.PassengerFeedback) error
	getByPassengerId(id int32) ([]*pb.PassengerFeedback, error)
	getByBookingCode(code string) ([]*pb.PassengerFeedback, error)
	deleteByPassengerId(id int32)
}

// Handler handle passenger feedback by id and booking code
type Handler struct {
	mapBookingFeedback     map[string]*pb.PassengerFeedback
	mapPassengerIDFeedback map[int32][]*pb.PassengerFeedback
}

func (h *Handler) addPassengerFeedback(feedback *pb.PassengerFeedback) (err error) {
	if _, ok := h.mapBookingFeedback[feedback.BookingCode]; ok {
		err = errors.New("Booking code feedback is existed")
	} else {
		h.mapBookingFeedback[feedback.BookingCode] = feedback
		h.mapPassengerIDFeedback[feedback.PassengerID] = append(h.mapPassengerIDFeedback[feedback.PassengerID], feedback)
	}
	return
}

func (h *Handler) getByPassengerId(id int32) ([]*pb.PassengerFeedback, error) {
	if val, ok := h.mapPassengerIDFeedback[id]; ok {
		return val, nil
	}
	return nil, errors.New("Passenger id is not existed")
}

func (h *Handler) getByBookingCode(code string) (*pb.PassengerFeedback, error) {
	if val, ok := h.mapBookingFeedback[code]; ok {
		return val, nil
	}
	return nil, errors.New("Booking code is not existed")
}

func (h *Handler) deleteByPassengerId(id int32) (totalDeletedFeedback int32) {
	if val, ok := h.mapPassengerIDFeedback[id]; ok {
		totalDeletedFeedback = int32(len(val))
		delete(h.mapPassengerIDFeedback, id)
	}
	return
}

var handler = Handler{
	make(map[string]*pb.PassengerFeedback),
	make(map[int32][]*pb.PassengerFeedback),
}

func (s *server) AddPassengerFeedback(ctx context.Context, c *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {
	log.Printf("AddPassengerFeedback received: %v\n", c.PassengerFeedback)
	err := handler.addPassengerFeedback(c.PassengerFeedback)
	val := &pb.AddPassengerFeedbackResponse{PassengerFeedback: c.PassengerFeedback}
	return val, err
}

func (s *server) GetByPassengerId(ctx context.Context, c *pb.GetByPassengerIdRequest) (*pb.GetByPassengerIdResponse, error) {
	log.Printf("GetByPassengerId received: %v\n", c.PassengerID)
	val, err := handler.getByPassengerId(c.PassengerID)
	return &pb.GetByPassengerIdResponse{PassengerFeedbacks: val}, err
}

func (s *server) GetByBookingCode(ctx context.Context, c *pb.GetByBookingCodeRequest) (*pb.GetByBookingCodeResponse, error) {
	log.Printf("GetByBookingCode received: %v\n", c.BookingCode)
	val, err := handler.getByBookingCode(c.BookingCode)
	return &pb.GetByBookingCodeResponse{PassengerFeedback: val}, err
}

func (s *server) DeleteByPassengerId(ctx context.Context, c *pb.DeleteByPassengerIdRequest) (*pb.DeleteByPassengerIdResponse, error) {
	log.Printf("DeleteByPassengerId deleted: %v\n", c.PassengerID)
	deleted := handler.deleteByPassengerId(c.PassengerID)
	return &pb.DeleteByPassengerIdResponse{TotalDeletedFeedback: deleted}, nil
}

func main() {
	// Listen announces on the local network address.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPassengerFeedbackServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
