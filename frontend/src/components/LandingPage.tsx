// src/components/LandingPage.tsx
import React, { useState, useEffect } from 'react';
import { Box, Typography, Button, Card, CardMedia, CardContent } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import axiosInstance from '../axiosConfig'; // Using axios for fetching posts

interface Post {
    id: number;
    user_id: number;
    image_url: string;
    caption: string;
    created_at: string;
}

const LandingPage: React.FC = () => {
    const [posts, setPosts] = useState<Post[]>([]); // State to store posts
    const [error, setError] = useState<string | null>(null); // Error handling
    const navigate = useNavigate();

    useEffect(() => {
        // Fetch posts when the component mounts
        axiosInstance.get('http://localhost:8080/post/user/15')
            .then(response => {
                setPosts(response.data); // Set the posts data in state
                setError(null);
            })
            .catch(_ => {
                setError('Failed to fetch posts.'); // Handle errors
            });
    }, []); // Empty dependency array means it runs once when the component mounts

    const handleLogout = () => {
        localStorage.removeItem('token'); // Clear the token on logout
        navigate('/'); // Redirect back to login page
    };

    return (
        <Box
            sx={{
                height: '100vh',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                paddingTop: 2,
            }}
        >
            <Typography variant="h4" component="h1" gutterBottom>
                Instagram-like Feed
            </Typography>

            {/* Display posts */}
            <Box
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    width: '100%',
                    maxWidth: '600px',
                }}
            >
                {error && (
                    <Typography color="error" variant="body2">
                        {error}
                    </Typography>
                )}

                {posts.map(post => (
                    <Card key={post.id} sx={{ width: '100%', mb: 2 }}>
                        {/* Image */}
                        <CardMedia
                            component="img"
                            height="400"
                            image={post.image_url}
                            alt={post.caption}
                        />

                        {/* Caption and Created At */}
                        <CardContent>
                            <Typography variant="body1">
                                {post.caption}
                            </Typography>
                            <Typography variant="body2" color="textSecondary">
                                {new Date(post.created_at).toLocaleString()}
                            </Typography>
                        </CardContent>
                    </Card>
                ))}
            </Box>

            {/* Logout Button */}
            <Button variant="contained" color="primary" onClick={handleLogout} sx={{ mt: 4 }}>
                Logout
            </Button>
        </Box>
    );
};

export default LandingPage;