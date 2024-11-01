import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import AuthRoutes from './AuthRoutes';
import DashboardRoutes from './DashboardRoutes';
import Redirect from '../components/Redirect';
import { useEffect, useState } from 'react';

const isAuthenticated = () => {
    try {
        const token = localStorage.getItem('token');
        if (token) {
            return token.length > 0;
        } else {
            return false;
        }
    } catch (error) {
        console.error('Error accessing localStorage', error);
        return false;
    }
}

const AppRoutes: React.FC = () => {
    // Check if user is authenticated before rendering the dashboard or auth routes
    const [isAuth, setAuth] = useState(isAuthenticated());

    useEffect(() => {
        const handleStorageChange = () => {
            console.log('AppRoutes: User authentication status updated', isAuth);
            setAuth(isAuthenticated());
        }

        console.log('Add event listener for storage changes to update user authentication status');
        window.addEventListener('storage', handleStorageChange);

        return () => {
            window.removeEventListener('storage', handleStorageChange);
        }
    });

    return (
        <Router>
            <Routes>
                <Route path='/' element={isAuth ? <DashboardRoutes /> : <Redirect to='/auth' />} />
                <Route path='/auth/*' element={isAuth ? <Redirect to='/' /> : <AuthRoutes />} />
            </Routes>
        </Router>
    );
}

export default AppRoutes;