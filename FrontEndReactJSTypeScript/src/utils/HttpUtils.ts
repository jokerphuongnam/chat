const isSuccessfulResponse = (statusCode: number): boolean => {
    return statusCode >= 200 && statusCode < 300;
}
 
export default isSuccessfulResponse;