import React, { useState } from 'react';
import { Button, Form } from 'react-bootstrap';
import './RegisterTab.css';
import 'bootstrap/dist/css/bootstrap.min.css';

const RegisterTab: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');

    const handleRegister = () => {
        
    };

    return (
        <Form>
            <Form.Group controlId="registerUsername">
                <Form.Label>Username</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="Enter username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
            </Form.Group>

            <Form.Group controlId="registerPassword" className="mt-3">
                <Form.Label>Password</Form.Label>
                <Form.Control
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
            </Form.Group>

            <Form.Group controlId="registerName" className="mt-3">
                <Form.Label>Name</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="Enter your name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
            </Form.Group>

            <Button variant="primary" onClick={handleRegister} className="mt-3 register-btn">
                Register
            </Button>
        </Form>
    );
};

export default RegisterTab;