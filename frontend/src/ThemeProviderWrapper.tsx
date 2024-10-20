// src/ThemeProviderWrapper.tsx
import React, { useContext } from 'react';
import { ThemeProvider } from '@mui/material/styles';
import { lightTheme, darkTheme } from './theme';
import { ColorModeContext } from './ColorModeContext';

const ThemeProviderWrapper: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const { mode } = useContext(ColorModeContext);
    const theme = mode === 'light' ? lightTheme : darkTheme;

    return <ThemeProvider theme={theme}>{children}</ThemeProvider>;
};

export default ThemeProviderWrapper;