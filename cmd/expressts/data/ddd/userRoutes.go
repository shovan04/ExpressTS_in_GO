package ddd

func GetUserRoutesContent() []byte {
	return []byte(`
import { Router } from "express";
import { UserController } from "../controllers/UserController.js";
import { CreateUserUseCase } from "../../../application/use-cases/CreateUserUseCase.js";
import { InMemoryUserRepository } from "../../../infrastructure/persistence/InMemoryUserRepository.js";

const router = Router();

const userRepository = new InMemoryUserRepository();
const createUserUseCase = new CreateUserUseCase(userRepository);
const userController = new UserController(createUserUseCase);

router.post("/users", (req, res) => userController.create(req, res));

export { router as userRouter };
`)
}
