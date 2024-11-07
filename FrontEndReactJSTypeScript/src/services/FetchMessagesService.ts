import axios from "axios";
import { ResponseModel } from "./ResponseModel";

type GetMessagesResponseSnakeCase = {
    id: string;            // Message ID
    sender_id: string;     // ID of the sender
    time: string;          // Timestamp of when the message was sent
    content: string;       // Message content
    content_type: "text" | "image" | "audio" | "video" | "location" | "contact"; // Enum for message type
    nickname: string;      // Nickname of sender, fallback to name if blank or null
    avatar_url: string;    // URL of sender's avatar
    color: string; // Color of room
    is_current_user: boolean // is current user
};

export type GetMessagesResponse = {
    id: string;            // Message ID
    senderId: string;     // ID of the sender
    time: string;          // Timestamp of when the message was sent
    content: string;       // Message content
    contentType: "text" | "image" | "audio" | "video" | "location" | "contact"; // Enum for message type
    nickname: string;      // Nickname of sender, fallback to name if blank or null
    avatarUrl: string;    // URL of sender's avatar
    color: string;  // Color of room
    token: string; // Token
    isCurrentUser: boolean // is current user
}

const FetchMessagesService = async (roomID: string): Promise<ResponseModel<GetMessagesResponse[]>> => {
    const token = localStorage.token;
    try {
        const response = await axios.get<ResponseModel<GetMessagesResponseSnakeCase[]>>(`/v1/rooms/${roomID}/messages`, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });
        return {
            code: response.data.code,
            message: response.data.message,
            data: response.data.data.map((message: GetMessagesResponseSnakeCase) => ({
                id: message.id,
                senderId: message.sender_id,
                time: message.time,
                content: message.content,
                contentType: message.content_type,
                nickname: message.nickname,
                avatarUrl: message.avatar_url,
                color: message.color,
                token: token,
                isCurrentUser: message.is_current_user
            }))
        }
    } catch (error) {
        throw error;
    }
}

export default FetchMessagesService;