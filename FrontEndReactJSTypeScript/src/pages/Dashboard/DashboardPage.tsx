import 'bootstrap/dist/css/bootstrap.min.css';
import { Stack } from 'react-bootstrap';
import { useEffect, useState } from 'react';
import './DashboardPage.css';
import SearchResultBlock from './SearchResult/SearchResultBlock';
import RoomsBlock from './Rooms/RoomsBlock';
import SearchBlock from './Search/SearchBlock';
import HeaderBlock from './Header/HeaderBlock';
import MessagesBlock from './Messages/MessagesBlock';
import FecthUserInfoService from '../../services/FecthUserInfoService';
import { UserInfo } from '../../services/FecthUserInfoService';
import isSuccessfulResponse from '../../utils/HttpUtils';

const DashboardPage: React.FC = () => {
    const [searchQuery, setSearchQuery] = useState('');
    const [isSearchResultVisible, setSearchResultVisible] = useState(false);
    const [userInfo, setUserInfo] = useState<UserInfo | null>(null);
    const [roomId, setRoomId] = useState('');

    const fetchUserInfo = async () => {
        try {
            const response = await FecthUserInfoService();
            if (isSuccessfulResponse(response.code)) {
                setUserInfo(response.data);
            } else {

            }
        } catch (error) {

        }
    }

    const onRoomSelect = (roomId: string) => {
        setRoomId(roomId);
    }

    useEffect(() => {
        fetchUserInfo();
    }, []);

    return (
        <>
            <div className='dashboard-root mt-0 mb-0'>
                <Stack direction='horizontal' gap={2} className='mt-0 mb-0 display-flex dashboard-stack-container'>
                    <div className='p-2 col-3 dashboard-left display-flex'>
                        <Stack direction='vertical' className='display-flex' gap={2}>
                            <div className='dashboard-search'>
                                <SearchBlock searchQuery={searchQuery} setSearchQuery={setSearchQuery} onFocus={setSearchResultVisible} />
                            </div>
                            {isSearchResultVisible ? <SearchResultBlock searchQuery={searchQuery} onClicked={onRoomSelect}/> : < RoomsBlock />}
                        </Stack>
                    </div>
                    <div className='p-2 col-8 dashboard-right'>
                        <div className='dashboard-header'>
                            <HeaderBlock avatarUrl={userInfo?.avatarUrl ?? ''} name={userInfo?.name ?? ''} />
                        </div>
                        <MessagesBlock roomId={roomId} />
                    </div>
                </Stack>
            </div>
        </>
    );
}

export default DashboardPage;