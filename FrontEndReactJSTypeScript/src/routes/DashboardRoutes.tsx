import { Route, Routes } from 'react-router-dom';
import DashboardPage from '../pages/Dashboard/DashboardPage';

const DashboardRoutes: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<DashboardPage />} />
        </Routes>
    );
}

export default DashboardRoutes;