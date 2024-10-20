import React, { useState } from 'react';
import { Box, Button, IconButton, TextField, Typography } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import axiosInstance from '../axiosConfig';
import { Comment } from "../models/models.tsx";

interface CommentsSectionProps {
    postId: number;
    comments: Comment[];
    onNewComment: (newComment: Comment) => void;
    onDeleteComment: (commentId: number) => void;
    postUserId: number;  // Add a prop for the user ID of the post owner
}

const CommentsSection: React.FC<CommentsSectionProps> = ({ postId, comments, onNewComment, onDeleteComment }) => {
    const [newComment, setNewComment] = useState<string>('');
    const [error, setError] = useState<string | null>(null);
    const currentUserId = Number(localStorage.getItem('userId')); // Get the currently logged-in user's ID

    const handleCommentSubmit = () => {
        if (!newComment.trim()) {
            setError('Comment cannot be empty.');
            return;
        }

        const payload = {
            user_id: currentUserId,  // Use the logged-in user's ID
            post_id: postId,
            content: newComment,
        };

        // Post the comment to the API
        axiosInstance.post("/comment/", payload)
            .then(response => {
                // Clear the input field after successful submission
                setNewComment('');
                setError(null);

                onNewComment(response.data);
            })
            .catch(() => {
                setError('Failed to post comment.');
            });
    };

    const handleDeleteComment = (commentId: number) => {
        // Call the API to delete the comment
        axiosInstance.delete(`http://localhost:8080/comment/${commentId}`)
            .then(() => {
                // Call the parent component's callback to remove the comment from the state
                onDeleteComment(commentId);
            })
            .catch(() => {
                setError('Failed to delete comment.');
            });
    };

    return (
        <Box sx={{ mt: 2 }}>
            {/* Existing comments */}
            {comments.map(comment => (
                <Box key={comment.id} sx={{ mb: 1, ml: 2, display: 'flex', alignItems: 'center' }}>
                    <Typography variant="body2" sx={{ flexGrow: 1 }}>
                        {comment.content}
                    </Typography>
                    {/* Conditionally render the delete button only if the logged-in user matches the post's user_id */}
                    {comment.user_id === currentUserId && (
                        <IconButton
                            aria-label="delete"
                            onClick={() => handleDeleteComment(comment.id)}
                            size="small"
                        >
                            <DeleteIcon fontSize="small" />
                        </IconButton>
                    )}
                    <Typography variant="caption" color="text.secondary" sx={{ mr: 2 }}>
                        {new Date(comment.created_at).toLocaleString()}
                    </Typography>
                </Box>
            ))}

            {/* New comment input */}
            <Box sx={{ display: 'flex', flexDirection: 'column', mt: 2 }}>
                <TextField
                    variant="outlined"
                    placeholder="Add a comment"
                    value={newComment}
                    onChange={(e) => setNewComment(e.target.value)}
                    fullWidth
                    multiline
                    rows={2}
                />
                {error && <Typography color="error" variant="body2">{error}</Typography>}
                <Button
                    variant="contained"
                    color="primary"
                    sx={{ mt: 2, alignSelf: 'flex-end' }}
                    onClick={handleCommentSubmit}
                >
                    Submit Comment
                </Button>
            </Box>
        </Box>
    );
};

export default CommentsSection;