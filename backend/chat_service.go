package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"backend/proto"
)

// Server represents the gRPC server
type ChatServiceServer struct {
    messaging.UnimplementedChatServiceServer
}


// SendMessage saves the message to the database and broadcasts it to other connected clients.
func (s *ChatServiceServer) SendMessage(ctx context.Context, message *messaging.ChatMessage) (*messaging.ConnectionAck, error) {
	// Save message to the database
	// Note: You'll need to implement the actual database logic.
	err := saveMessageToDatabase(message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to save message to database: %v", err)
	}

	// Broadcast message to other connected clients
	// Note: You'll need to implement the actual broadcasting logic.
	err = broadcastMessageToClients(message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to broadcast message: %v", err)
	}

	return &messaging.ConnectionAck{
		Success: true,
		Message: "Message sent successfully",
	}, nil
}

// ReceiveMessages streams chat messages in real-time.
func (s *ChatServiceServer) ReceiveMessages(user *messaging.User, stream messaging.ChatService_ReceiveMessagesServer) error {
	// Fetch messages from the database
	// Note: You'll need to implement the actual database fetching logic.
	messages, err := fetchMessagesFromDatabase(user)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to fetch messages: %v", err)
	}

	// Stream messages to the client
	for _, message := range messages {
		if err := stream.Send(message); err != nil {
			return err
		}
	}

	return nil
}

// saveMessageToDatabase saves the chat message to the database.
// Note: This is a placeholder function. You'll need to implement the actual database logic.
func saveMessageToDatabase(message *messaging.ChatMessage) error {
	// TODO: Implement the logic to save the message to the PostgreSQL database.
	log.Println("Message saved to database:", message.Text)
	return nil
}

// broadcastMessageToClients broadcasts the chat message to other connected clients.
// Note: This is a placeholder function. You'll need to implement the actual broadcasting logic.
func broadcastMessageToClients(message *messaging.ChatMessage) error {
	// TODO: Implement the logic to broadcast the message to other connected clients.
	log.Println("Message broadcasted:", message.Text)
	return nil
}

// fetchMessagesFromDatabase fetches chat messages for a user from the database.
// Note: This is a placeholder function. You'll need to implement the actual database fetching logic.
func fetchMessagesFromDatabase(user *messaging.User) ([]*messaging.ChatMessage, error) {
	// TODO: Implement the logic to fetch messages for the user from the PostgreSQL database.
	return []*messaging.ChatMessage{}, nil
}
