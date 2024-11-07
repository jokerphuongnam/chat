import axios from 'axios';
import { ResponseModel } from './ResponseModel';

export interface UserInfo {
    id: string;
    name: string;
    avatarUrl: string;
}

type GetUserInfoResponse = {
    id: string;
    name: string;
    avatar_url: string;
}

const FecthUserInfoService = async (): Promise<ResponseModel<UserInfo | null>> => {
    const token = localStorage.token;
    try {
        let loginResponse = await axios.get<ResponseModel<GetUserInfoResponse | null>>(`/v1/user-info`, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });
        console.log('FetchRoomsService', loginResponse.data);
        return {
            code: loginResponse.data.code,
            message: loginResponse.data.message,
            data: loginResponse.data.data ? {
                id: loginResponse.data.data.id,
                name: loginResponse.data.data.name,
                avatarUrl: loginResponse.data.data.avatar_url
            } : null,
        };
    } catch (error) {
        console.error('Error logging in', error || 'Unknown error')
        throw error
    }
}

export default FecthUserInfoService;