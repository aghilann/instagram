// src/ColorModeContext.tsx
import React, { createContext, useState, useMemo } from 'react';

interface ColorModeContextType {
    mode: 'light' | 'dark';
    toggleColorMode: () => void;
}

export const ColorModeContext = createContext<ColorModeContextType>({
    mode: 'light',
    toggleColorMode: () => {},
});

export const ColorModeContextProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [mode, setMode] = useState<'light' | 'dark'>('light');

    const toggleColorMode = () => {
        setMode(prevMode => (prevMode === 'light' ? 'dark' : 'light'));
    };

    const value = useMemo(() => ({ mode, toggleColorMode }), [mode]);

    return <ColorModeContext.Provider value={value}>{children}</ColorModeContext.Provider>;
};