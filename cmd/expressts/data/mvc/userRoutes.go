package mvc

func GetUserRoutesContent() []byte {
	return []byte(`import { Router } from "express";
import { UserController } from "../controller/users.js";
import { validateDto } from "../middleware/validateDtos.js";
import { CreateUserDTO } from "../utilities/dtos/CreateUserDTO.js";

const router = Router();
const userController = new UserController();

router.post("/", validateDto(CreateUserDTO), userController.save);

export default router;
`)
}
