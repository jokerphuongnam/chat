import axios from "axios"
import { ResponseModel } from "./ResponseModel"

const FetchRoomIDService = async (userId: string): Promise<ResponseModel<string | null>> => {
    try {
        const token = localStorage.token
        const response = await axios.get<ResponseModel<string | null>>(`/v1/room/get_room_id/${userId}`, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });
        return response.data;
    } catch (error) {
        throw error;
    }
}

export default FetchRoomIDService;