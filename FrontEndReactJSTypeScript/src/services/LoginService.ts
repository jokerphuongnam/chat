import axios from 'axios'
import { ResponseModel } from './ResponseModel';

type LoginResponse = {
    name: string;
    avatarUrl: string;
    token: string;
}

type LoginResponseSnakeCase = {
    name: string;
    avatar_url: string;
    token: string;
}

const LoginService = async (username: string, password: string): Promise<ResponseModel<LoginResponse>> => {
    try {
        let loginResponse = await axios.post<ResponseModel<LoginResponseSnakeCase>>('/v1/login', { username, password }, {
            headers: {
                'Content-Type': 'application/json',
            },
        })
        return  {
            code: loginResponse.data.code,
            message: loginResponse.data.message,
            data: {
                name: loginResponse.data.data.name,
                avatarUrl: loginResponse.data.data.avatar_url,
                token: loginResponse.data.data.token,
            }
        };
    } catch (error) {
        console.error('Error logging in', error || 'Unknown error')
        throw error
    }
}

export default LoginService