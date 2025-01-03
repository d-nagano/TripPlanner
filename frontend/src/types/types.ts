export type Prefecture = {
    id: number
    name: string
}

export type User = {
    name: string
    email: string
    password: string
}

export type LoginRequest = {
    email: string
    password: string
}