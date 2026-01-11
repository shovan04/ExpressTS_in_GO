package mvc

func GetErrorResponseDTOContent() []byte {
	return []byte(`export default class ErrorResponseDTO {
    private readonly status: number;
    private readonly apiPath: string;
    private readonly message: string;
    private readonly timestamp: string;

    constructor(apiPath: string, status: number, message: string) {
        this.status = status;
        this.apiPath = apiPath;
        this.message = message;
        this.timestamp = new Date().toISOString();
    }
}
`)
}
