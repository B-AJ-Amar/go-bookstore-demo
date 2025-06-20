package main

import (
	"context"
	"log"

	"github.com/B-AJ-Amar/go-bookstore-demo/proto/paymentpb"

	config "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	models "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	repositories "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/repositories"
)

type PaymentServer struct {
    paymentpb.UnimplementedPaymentServiceServer
}

// ? one grpc endpoint for test
func (s *PaymentServer) BuyBook(ctx context.Context, req *paymentpb.BuyBookRequest) (*paymentpb.BuyBookResponse, error) {
    log.Printf("Processing payment for Book ID: %d, Amount: %d", req.BookId, req.Quantity)
	bookRepo := repositories.NewBookRepository(config.GetDB())
	paymentRepo := repositories.NewPaymentRepository(config.GetDB())

	// Fetch the book from the database
	book, err := bookRepo.GetByID(uint(req.BookId))
	if err != nil {
		// log.Printf("Book not found: %v", err)
		return &paymentpb.BuyBookResponse{
			Success: false,
			Message: "Book not found",
		}, nil
	}

	// Check if enough quantity is available
	if book.Quantity < int(req.Quantity) {
		// log.Printf("Not enough quantity for Book ID: %s", req.BookId)
		return &paymentpb.BuyBookResponse{
			Success: false,
			Message: "Not enough quantity available",
		}, nil
	}

	// Decrease the quantity
	book.Quantity -= int(req.Quantity)
	if err := bookRepo.UpdateBookQuantity(book.ID, book.Quantity); err != nil {
		// log.Printf("Failed to update book quantity: %v", err)
		return &paymentpb.BuyBookResponse{
			Success: false,
			Message: "Failed to update book quantity",
		}, nil
	}

	// Add payment record to the payment database
	payment := models.Payment{
		BookID:   int(book.ID),
		Quantity: int(req.Quantity),
	}
	if err := paymentRepo.Create(&payment); err != nil {
		// log.Printf("Failed to add payment: %v", err)
		return &paymentpb.BuyBookResponse{
			Success: false,
			Message: "Failed to process payment",
		}, nil
	}


    // ? success
    return &paymentpb.BuyBookResponse{
        Success: true,
        Message: "Payment processed successfully",
    }, nil
}