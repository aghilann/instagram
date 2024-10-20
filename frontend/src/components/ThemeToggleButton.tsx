import React, { useContext } from 'react';
import { IconButton } from '@mui/material';
import { Brightness4, Brightness7 } from '@mui/icons-material';
import { ColorModeContext } from "../ColorModeContext.tsx";

const ThemeToggleButton: React.FC = () => {
    const { mode, toggleColorMode } = useContext(ColorModeContext);

    return (
        <IconButton onClick={toggleColorMode} color="inherit">
            {mode === 'dark' ? <Brightness7 /> : <Brightness4 />}
        </IconButton>
    );
};

export default ThemeToggleButton;