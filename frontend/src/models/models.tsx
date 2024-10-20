export interface User {
    id: number;
    username: string;
    email: string;
    bio?: string;          // Optional field
    profile_image?: string; // Optional field
}

export interface Post {
    id: number;
    user_id: number;
    image_url: string;
    caption: string;
    created_at: string
}

// Define a new interface that combines both User and Post
export interface FeedPost extends Post {
    username: string;       // from User
    email: string;          // from User
    bio?: string;           // Optional field from User
    profile_image: string; // Optional field from User
}

export interface Comment {
    id: number;
    post_id: number;
    user_id: number;
    content: string;
    created_at: string;
}