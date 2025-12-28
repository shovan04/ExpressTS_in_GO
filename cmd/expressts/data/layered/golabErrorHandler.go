package layered

func GetGlobalErrorHandlerContent() []byte {
	return []byte(`import { ValidationError } from "class-validator";
import { Request, Response, NextFunction } from "express";
import ConstraintValidationErrorResponse from "../exceptions/valodationErrors.js";
import HttpResponseCode from "../constants/httpResponseCode.js";
import { ResponseDTO } from "../DTOClass/response.DTO.js";
import ConflictException from "../exceptions/conflictExceptions.js";
import ErrorResponseDTO from "../DTOClass/errorResponse.DTO.js";

export default function GlobalErrorHandler(
  err: any,
  req: Request,
  res: Response,
  next: NextFunction
) {
  const path = req.originalUrl;
  // Class-validator error (e.g., DTO validation)
  if (Array.isArray(err) && err[0] instanceof ValidationError) {
    const validationErrors: Record<string, string> = {};
    err.forEach((e: ValidationError) => {
      if (e.constraints) {
        validationErrors[e.property] = Object.values(e.constraints).join(", ");
      }
    });

    const errorDto = new ResponseDTO<ConstraintValidationErrorResponse>();
    errorDto.setStatus(false);
    errorDto.setMessage("Validation failed");
    errorDto.setData(
      new ConstraintValidationErrorResponse(
        path,
        HttpResponseCode.BAD_REQUEST,
        "Validation failed",
        validationErrors
      )
    );
    return res.status(HttpResponseCode.BAD_REQUEST).json(errorDto);
  }

  // Conflict error
  if (err instanceof ConflictException) {
    const errorDto = new ResponseDTO<ErrorResponseDTO>();
    errorDto.setStatus(false);
    errorDto.setMessage("Conflict occurred");
    errorDto.setData(
      new ErrorResponseDTO(
        path,
        HttpResponseCode.CONFLICT,
        err?.message
      )
    );
    return res.status(HttpResponseCode.CONFLICT).json(errorDto);
  }

  // Generic error handler
  const errorDto = new ResponseDTO<ErrorResponseDTO>();
  errorDto.setStatus(false);
  errorDto.setMessage("An error occurred");
  errorDto.setData(
    new ErrorResponseDTO(
      path,
      HttpResponseCode.INTERNAL_SERVER_ERROR,
      err?.message
    )
  );
  return res.status(HttpResponseCode.INTERNAL_SERVER_ERROR).json(errorDto);
}`)
}