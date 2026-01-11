package ddd

func GetUserRoutesContent() []byte {
	return []byte(`import { Router } from "express";
import { UserController } from "../controllers/UserController.js";
import { CreateUserUseCase } from "../../../application/use-cases/CreateUserUseCase.js";
import { InMemoryUserRepository } from "../../../infrastructure/persistence/InMemoryUserRepository.js";
import { validateDto } from "../middlewares/ValidateDto.js";
import { CreateUserDTO } from "../../../application/dtos/CreateUserDTO.js";

const router = Router();

const userRepository = new InMemoryUserRepository();
const createUserUseCase = new CreateUserUseCase(userRepository);
const userController = new UserController(createUserUseCase);

router.post("/", validateDto(CreateUserDTO), userController.create);

export { router as userRouter };
`)
}
