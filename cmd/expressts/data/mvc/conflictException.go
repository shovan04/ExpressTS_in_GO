package mvc

func GetConflictExceptionContent() []byte {
	return []byte(`import HttpResponseCode from "../httpResponseCode.js";

class ConflictException extends Error {
  statusCode: number;
  code: string;
  override cause?: unknown;

  constructor(message: string, cause?: unknown) {
    super(message);

    this.name = "ConflictException";
    this.code = "CONFLICT_ERROR";
    this.statusCode = HttpResponseCode.CONFLICT;
    this.cause = cause;

    Object.setPrototypeOf(this, new.target.prototype);
    Error.captureStackTrace?.(this, this.constructor);
  }
}

export default ConflictException;
`)
}
