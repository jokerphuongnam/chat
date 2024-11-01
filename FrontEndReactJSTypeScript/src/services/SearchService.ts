import axios from 'axios'
import { ResponseModel } from './ResponseModel'

type SearchResponseSnakeCase = {
    id: string;
    name: string;
    avatar_url: string;
}

export type SearchResponse = {
    id: string;
    name: string;
    avatarUrl: string;
}

const SearchService = async (name: string | undefined): Promise<ResponseModel<SearchResponse[]>> => {
    try {
        name = name?.trim();
        const response = await axios.get<ResponseModel<SearchResponseSnakeCase[]>>(`/v1/search?name=${name}`, {
            headers: {
                'Content-Type': 'application/json',
            },
        });
        console.log('SearchService', response.data);
        return {
            code: response.data.code,
            message: response.data.message,
            data: response.data.data.map((item: SearchResponseSnakeCase) => ({
                id: item.id,
                name: item.name,
                avatarUrl: item.avatar_url,
            })),
        }
    } catch (error) {
        throw error;
    }
}

export default SearchService;