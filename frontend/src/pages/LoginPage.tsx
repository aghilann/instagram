import React, { useState, useEffect } from 'react';
import { Box, Button, TextField, Typography, FormControl, FormLabel, Link, Divider, CssBaseline } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { styled } from '@mui/material/styles';
import axiosInstance from '../axiosConfig';

const Card = styled(Box)(({ theme }) => ({
    display: 'flex',
    flexDirection: 'column',
    width: '100%',
    padding: theme.spacing(4),
    gap: theme.spacing(2),
    boxShadow:
        'hsla(220, 30%, 5%, 0.05) 0px 5px 15px 0px, hsla(220, 25%, 10%, 0.05) 0px 15px 35px -5px',
    borderRadius: theme.shape.borderRadius,
    [theme.breakpoints.up('sm')]: {
        maxWidth: '450px',
        margin: 'auto',
    },
}));

const SignInContainer = styled(Box)(({ theme }) => ({
    height: '100vh',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    padding: theme.spacing(2),
    backgroundColor: theme.palette.background.default,
}));

const LoginPage: React.FC = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const savedToken = localStorage.getItem('token');
        if (savedToken) {
            navigate('/landing'); // Redirect to the landing page if already logged in
        }
    }, [navigate]);

    const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axiosInstance.post('/auth/login', { email, password });
            const receivedToken = response.data.token;
            const receivedUserId = response.data.id;
            localStorage.setItem('token', receivedToken);
            localStorage.setItem('userId', receivedUserId);
            setError(null);
            navigate('/landing');
        } catch (err) {
            setError('Login failed, please check your credentials.');
        }
    };

    return (
        <SignInContainer>
            <CssBaseline />
            <Card>
                <Typography component="h1" variant="h4" textAlign="center">
                    Sign In
                </Typography>
                <Box component="form" onSubmit={handleLogin} noValidate>
                    <FormControl fullWidth sx={{ mb: 2 }}>
                        <FormLabel htmlFor="email">Email</FormLabel>
                        <TextField
                            id="email"
                            type="email"
                            name="email"
                            placeholder="your@email.com"
                            required
                            fullWidth
                            variant="outlined"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            error={!!error}
                        />
                    </FormControl>

                    <FormControl fullWidth sx={{ mb: 2 }}>
                        <FormLabel htmlFor="password">Password</FormLabel>
                        <TextField
                            id="password"
                            type="password"
                            name="password"
                            placeholder="••••••"
                            required
                            fullWidth
                            variant="outlined"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            error={!!error}
                        />
                    </FormControl>

                    {error && (
                        <Typography color="error" variant="body2">
                            {error}
                        </Typography>
                    )}
                    <Button type="submit" fullWidth variant="contained" color="primary" sx={{ mt: 2 }}>
                        Sign In
                    </Button>
                </Box>

                <Link href="#" variant="body2" sx={{ mt: 2, display: 'block', textAlign: 'center' }}>
                    Forgot your password?
                </Link>

                <Divider sx={{ my: 2 }}>or</Divider>

                <Button fullWidth variant="outlined" onClick={() => alert('Sign in with Google')}>
                    Sign in with Google
                </Button>
            </Card>
        </SignInContainer>
    );
};

export default LoginPage;