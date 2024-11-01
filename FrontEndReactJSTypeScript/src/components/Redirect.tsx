import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

interface RedirectProps {
    to: string;
}

const Redirect: React.FC<RedirectProps> = ({ to }) => {
    const navigate = useNavigate();
    useEffect(() => {
        console.log("Redirecting to:", to);
        navigate(to);
    });
    return (
        <></>
    )
}


export default Redirect;