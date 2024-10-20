import React from 'react';
import { Box, Typography } from '@mui/material';
import ThemeToggleButton from './ThemeToggleButton.tsx';

const Header: React.FC = () => (
    <Box
        sx={{
            width: '100%',
            maxWidth: '900px',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            mb: 2,
        }}
    >
        <Typography variant="h4" component="h1">
            Not Bad Feed
        </Typography>
        <ThemeToggleButton />
    </Box>
);

export default Header;