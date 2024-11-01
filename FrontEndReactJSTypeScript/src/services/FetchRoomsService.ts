import axios from 'axios';
import { ResponseModel } from './ResponseModel';

export type RoomDetails = {

}

const FetchRoomsService = async (): Promise<ResponseModel<RoomDetails[]>> => {
    const token = localStorage.token;
    try {
        let loginResponse = await axios.get<ResponseModel<RoomDetails[]>>('/v1/rooms', {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        console.log('FetchRoomsService', loginResponse.data);
        return loginResponse.data;
    } catch (error) {
        console.error('Error logging in', error || 'Unknown error')
        throw error
    }
}

export default FetchRoomsService;