import axios from 'axios';
import type { User, LoginRequest } from '../types/users.ts';

export const fetchUserSignUp = async (user: User) => {
    try {
        await axios.post('http://localhost/api/signup', user);
    } catch (error: any) {
        throw error;
    }
};

export const fetchUserLogin = async (user: LoginRequest) => {
    try {
        await axios.post('http://localhost/api/login', user);
    } catch (error: any) {
        throw error;
    }
};