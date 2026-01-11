package ddd

func GetGlobalErrorHandlerContent() []byte {
	return []byte(`import { ValidationError } from "class-validator";
import { Request, Response, NextFunction } from "express";
import ConstraintValidationErrorResponse from "../../../domain/exceptions/ValidationException.js";
import HttpResponseCode from "../../../constants/httpResponseCode.js";
import { ResponseDTO } from "../../../application/dtos/ResponseDTO.js";
import ConflictException from "../../../domain/exceptions/ConflictException.js";
import ErrorResponseDTO from "../../../application/dtos/ErrorResponseDTO.js";

/* -------- helper to extract nested validation errors -------- */
function extractValidationErrors(
  errors: ValidationError[],
  parentPath = ""
): Record<string, string> {
  const result: Record<string, string> = {};

  for (const error of errors) {
    const path = parentPath
      ? parentPath + "." + error.property
      : error.property;

    if (error.constraints) {
      result[path] = Object.values(error.constraints).join(", ");
    }

    if (error.children && error.children.length > 0) {
      Object.assign(result, extractValidationErrors(error.children, path));
    }
  }

  return result;
}

/* ------------------ GLOBAL ERROR HANDLER ------------------ */
export default function GlobalErrorHandler(
  err: unknown,
  req: Request,
  res: Response,
  _next: NextFunction
) {
  const path = req.originalUrl;

  /* ---------- DTO VALIDATION ERRORS ---------- */
  if (
    Array.isArray(err) &&
    err.length > 0 &&
    err.every((e) => e instanceof ValidationError)
  ) {
    const validationErrors = extractValidationErrors(err);

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

  /* ---------- CUSTOM CONFLICT ERROR ---------- */
  if (err instanceof ConflictException) {
    const errorDto = new ResponseDTO<ErrorResponseDTO>();
    errorDto.setStatus(false);
    errorDto.setMessage("Conflict occurred");
    errorDto.setData(
      new ErrorResponseDTO(
        path,
        HttpResponseCode.CONFLICT,
        err.message || "Conflict occurred"
      )
    );

    return res.status(HttpResponseCode.CONFLICT).json(errorDto);
  }

  /* ---------- INVALID JSON BODY ---------- */
  if (
    typeof err === "object" &&
    err !== null &&
    (err as any).type === "entity.parse.failed"
  ) {
    const errorDto = new ResponseDTO<ErrorResponseDTO>();
    errorDto.setStatus(false);
    errorDto.setMessage("Invalid JSON payload");
    errorDto.setData(
      new ErrorResponseDTO(
        path,
        HttpResponseCode.BAD_REQUEST,
        "Malformed JSON in request body"
      )
    );

    return res.status(HttpResponseCode.BAD_REQUEST).json(errorDto);
  }

  /* ---------- GENERIC HTTP ERRORS ---------- */
  if (typeof err === "object" && err !== null && "statusCode" in err) {
    const statusCode =
      (err as any).statusCode || HttpResponseCode.INTERNAL_SERVER_ERROR;

    const errorDto = new ResponseDTO<ErrorResponseDTO>();
    errorDto.setStatus(false);
    errorDto.setMessage("Request failed");
    errorDto.setData(
      new ErrorResponseDTO(
        path,
        statusCode,
        (err as any).message || "Request failed"
      )
    );

    return res.status(statusCode).json(errorDto);
  }

  /* ---------- FALLBACK ---------- */
  console.error("Unhandled Error:", err);

  const errorDto = new ResponseDTO<ErrorResponseDTO>();
  errorDto.setStatus(false);
  errorDto.setMessage("An error occurred");
  errorDto.setData(
    new ErrorResponseDTO(
      path,
      HttpResponseCode.INTERNAL_SERVER_ERROR,
      "Internal server error"
    )
  );
  return res.status(HttpResponseCode.INTERNAL_SERVER_ERROR).json(errorDto);
}
`)
}
