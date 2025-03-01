export type TripPlan = {
    id: string
    title: string
    destination: string
    start_date: Date
    end_date: Date
}

export type RegisterTripPlanResponse = {
    id: string
}