import axios from 'axios';
import type { TripPlan } from '../types/tripPlans.ts';

export const fetchTripPlans = async (): Promise<TripPlan[]> => {
    try {
        const response = await axios.get<TripPlan[]>('http://localhost/api/trip-plan')
        return response.data
    } catch (error: any) {
        throw error;
    }
};