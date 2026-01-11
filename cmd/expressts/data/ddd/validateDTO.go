package ddd

func GetValidateDTOContent() []byte {
	return []byte(`import { plainToInstance } from "class-transformer";
import { validate } from "class-validator";
import { NextFunction, Request, Response } from "express";
import ConflictException from "../../../domain/exceptions/ConflictException.js";

type ReqSource = "body" | "query" | "params";

export const validateDto = (DTOClass: any, source: ReqSource = "body") => {
  return async (req: Request, _res: Response, next: NextFunction) => {
    try {
      const rawData = req[source];

      if (rawData === undefined) {
        throw new ConflictException("DTO validation middleware applied but " + source + " is undefined");
      }

      const dtoObj = plainToInstance(DTOClass, rawData, {
        enableImplicitConversion: true,
        exposeDefaultValues: true,
      });

      const errors = await validate(dtoObj, {
        whitelist: true,
        forbidNonWhitelisted: true,
        forbidUnknownValues: true,
        validationError: {
          target: false,
          value: false,
        },
      });

      if (errors.length > 0) {
        return next(errors);
      }

      // overwrite the same source so next handlers use validated + transformed data
      (req as any)[source] = dtoObj;

      next();
    } catch (err) {
      next(err);
    }
  };
};
`)
}
