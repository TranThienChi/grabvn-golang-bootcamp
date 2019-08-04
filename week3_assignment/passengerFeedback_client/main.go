package main

import (
	"context"
	"log"
	"time"

	pb "github.com/TranThienChi/grabvn-golang-bootcamp/week3_assignment/passengerFeedback"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func addPassengerFeedback(ctx context.Context, c pb.PassengerFeedbackServiceClient, f *pb.PassengerFeedback) error {
	_, err := c.AddPassengerFeedback(ctx, &pb.AddPassengerFeedbackRequest{PassengerFeedback: f})
	return err
}

func getByPassengerId(ctx context.Context, c pb.PassengerFeedbackServiceClient, id int32) (*pb.GetByPassengerIdResponse, error) {
	val, err := c.GetByPassengerId(ctx, &pb.GetByPassengerIdRequest{PassengerID: id})
	return val, err
}

func getByBookingCode(ctx context.Context, c pb.PassengerFeedbackServiceClient, code string) (*pb.GetByBookingCodeResponse, error) {
	val, err := c.GetByBookingCode(ctx, &pb.GetByBookingCodeRequest{BookingCode: code})
	return val, err
}

func deleteByPassengerId(ctx context.Context, c pb.PassengerFeedbackServiceClient, id int32) (*pb.DeleteByPassengerIdResponse, error) {
	val, err := c.DeleteByPassengerId(ctx, &pb.DeleteByPassengerIdRequest{PassengerID: id})
	return val, err
}

func test(ctx context.Context, c pb.PassengerFeedbackServiceClient) {
	// First passenger feedback
	feedback1 := pb.PassengerFeedback{
		BookingCode: "#1",
		PassengerID: 1,
		Feedback:    "Good!",
	}

	// Second passenger feedback
	feedback2 := pb.PassengerFeedback{
		BookingCode: "#2",
		PassengerID: 1,
		Feedback:    "Excellent!",
	}

	// Third passenger feedback
	feedback3 := pb.PassengerFeedback{
		BookingCode: "#3",
		PassengerID: 2,
		Feedback:    "Not good!",
	}

	// // Fourth passenger feedback
	// feedback4 := pb.PassengerFeedback{
	// 	BookingCode: "#3",
	// 	PassengerID: 3,
	// 	Feedback:    "Bad!",
	// }

	// Add first passenger feedback
	if err := addPassengerFeedback(ctx, c, &feedback1); err != nil {
		panic(err)
	}
	log.Printf("Already added feedback: %v\n", &feedback1)

	// Add second passenger feedback
	if err := addPassengerFeedback(ctx, c, &feedback2); err != nil {
		panic(err)
	}
	log.Printf("Already added feedback: %v\n", &feedback2)

	// Add third passenger feedback
	if err := addPassengerFeedback(ctx, c, &feedback3); err != nil {
		panic(err)
	}
	log.Printf("Already added feedback: %v\n", &feedback3)

	// // Add fourth passenger feedback
	// if err := addPassengerFeedback(ctx, c, &feedback4); err != nil {
	// 	panic(err)
	// }
	// log.Printf("Already added feedback: %v\n", &feedback4)

	// Get feedback from first passenger id
	if val, err := getByPassengerId(ctx, c, feedback1.PassengerID); err != nil {
		panic(err)
	} else {
		log.Printf("Passenger ID %v review %v\n", feedback1.PassengerID, val)
	}

	// Get feedback from second booking code
	if val, err := getByBookingCode(ctx, c, feedback2.BookingCode); err != nil {
		panic(err)
	} else {
		log.Printf("Passenger booking code %v review %v\n", feedback2.BookingCode, val)
	}

	// Delete feedback by passenger id
	if val, err := deleteByPassengerId(ctx, c, feedback1.PassengerID); err != nil {
		panic(err)
	} else {
		log.Printf("Deleted total %v feedbacks related with passenger id %v\n", val, feedback1.PassengerID)
	}
}

func main() {
	// Create a client connection
	connect, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()

	c := pb.NewPassengerFeedbackServiceClient(connect)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test case
	test(ctx, c)
}
