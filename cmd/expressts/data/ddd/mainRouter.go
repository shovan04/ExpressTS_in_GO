package ddd

func GetMainRouterIndexContent() []byte {
	return []byte(`import { Router, Request, Response } from "express";
import { userRouter } from "./UserRoutes.js";

const mainRouter = Router();

mainRouter.get("/", (req: Request, res: Response): void => {
  res.send("DDD API is running...");
});

// Mount module routers
mainRouter.use("/users", userRouter);

export default mainRouter;
`)
}
