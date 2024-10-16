// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginPage from './components/LoginPage';
import LandingPage from './components/LandingPage';

const App: React.FC = () => {
    return (
        <Router>
            <div className="App">
                <Routes>
                    {/* Route for the Login Page */}
                    <Route path="/" element={<LoginPage />} />

                    {/* Route for the Landing Page (shown after successful login) */}
                    <Route path="/landing" element={<LandingPage />} />
                </Routes>
            </div>
        </Router>
    );
};

export default App;