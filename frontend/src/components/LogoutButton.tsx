import React from 'react';
import { Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const LogoutButton: React.FC = () => {
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/');
    };

    return (
        <Button variant="contained" color="primary" onClick={handleLogout} sx={{ mt: 4 }}>
            Logout
        </Button>
    );
};

export default LogoutButton;