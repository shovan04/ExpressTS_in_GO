package layered

func GetMainRouterIndexContent() []byte {
	return []byte(`import { Router, Request, Response } from "express";
import { WellcomeRoutes } from "../configs/routes/wellcome.js";
import wellcomeRoutes from "./hello/wellcome.js";

const mainRouter = Router();

mainRouter.get("/", (req: Request, res: Response): void => {
  res.send("API is running... Go to /hello/wellcome to see the wellcome message.");
});

mainRouter.use(WellcomeRoutes.BASE_PATH, wellcomeRoutes)

export default mainRouter;

`)}