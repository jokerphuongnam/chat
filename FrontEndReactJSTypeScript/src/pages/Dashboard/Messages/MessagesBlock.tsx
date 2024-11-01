import 'bootstrap/dist/css/bootstrap.min.css';
import { useEffect } from 'react';

interface MessagesBlockProps {
    roomId: string;
}

const MessagesBlock: React.FC<MessagesBlockProps> = ({ roomId }) => {
    useEffect(() => {

    });

    return (
        <>
            <div className='mesesages'>

            </div>
        </>
    );
}

export default MessagesBlock;