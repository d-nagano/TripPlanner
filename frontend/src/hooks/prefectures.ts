import axios from 'axios'
import type { Prefecture } from '../types/types';

export const fetchPrefectures = async (): Promise<Prefecture[]> => {
    try {
        const response = await axios.get<Prefecture[]>('http://localhost/api/prefectures');
        return response.data;
    } catch (error) {
        console.error('error fetching prefectures:', error);
        throw error;
    }
};