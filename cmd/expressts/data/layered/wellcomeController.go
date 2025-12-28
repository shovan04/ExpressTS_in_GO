package layered

func GetWellcomeControllerContent() []byte {
	return []byte(`import { Request, Response } from "express";
import { mapSayWellcomeResponse } from "../../mappers/hello/wellcome.js";
import { HelloService } from "../../services/hello/wellcome.service.js";

export const sayWellcome = (req: Request, res: Response) => {
  const data: string = HelloService.sayWellcome();
    const response = mapSayWellcomeResponse(data);
  res.status(200).json(response);
};`)
}
