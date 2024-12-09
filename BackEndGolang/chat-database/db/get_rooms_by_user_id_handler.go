package database

import (
	"chat-database/ent"
	"chat-database/ent/member"
	"chat-database/ent/message"
	"context"

	"github.com/google/uuid"
)

// RoomDetails represents the structure for the room information along with last message and members.
type RoomDetails struct {
	ID          uuid.UUID            `json:"id"`
	Name        *string              `json:"name"`
	ImageURL    *string              `json:"image_url"`
	LastMessage *LastMessageResponse `json:"last_message"`
	Members     []MemberResponse     `json:"members"`
}

// LastMessageResponse represents the structure for the last message in the room.
type LastMessageResponse struct {
	Content  string              `json:"content"`
	SendAt   uint64              `json:"send_at"`
	SenderID uuid.UUID           `json:"sender"`
	Type     message.TypeMessage `json:"type_message"`
}

// MemberResponse represents the structure for a room member.
type MemberResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AvatarURL *string   `json:"avatar_url,omitempty"`
}

func (db *Database) GetRoomsByUserHandler(userID uuid.UUID) ([]*RoomDetails, error) {
	// Get rooms where the user is a member
	members, err := db.Client.Member.Query().
		Where(member.UserID(userID)).
		WithRooms().
		All(context.Background())
	if err != nil {
		return nil, err
	}

	roomsDetails := make([]*RoomDetails, 0)

	for _, member := range members {
		r := member.Edges.Rooms

		// Fetch room info, handle not found
		roomInfo, err := r.QueryRoomInfo().First(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}

		// Get the last message in the room
		lastMessage, err := r.QueryMessages().
			Order(ent.Desc(message.FieldDateSend)).
			First(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}

		// Get all members in the room
		membersInRoom, err := r.QueryMembers().WithUsers().All(context.Background())
		if err != nil {
			return nil, err
		}

		// Construct room details
		roomDetails := &RoomDetails{
			ID:       member.RoomID,
			Name:     getRoomName(roomInfo, membersInRoom),
			ImageURL: getRoomImageURL(roomInfo, membersInRoom),
			LastMessage: &LastMessageResponse{
				Content:  getLastMessageContent(lastMessage),
				SendAt:   getLastMessageSendAt(lastMessage),
				SenderID: getLastMessageSenderID(lastMessage),
				Type:     getLastMessageType(lastMessage),
			},
			Members: getMembersResponse(membersInRoom),
		}

		roomsDetails = append(roomsDetails, roomDetails)
	}

	// // Sort the room details by the last message send time (ascending order)
	// sort.Slice(roomsDetails, func(i, j int) bool {
	// 	if roomsDetails[i].LastMessage != nil && roomsDetails[j].LastMessage != nil {
	// 		return roomsDetails[i].LastMessage.SendAt < roomsDetails[j].LastMessage.SendAt
	// 	}
	// 	return false
	// })

	return roomsDetails, nil
}

// getRoomName returns the room name based on room info and members.
func getRoomName(roomInfo *ent.RoomInfo, membersInRoom []*ent.Member) *string {
	if roomInfo != nil && roomInfo.Name != "" {
		if roomInfo.Name == "" {
			return nil
		}
		return &roomInfo.Name
	}
	if len(membersInRoom) > 0 {
		// If the nickname is available, return it; otherwise, return the user's name
		if membersInRoom[0].NickName != "" {
			return &membersInRoom[0].NickName
		}
		return &membersInRoom[0].Edges.Users.Name // Assuming User has a Name field
	}
	return nil
}

// getRoomImageURL returns the room image URL based on room info and members.
func getRoomImageURL(roomInfo *ent.RoomInfo, membersInRoom []*ent.Member) *string {
	if roomInfo != nil {
		if roomInfo.Name == "" {
			return nil
		}
		return &roomInfo.RoomImageURL
	}
	if len(membersInRoom) > 0 {
		return &membersInRoom[0].Edges.Users.AvatarURL // Assuming User has an AvatarURL field
	}
	return nil
}

// getLastMessageContent extracts the content of the last message.
func getLastMessageContent(lastMessage *ent.Message) string {
	if lastMessage != nil {
		return lastMessage.Content
	}
	return ""
}

// getLastMessageSendAt extracts the send timestamp of the last message.
func getLastMessageSendAt(lastMessage *ent.Message) uint64 {
	if lastMessage != nil {
		return lastMessage.DateSend
	}
	return 0
}

// getLastMessageSenderID extracts the sender's user ID from the last message.
func getLastMessageSenderID(lastMessage *ent.Message) uuid.UUID {
	if lastMessage != nil {
		return lastMessage.IDUserSend
	}
	return uuid.UUID{}
}

// getLastMessageType extracts the message type from the last message.
func getLastMessageType(lastMessage *ent.Message) message.TypeMessage {
	if lastMessage != nil {
		return lastMessage.TypeMessage
	}
	return ""
}

// getMembersResponse constructs the response structure for members in the room.
func getMembersResponse(membersInRoom []*ent.Member) []MemberResponse {
	members := make([]MemberResponse, 0, len(membersInRoom))
	for _, member := range membersInRoom {
		user := member.Edges.Users // Assuming each member has an associated user
		var avatarURL *string
		if user.AvatarURL == "" {
			avatarURL = nil
		} else {
			avatarURL = &user.AvatarURL
		}
		members = append(members, MemberResponse{
			ID:        member.UserID,
			Name:      user.Name,
			AvatarURL: avatarURL,
		})
	}
	return members
}
