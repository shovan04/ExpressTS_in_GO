package layered

func GetWellcomeRouterIndexContent() []byte {
	return []byte(`import { Router } from "express";
import { sayWellcome } from "../../controllers/hello/wellcome.controller.js";
import { WellcomeRoutes } from "../../configs/routes/wellcome.js";

const wellcomeRoutes = Router();

wellcomeRoutes.get(WellcomeRoutes.WELLCOME, sayWellcome);

export default wellcomeRoutes;`)
}
