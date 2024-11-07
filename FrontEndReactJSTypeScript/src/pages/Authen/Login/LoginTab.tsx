import React, { useState } from 'react';
import { Button, Form } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../../../styles/globals.css'
import './LoginTab.css'
import LoginService from '../../../services/LoginService';
import LoadingComponent from '../../../components/Loading/LoadingComponent';
import isSuccessfulResponse from '../../../utils/HttpUtils';

const LoginTab: React.FC = () => {
    const [isLoading, setLoading] = useState(false);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleLogin = async () => {
        try {
            setLoading(true);
            let response = await LoginService(username, password);
            if (isSuccessfulResponse(response.code)) {
                localStorage.setItem('token', response.data.token);
                window.location.reload();
                console.log('Login successful', response);
            } else {
                setError(response.message);
            }
        } catch (error) {
            console.error('Error logging in', error);
            // Handle error appropriately (e.g., display error message to user)
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            <Form>
                <Form.Group controlId="loginUsername">
                    <Form.Label>Username</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="loginPassword" className="mt-3">
                    <Form.Label>Password</Form.Label>
                    <Form.Control
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </Form.Group>

                {error && <div className="error-text">{error}</div>}

                <Button variant="primary" onClick={handleLogin} className="mt-3 login-btn">
                    Login
                </Button>
            </Form>

            {
                isLoading && (<LoadingComponent />)
            }
        </>
    );
};

export default LoginTab;
