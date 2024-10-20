// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import LandingPage from './pages/LandingPage';

const App: React.FC = () => {
    return (
        <Router>
            <div className="App">
                <Routes>
                    {/* Route for the Root URL */}
                    <Route path="/" element={<LandingPage />} />

                    {/* Route for the Login Page */}
                    <Route path="/login" element={<LoginPage />} />

                    {/* Route for the Landing Page (shown after successful login) */}
                    <Route path="/landing" element={<LandingPage />} />
                </Routes>
            </div>
        </Router>
    );
};

export default App;