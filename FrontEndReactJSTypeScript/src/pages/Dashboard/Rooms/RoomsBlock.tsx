import 'bootstrap/dist/css/bootstrap.min.css';
import './RoomBlock.css'
import { useEffect, useState } from 'react';
import FetchRoomsService, { RoomDetails } from '../../../services/FetchRoomsService';
import isSuccessfulResponse from '../../../utils/HttpUtils';
import { Alert, ListGroup, Spinner } from 'react-bootstrap';
import LoadingComponent from '../../../components/Loading/LoadingComponent';

const RoomBlock: React.FC = () => {
    const [roomDetails, setRoomDetails] = useState<RoomDetails[]>([]);
    const [isLoading, setLoading] = useState(true);
    const [error, setError] = useState('');

    const fetchRooms = async () => {
        setLoading(true);
        try {
            const result = await FetchRoomsService();
            if (isSuccessfulResponse(result.code)) {
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
                            {roomDetails.map((item, index) => (
                                <div />
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