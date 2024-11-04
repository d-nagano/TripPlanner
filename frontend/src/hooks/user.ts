import axios from 'axios';
import type { User } from '../types/types';

export const fetchUserSignUp = async (user: User) => {
    try {
        await axios.post('http://localhost/api/signup', user);
    } catch (error: any) {
        throw error;
    }
};