import axios from 'axios';
import { ResponseModel } from './ResponseModel';
import { snakeToCamel } from '../utils/snakeToCamel';

export type Member = {
    id: string;
    name: string;
}

export type LastMessage = {
    content: string;
    sendAt: number;
    sender: string;
    typeMessage: string;
}

export type RoomDetails = {
    id: string;
    name: string;
    imageUrl: string;
    lastMessage: LastMessage;
    members: Member[];
}

const FetchRoomsService = async (): Promise<ResponseModel<RoomDetails[]>> => {
    const token = localStorage.token;
    try {
        let response = await axios.get('/v1/rooms', {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });

        return {
            code: response.data.code,
            message: response.data.message,
            data: response.data.data.map((room: any) => snakeToCamel(room)),
        };
    } catch (error) {
        console.error('Error logging in', error || 'Unknown error')
        throw error
    }
}

export default FetchRoomsService;