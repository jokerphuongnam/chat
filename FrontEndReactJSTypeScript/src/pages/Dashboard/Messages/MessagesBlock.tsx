import 'bootstrap/dist/css/bootstrap.min.css';
import { useEffect, useState } from 'react';
import { Container, InputGroup, FormControl, Button } from 'react-bootstrap';
import FetchMessagesService, { GetMessagesResponse } from '../../../services/FetchMessagesService';
import isSuccessfulResponse from '../../../utils/HttpUtils';
import { SendMessageRequest } from '../../../hooks/useWebSocket';
import SendMessageService from '../../../services/SendMessageService';

interface MessagesBlockProps {
    roomId: string;
    userId?: string;
    sendMessage: (message: SendMessageRequest) => void;
}

const MessagesBlock: React.FC<MessagesBlockProps> = ({ roomId, userId, sendMessage }) => {
    const [messages, setMessages] = useState<GetMessagesResponse[]>([]);
    const [isLoading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const [inputMessage, setInputMessage] = useState('');

    const getMessages = async () => {
        if (!roomId) return;
        try {
            const response = await FetchMessagesService(roomId);
            if (isSuccessfulResponse(response.code)) {
                setMessages(response.data);
            } else {
                setError(response.message);
            }
        } catch (error) {
            console.error('Error fetching messages', error);
        }
    }

    const sendMessageService = async (message: string) => {
        if (!userId) return;
        try {
            const response = await SendMessageService(userId, message, 'text');
            if (isSuccessfulResponse(response.code)) {
                window.location.reload();
            }
        } catch (error) {

        }
    }

    const handleSendMessage = () => {
        const newMessage = inputMessage.trim();
        if (newMessage !== '') {
            if (roomId) {
                sendMessage({
                    to: roomId,
                    message: newMessage,
                    messageType: 'text',
                });
            } else {
                sendMessageService(newMessage);
            }
        }
    };

    useEffect(() => {
        getMessages();
    });

    return (
        <>
            <Container fluid className='d-flex flex-column'>
                <div className='flex-grow-1 overflow-auto messages-list'>
                    {messages.map((msg) => (
                        <MessageBubble
                            key={msg.id}
                            id={msg.id}
                            isCurrentUser={msg.isCurrentUser}
                            content={msg.content}
                            time={msg.time}
                            nickname={msg.nickname}
                            avatarUrl={msg.avatarUrl}
                            color={msg.color}
                            senderId={msg.senderId}
                            contentType={msg.contentType}
                        />
                    ))}
                </div>

                <InputGroup>
                    <FormControl
                        placeholder='Type a message...'
                        value={inputMessage}
                        onChange={(e) => setInputMessage(e.target.value)}
                        onKeyPress={(e) => {
                            if (e.key === 'Enter') handleSendMessage();
                        }}
                    />
                    <Button variant='primary' onClick={handleSendMessage}>Send</Button>
                </InputGroup>
            </Container>
        </>
    );
}

export default MessagesBlock;

interface MessageBubbleProps {
    isCurrentUser: boolean;
    id: string;
    senderId: string;
    time: string;
    content: string;
    contentType: 'text' | 'image' | 'audio' | 'video' | 'location' | 'contact';
    nickname: string;
    avatarUrl: string;
    color: string;
}

const MessageBubble: React.FC<MessageBubbleProps> = ({
    isCurrentUser,
    id,
    senderId,
    time,
    content,
    contentType,
    nickname,
    avatarUrl,
    color,
}) => {
    return (
        <div className={`d-flex ${isCurrentUser ? 'justify-content-end' : 'justify-content-start'} mb-2`}>
            {!isCurrentUser && (
                <div className='d-flex align-items-center me-2'>
                    {avatarUrl && (
                        <img
                            src={avatarUrl}
                            alt='avatar'
                            className='rounded-circle'
                        />
                    )}
                    {nickname && <span className='ms-2 text-muted small'>{nickname}</span>}
                </div>
            )}
            <div
                className={`p-2 rounded-pill ${isCurrentUser ? 'text-white' : 'text-dark'}`}
                style={{
                    backgroundColor: isCurrentUser ? color : '#f0f0f0',
                    maxWidth: '60%',
                }}
            >
                <span>{content}</span>
            </div>
            <div className='text-muted small ms-2 align-self-end'>
                {time}
            </div>
        </div>
    );
};