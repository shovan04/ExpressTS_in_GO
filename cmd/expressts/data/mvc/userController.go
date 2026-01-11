package mvc

func GetUserControllerContent() []byte {
	return []byte(`import { Request, Response, NextFunction } from "express";
import { CreateUserDTO } from "../utilities/dtos/CreateUserDTO.js";

export class UserController {
  save(req: Request, res: Response, next: NextFunction) {
      try {
        const user: CreateUserDTO = req.body;
        // Logic to save user would go here
        res.status(201).json({ message: "User created successfully", data: user });
      } catch (error) {
        next(error);
      }
  }
}
`)
}
