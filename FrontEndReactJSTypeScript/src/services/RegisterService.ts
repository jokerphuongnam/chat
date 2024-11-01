import axios from 'axios'
import { ResponseModel } from './ResponseModel';

type RegisterResponse = {
    username: string;
    name: string;
    token: string;
}

const RegisterService = async (username: string, password: string, name: string): Promise<ResponseModel<RegisterResponse>> => {
    try {
        let registerResponse = await axios.post<ResponseModel<RegisterResponse>>('/v1/register', { username, password, name }, {
            headers: {
                'Content-Type': 'application/json',
            },
        })
        return registerResponse.data;
    } catch (error) {
        console.error('Error logging in', error || 'Unknown error')
        throw error
    }
}

export default RegisterService;