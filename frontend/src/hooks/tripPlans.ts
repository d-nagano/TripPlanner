import axios from 'axios';
import type { TripPlan, RegisterTripPlanResponse } from '../types/tripPlans.ts';

export const fetchTripPlans = async (): Promise<TripPlan[]> => {
    try {
        const response = await axios.get<TripPlan[]>('http://localhost/api/trip-plan')
        return response.data
    } catch (error: any) {
        throw error;
    }
};

export const registerTripPlan = async (tripPlan: TripPlan): Promise<string> => {
    try {
        const response = await axios.post<RegisterTripPlanResponse>('http://localhost/api/trip-plan', tripPlan)
        return response.data.id
    } catch (error: any) {
        throw error;
    }
};

export const uploadImage = async (tripID: string) => {
    try {
        await axios.post(`http://localhost/api/trip-plan/${tripID}/image`);
    } catch (error: any) {
        throw error;
    }
};