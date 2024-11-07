import axios from "axios";
import { Message, MessageSnakeCase } from "../hooks/useWebSocket";
import { ResponseModel } from "./ResponseModel";

const SendMessageService = async (to: string, message: string, messageType: 'text' | 'image' | 'audio' | 'video' | 'location' | 'contact'): Promise<ResponseModel<Message>> => {
    try {
        console.log('send message', to, message, messageType);
        const token = localStorage.token
        const response = await axios.post<ResponseModel<MessageSnakeCase>>(`/v1/send-message`, {
            to,
            message,
            message_type: messageType,
        }, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });
        return {
            code: response.data.code,
            message: response.data.message,
            data: {
                messageType,
                content: message,
                senderId: response.data.data.sender_id,
                sentAt: response.data.data.sent_at,
            },
        };
    } catch (error) {
        throw error;
    }
}

export default SendMessageService;