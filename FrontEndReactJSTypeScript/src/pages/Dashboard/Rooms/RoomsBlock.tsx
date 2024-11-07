import 'bootstrap/dist/css/bootstrap.min.css';
import { useEffect, useState } from 'react';
import { Alert, ListGroup, Spinner } from 'react-bootstrap';
import './RoomBlock.css'
import FetchRoomsService, { RoomDetails } from '../../../services/FetchRoomsService';
import isSuccessfulResponse from '../../../utils/HttpUtils';

interface RoomBlockProps {
    roomDetails: RoomDetails[];
    setRoomDetails: (roomDetails: RoomDetails[]) => void;
    selectRoom: (roomId: string | null) => void;
}

const RoomBlock: React.FC<RoomBlockProps> = ({ roomDetails, setRoomDetails, selectRoom }) => {
    const [isLoading, setLoading] = useState(true);
    const [error, setError] = useState('');

    const fetchRooms = async () => {
        setLoading(true);
        try {
            const result = await FetchRoomsService();
            if (isSuccessfulResponse(result.code)) {
                setError('');
                setRoomDetails(result.data);
            } else {
                setError(result.message);
            }
        } catch (error) {
            console.error('Error fetching rooms', error);
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        fetchRooms();
    }, []);

    return (
        <>
            <div className='room-container h-100'>
                {
                    isLoading ? (
                        <Alert variant="error" className='non-scroll-container'>
                            <Spinner className='loading-spinner' animation='border' variant='dark' />
                        </Alert>
                    ) : error ? (
                        <Alert variant="error" className='non-scroll-container'>
                            {error}
                        </Alert>
                    ) : roomDetails && roomDetails.length > 0 ? (
                        <ListGroup>
                            {roomDetails.map((room, index) => (
                                <div key={room.name} className="room-item" onClick={() => {
                                    selectRoom(room.id);
                                    console.log(room.id);
                                }}>
                                    <img
                                        src={room.imageUrl || '/default-avatar.png'}
                                        alt={room.name}
                                        className="room-image"
                                        width={32}
                                        height={32}
                                    />
                                    <div className="room-info">
                                        <div className="room-name">{room.name}</div>
                                        <div className="room-last-message">
                                            {room.lastMessage.content}
                                        </div>
                                    </div>
                                    <div className="room-time">
                                        {new Date(room.lastMessage.sendAt).toLocaleTimeString()}
                                    </div>
                                </div>
                            ))}
                        </ListGroup>
                    ) : (
                        <Alert variant="info" className='non-scroll-container'>
                            No items to display.
                        </Alert>
                    )
                }
            </div>
        </>
    );
}

export default RoomBlock;