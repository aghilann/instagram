// src/theme.ts
import { createTheme, ThemeOptions } from '@mui/material/styles';

// Common theme settings
const common: Partial<ThemeOptions> = {
    typography: {
        fontFamily: 'Roboto, sans-serif',
        h4: {
            fontWeight: 600,
        },
        body1: {
            fontSize: '1rem',
        },
        // Add more typography settings as needed
    },
    shape: {
        borderRadius: 8,
    },
};

// Light theme configuration
export const lightTheme = createTheme({
    ...common,
    palette: {
        mode: 'light',
        primary: {
            main: '#E1306C', // Instagram pink
        },
        secondary: {
            main: '#405DE6', // Instagram blue
        },
        background: {
            default: '#fafafa', // Light background
            paper: '#ffffff',
        },
        text: {
            primary: '#262626',
            secondary: '#8e8e8e',
        },
    },
    components: {
        // Customize component styles if needed
        MuiButton: {
            styleOverrides: {
                root: {
                    borderRadius: 20,
                },
            },
        },
        // Add more component customizations as needed
    },
});

// Dark theme configuration
export const darkTheme = createTheme({
    ...common,
    palette: {
        mode: 'dark',
        primary: {
            main: '#E1306C', // Instagram pink
        },
        secondary: {
            main: '#405DE6', // Instagram blue
        },
        background: {
            default: '#121212', // Dark background
            paper: '#1e1e1e',
        },
        text: {
            primary: '#ffffff',
            secondary: '#b0b0b0',
        },
    },
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    borderRadius: 20,
                },
            },
        },
        // Add more component customizations as needed
    },
});