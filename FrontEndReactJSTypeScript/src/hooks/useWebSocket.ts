import { useCallback, useEffect, useRef, useState } from 'react';

// Request Type
export type SendMessageRequest = {
    to: string;            // ID of the recipient
    message: string;       // Message content
    messageType: 'text' | 'image' | 'audio' | 'video' | 'location' | 'contact'; // Enum for message type
}

// Response Type
export type Message = {
    messageType: 'text' | 'image' | 'audio' | 'video' | 'location' | 'contact'; // Enum for message type
    content: string;      // Message content
    senderId: string;     // UUID of the sender
    sentAt: number;       // Timestamp in Unix format
}

export type MessageSnakeCase = {
    message_type: 'text' | 'image' | 'audio' | 'video' | 'location' | 'contact'; // Enum for message type
    content: string;      // Message content
    sender_id: string;     // UUID of the sender
    sent_at: number;       // Timestamp in Unix format
}

const useWebSocket = (url: string, token: string): [Message | undefined, (message: SendMessageRequest) => void] => {
    const ws = useRef<WebSocket | null>(null);
    const [message, setMessage] = useState<Message>();

    useEffect(() => {
        ws.current = new WebSocket(url, token);

        ws.current.onopen = () => {
            console.log('WebSocket connection established');
        };

        ws.current.onclose = () => {
            console.log('WebSocket connection closed');
        };

        return () => {
            ws.current?.close();
        };
    }, [url, token]);

    useEffect(() => {
        if (!ws.current) return;

        ws.current.onmessage = (event) => {
            const data: MessageSnakeCase = JSON.parse(event.data);
            setMessage({
                messageType: data.message_type,
                content: data.content,
                senderId: data.sender_id,
                sentAt: data.sent_at,
            });
        }
    }, []);

    const sendMessage = useCallback((message: SendMessageRequest) => {
        ws.current?.send(JSON.stringify({
            to: message.to,
            message: message.message,
            message_type: message.messageType,
        }));
    }, []);

    return [message, sendMessage];
}

export default useWebSocket;