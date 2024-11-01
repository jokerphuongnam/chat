export interface ResponseModel<Data> {
    code: number;
    message: string;
    data: Data;
}