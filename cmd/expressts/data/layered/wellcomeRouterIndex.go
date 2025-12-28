package layered

func GetWellcomeRouterIndexContent() []byte {
	return []byte(`import { Router } from "express";
import { sayWellcome } from "../../controllers/hello/wellcome.js";
import { WellcomeRoutes } from "../../config/routes/wellcome.js";

const wellcomeRoutes = Router();

wellcomeRoutes.get(WellcomeRoutes.WELLCOME, sayWellcome);

export default wellcomeRoutes;`)
}
