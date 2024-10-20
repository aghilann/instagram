import React from 'react';
import { Avatar, Box, Card, CardContent, CardMedia, IconButton, Typography } from '@mui/material';
import FavoriteIcon from '@mui/icons-material/Favorite';
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';
import {FeedPost} from "../models/models.tsx";

interface PostProps {
    post: FeedPost;
    onFetchComments: (postId: number) => void;
    commentsVisible: boolean;
    children: React.ReactNode;
}

const PostCard: React.FC<PostProps> = ({ post, onFetchComments, commentsVisible, children }) => (
    console.log(post.created_at),
    <Card sx={{ width: '100%', borderRadius: 4, boxShadow: 3, bgcolor: 'background.paper' }}>
        {/* User Info */}
        <Box sx={{ display: 'flex', alignItems: 'center', padding: 1 }}>
            <Avatar
                sx={{ width: 40, height: 40, marginRight: 2 }}
                src={post.profile_image}
                alt="User Avatar"
            />
            <Typography variant="body1">
                {post.username}
            </Typography>
        </Box>

        {/* Image */}
        <CardMedia
            component="img"
            height="400"
            image={post.image_url}
            alt={post.caption}
            sx={{ borderRadius: 2 }}
        />

        {/* Caption and Created At */}
        <CardContent>
            <Typography variant="body1">
                {post.caption}
            </Typography>
            <Typography variant="body2" color="text.secondary">
                {new Date(post.created_at).toLocaleString()}
            </Typography>

            {/* Like and Comment Icons */}
            <Box sx={{ display: 'flex', justifyContent: 'flex-start', marginTop: 1 }}>
                <IconButton aria-label="like" color="primary">
                    <FavoriteIcon />
                </IconButton>
                <IconButton aria-label="comment" color="primary" onClick={() => onFetchComments(post.id)}>
                    <ChatBubbleOutlineIcon />
                </IconButton>
            </Box>

            {/* Comments Section */}
            {commentsVisible && children}
        </CardContent>
    </Card>
);

export default PostCard;