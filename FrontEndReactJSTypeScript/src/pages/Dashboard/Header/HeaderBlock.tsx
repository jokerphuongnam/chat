import 'bootstrap/dist/css/bootstrap.min.css';
import './HeaderBlock.css'
import { Stack } from 'react-bootstrap';

interface HeaderBlockProps {
    avatarUrl: string;
    name: string;
}

const HeaderBlock: React.FC<HeaderBlockProps> = ({ avatarUrl, name }) => {
    console.log('HeaderBlock props:', avatarUrl, name);
    return (
        <>
            <Stack className='header-container' direction='horizontal'>
                <Stack className='header-left' direction='horizontal'>
                    <img className='header-avatar' src={avatarUrl} alt='avatar' />
                    <p className='header-name'>{name}</p>
                </Stack>
                <Stack className='header-right' direction='horizontal'>

                </Stack>
            </Stack>
        </>
    );
}

export default HeaderBlock;