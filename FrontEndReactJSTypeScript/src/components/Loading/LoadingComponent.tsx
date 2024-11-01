import { Spinner } from 'react-bootstrap';
import './LoadingComponent.css';

const LoadingComponent: React.FC = () => {
    return (
        <div className='loading-dialog'>
            <div className='loading-content'>
                <Spinner className='loading-spinner' animation='border' variant='dark' />
            </div>
        </div>
    );
}

export default LoadingComponent;