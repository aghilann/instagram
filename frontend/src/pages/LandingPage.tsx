import React, { useEffect, useState } from 'react';
import { Box, Typography } from '@mui/material';
import { Grid2} from "@mui/material";
import axiosInstance from '../axiosConfig';
import Header from '../components/Header.tsx';
import Post from '../components/PostCard.tsx';
import CommentsSection from '../components/CommentsSection.tsx';
import LogoutButton from '../components/LogoutButton.tsx';
import {Comment, FeedPost} from "../models/models.tsx";

const LandingPage: React.FC = () => {
    const [posts, setPosts] = useState<FeedPost[]>([]);
    const [comments, setComments] = useState<{ [key: number]: Comment[] }>({});
    const [commentVisible, setCommentVisible] = useState<{ [key: number]: boolean }>({});
    const [error, setError] = useState<string | null>(null);

    const handleNewComment = (postId: number, newComment: Comment) => {
        setComments(prevComments => ({
            ...prevComments,
            [postId]: [...(prevComments[postId] || []), newComment],
        }));
    };

    const handleDeleteComment = (postId: number, commentId: number) => {
        setComments(prevComments => ({
            ...prevComments,
            [postId]: prevComments[postId].filter(comment => comment.id !== commentId),
        }));
    };

    useEffect(() => {
        const userId = localStorage.getItem('userId');
        axiosInstance.get(`/post/feed/${userId}`)
            .then(response => {
                setPosts(response.data);
                setError(null);
            })
            .catch(_ => {
                setError('Failed to fetch posts.');
            });
    }, []);

    const fetchComments = (postId: number) => {
        if (comments[postId]) {
            setCommentVisible(prev => ({ ...prev, [postId]: !prev[postId] }));
            return;
        }

        axiosInstance.get(`/comment/post/${postId}`)
            .then(response => {
                setComments(prevComments => ({ ...prevComments, [postId]: response.data }));
                setCommentVisible(prev => ({ ...prev, [postId]: true }));
            })
            .catch(() => {
                console.error(`Failed to fetch comments for post ${postId}`);
            });
    };

    return (
        <Box
            sx={{
                height: '100vh',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                paddingTop: 2,
                overflowY: 'auto',
                bgcolor: 'background.default',
                color: 'text.primary',
            }}
        >
            <Header />

            <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', width: '100%', maxWidth: '900px', padding: 2 }}>
                {error && <Typography color="error" variant="body2">{error}</Typography>}

                <Grid2 container spacing={2} sx={{ width: '75%' }}>
                    {posts.map(post => (
                        <Grid2 key={post.id} sx={{ width: '100%' }}>
                            <Post post={post} onFetchComments={fetchComments} commentsVisible={commentVisible[post.id]}>
                                {commentVisible[post.id] && comments[post.id] && (
                                    <CommentsSection
                                        onDeleteComment={(commentId) => handleDeleteComment(post.id, commentId)}
                                        postUserId={post.user_id}
                                        postId={post.id}
                                        comments={comments[post.id]}
                                        onNewComment={(newComment) => handleNewComment(post.id, newComment)}
                                    />
                                )}
                            </Post>
                        </Grid2>
                    ))}
                </Grid2>
            </Box>

            <LogoutButton />
        </Box>
    );
};

export default LandingPage;