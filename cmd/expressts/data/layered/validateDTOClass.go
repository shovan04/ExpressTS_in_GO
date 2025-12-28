package layered

func GetValidateDTOClassesContent() []byte {
	return []byte(`import { plainToInstance } from "class-transformer";
import { validate } from "class-validator";
import { NextFunction, Request, Response } from "express";

export const validateDto = (
  DTOClasses: any,
  source: "body" | "query" | "params" = "body"
) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    try {
      const dtoObj = plainToInstance(DTOClasses, req[source]);
      const errors = await validate(dtoObj);

      if (errors.length > 0) {
        return next(errors);
      }
      // assign the validated DTO back, so other middlewares can trust it
      Object.assign(req[source], dtoObj);
      next();
    } catch (error) {
      next(error);
    }
  };
};
`)
}
