import React, { useState } from 'react';
import { Tab, Tabs } from 'react-bootstrap';
import './AuthenPage.css';
import LoginTab from '../Login/LoginTab';
import RegisterTab from '../Register/RegisterTab';
import 'bootstrap/dist/css/bootstrap.min.css';

const AuthPage: React.FC = () => {
    const [key, setKey] = useState('login');

    return (
        <div className="container mt-5">
            <h2>Chat</h2>
            <Tabs activeKey={key} onSelect={(k) => {
                // eslint-disable-next-line no-mixed-operators
                if (k === 'login' || k ==='register' && key!== k) {
                    setKey(k)
                }
            }} className="mb-3">
                <Tab eventKey="login" title="Login">
                    <LoginTab />
                </Tab>
                <Tab eventKey="register" title="Register">
                    <RegisterTab />
                </Tab>
            </Tabs>
        </div>
    );
};

export default AuthPage;