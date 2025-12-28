package layered

func GetConflictExceptionContent() []byte {
	return  []byte(`class ConflictException extends Error {
  statusCode: number;

  constructor(message: string) {
    super(message);
    this.statusCode = 404;
    Object.setPrototypeOf(this, ConflictException.prototype);
  }
}

export default ConflictException;
`)}