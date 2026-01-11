package mvc

func GetMainRouterIndexContent() []byte {
	return []byte(`import { Router, Request, Response } from "express";
import userRoutes from "./users.js";

const mainRouter = Router();

mainRouter.get("/", (req: Request, res: Response): void => {
  res.send("MVC API is running...");
});

// Mount module routers
mainRouter.use("/users", userRoutes);

export default mainRouter;
`)
}
