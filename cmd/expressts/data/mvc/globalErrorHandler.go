package mvc

func GetGlobalErrorHandlerContent() []byte {
	return []byte(`import { ValidationError } from "class-validator";
import { Request, Response, NextFunction } from "express";
import HttpResponseCode from "../utilities/httpResponseCode.js";

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

    return res.status(HttpResponseCode.BAD_REQUEST).json({
        status: false,
        message: "Validation failed",
        data: {
            path,
            status: HttpResponseCode.BAD_REQUEST,
            message: "Validation failed",
            validationErrors
        }
    });
  }

  /* ---------- GENERIC HTTP ERRORS ---------- */
  if (typeof err === "object" && err !== null && "statusCode" in err) {
    const statusCode = (err as any).statusCode || HttpResponseCode.INTERNAL_SERVER_ERROR;
    return res.status(statusCode).json({
        status: false,
        message: "Request failed",
        data: {
             path,
             status: statusCode,
             message: (err as any).message || "Request failed"
        }
    });
  }

  /* ---------- FALLBACK ---------- */
  console.error("Unhandled Error:", err);

  return res.status(HttpResponseCode.INTERNAL_SERVER_ERROR).json({
      status: false,
      message: "An error occurred",
      data: {
          path,
          status: HttpResponseCode.INTERNAL_SERVER_ERROR,
          message: "Internal server error"
      }
  });
}
`)
}
