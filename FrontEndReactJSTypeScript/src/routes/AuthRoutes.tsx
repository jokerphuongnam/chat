import { Routes, Route } from "react-router-dom";
import AuthPage from "../pages/Authen/Authen/AuthPage";

const AuthRoutes: React.FC = () => (
    <Routes>
        <Route path="/" element={<AuthPage />} />
        <Route path="/login" element={<AuthPage />} />
        <Route path="/register" element={<AuthPage />} />
    </Routes>
)

export default AuthRoutes;