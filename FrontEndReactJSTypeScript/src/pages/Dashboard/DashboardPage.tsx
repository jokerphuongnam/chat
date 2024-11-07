import 'bootstrap/dist/css/bootstrap.min.css';
import { Stack } from 'react-bootstrap';
import { useEffect, useState } from 'react';
import './DashboardPage.css';
import SearchResultBlock from './SearchResult/SearchResultBlock';
import RoomsBlock from './Rooms/RoomsBlock';
import SearchBlock from './Search/SearchBlock';
import HeaderBlock from './Header/HeaderBlock';
import MessagesBlock from './Messages/MessagesBlock';
import { UserInfo } from '../../services/FecthUserInfoService';
import RoomInfoBlock from './RoomInfo/RoomInfoBlock';
import useWebSocket from '../../hooks/useWebSocket';
import { RoomDetails } from '../../services/FetchRoomsService';

export type RoomInfo = {
    roomImage: string;
    roomName: string;
}

const DashboardPage: React.FC = () => {
    const [searchQuery, setSearchQuery] = useState('');
    const [isSearchResultVisible, setSearchResultVisible] = useState(false);
    const [userInfo, setUserInfo] = useState<UserInfo | null>(null);
    const [roomId, setRoomId] = useState<string>('');
    const [message, sendMessage] = useWebSocket('ws://localhost:3031/ws', localStorage.token);
    const [roomDetails, setRoomDetails] = useState<RoomDetails[]>([]);

    useEffect(() => {

    });

    const onRoomSelect = (roomId: string | null) => {
        if (roomId) {
            setRoomId(roomId);
        }
        setSearchResultVisible(false);
        setSearchQuery('');
    }

    useEffect(() => {
        if (roomDetails.length > 0) {
            let firstRooom = roomDetails[0];
            setUserInfo({
                id: firstRooom.id,
                name: firstRooom.name,
                avatarUrl: firstRooom.imageUrl
            });
        }
    }, [roomDetails]);

    return (
        <>
            <div className='dashboard-root mt-0 mb-0'>
                <Stack direction='horizontal' gap={3} className='mt-0 mb-0 display-flex dashboard-stack-container'>
                    <div className='p-3 col-3 dashboard-left display-flex dashboard-element'>
                        <Stack direction='vertical' className='display-flex' gap={2}>
                            <div className='dashboard-search'>
                                <SearchBlock searchQuery={searchQuery} setSearchQuery={setSearchQuery} onFocus={setSearchResultVisible} />
                            </div>
                            {isSearchResultVisible ? <SearchResultBlock searchQuery={searchQuery} onClicked={onRoomSelect} setUserInfo={setUserInfo} /> : < RoomsBlock roomDetails={roomDetails} setRoomDetails={setRoomDetails} selectRoom={onRoomSelect} />}
                        </Stack>
                    </div>

                    {
                        userInfo && (
                            <>
                                <div className='p-2 col-5 dashboard-center dashboard-element'>
                                    <div className='dashboard-header'>
                                        <HeaderBlock avatarUrl={userInfo?.avatarUrl ?? ''} name={userInfo?.name ?? ''} />
                                    </div>
                                    <MessagesBlock roomId={roomId} sendMessage={sendMessage} userId={userInfo?.id} />
                                </div>
                                <div className='p-2 col-3 dashboard-right dashboard-element'>
                                    <RoomInfoBlock />
                                </div>
                            </>
                        )
                    }
                </Stack>
            </div>
        </>
    );
}

export default DashboardPage;