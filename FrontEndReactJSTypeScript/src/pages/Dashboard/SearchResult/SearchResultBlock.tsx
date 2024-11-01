import 'bootstrap/dist/css/bootstrap.min.css';
import './SearchResultBlock.css'
import React, { useEffect, useState } from 'react';
import SearchService, { SearchResponse } from '../../../services/SearchService';
import isSuccessfulResponse from '../../../utils/HttpUtils';
import { Alert, ListGroup, Spinner, Stack } from 'react-bootstrap';

interface SearchResultBlockProps {
    searchQuery: string;
    onClicked: (roomId: string) => void;
}

const SearchResultBlock: React.FC<SearchResultBlockProps> = ({ searchQuery, onClicked }) => {
    const [isLoading, setLoading] = useState(false);
    const [searchResults, setSearchResults] = useState<SearchResponse[]>([]);
    const [error, setError] = useState('');
    const [currentSearchQuery, setCurrentSearchQuery] = useState('');

    const fetchSearchResults = async () => {
        setLoading(true);
        try {
            const response = await SearchService(searchQuery);
            setCurrentSearchQuery(searchQuery);
            if (isSuccessfulResponse(response.code)) {
                setSearchResults(response.data);
            } else {
                setError(response.message);
            }
        } catch (error) {
            console.error('Error fetching search results', error);
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        if (searchQuery !== currentSearchQuery) {
            if (searchQuery.length > 0) {
                fetchSearchResults();
            } else {
                setLoading(false);
                setSearchResults([]);
                setError('');
            }
        }
    }, [searchQuery, currentSearchQuery]);

    return (
        <>
            <div className='search-result-container h-100'>
                {
                    isLoading ? (
                        <Alert variant="error" className='non-scroll-container'>
                            <Spinner className='loading-spinner' animation='border' variant='dark' />
                        </Alert>) : error ? <Alert variant="error" className='non-scroll-container'>
                            {error}
                        </Alert> : searchResults && searchResults.length > 0 ? (
                            <ListGroup className='search-result-list' variant='flush'>
                                {searchResults.map((item, index) => (
                                    <div onClick={() => onClicked(item.id)}>
                                        <SearchResultItem
                                            key={item.id}
                                            avatarUrl={item.avatarUrl}
                                            name={item.name}

                                        />
                                    </div>
                                ))}
                            </ListGroup>
                        ) : (
                        <Alert variant="info" className='non-scroll-container'>
                            No items to display.
                        </Alert>
                    )
                }
            </div>
        </>
    );
}

interface SearchResultItemProps {
    avatarUrl: string;
    name: string;
}

const SearchResultItem: React.FC<SearchResultItemProps> = ({ avatarUrl, name }) => {
    return (
        <>
            <Stack direction='horizontal' className='search-result-item'>
                <img className='search-result-avatar' src={avatarUrl} alt='avatar' />
                <div className='search-result-name'>{name}</div>
            </Stack>
        </>
    );
}

export default SearchResultBlock;