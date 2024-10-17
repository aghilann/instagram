// src/index.tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { ColorModeContextProvider } from './ColorModeContext';
import ThemeProviderWrapper from "./ThemeProviderWrapper.tsx";
import {CssBaseline} from "@mui/material";

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);

root.render(
    <React.StrictMode>
        <ColorModeContextProvider>
            <ThemeProviderWrapper>
                <CssBaseline />
                <App />
            </ThemeProviderWrapper>
        </ColorModeContextProvider>
    </React.StrictMode>
);