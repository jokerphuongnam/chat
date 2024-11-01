import 'bootstrap/dist/css/bootstrap.min.css';
import './SearchBlock.css'

interface SearchBlockProps {
    searchQuery: string;
    setSearchQuery: (searchQuery: string) => void;
    onFocus: (isFocus: boolean) => void;
}

const SearchBlock: React.FC<SearchBlockProps> = ({ searchQuery, setSearchQuery, onFocus }) => {
    return (
        <>
            <input
                type="text"
                value={searchQuery}
                onFocus={() => { onFocus(true); }}
                onBlur={(event) => { onFocus(false); }}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="Search for a room"
                className='search-input w-100'
            />
        </>
    )
}


export default SearchBlock;