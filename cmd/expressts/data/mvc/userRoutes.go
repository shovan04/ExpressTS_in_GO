package mvc

func GetUserRoutesContent() []byte {
	return []byte(`import { Router } from "express";
import { UserController } from "../controller/users.js";
import { validateDto } from "../middleware/validateDtos.js";
import { User } from "../model/users.js";

const router = Router();
const userController = new UserController();

router.post("/", validateDto(User), (req, res, next) => userController.save(req, res, next));

export default router;
`)
}
