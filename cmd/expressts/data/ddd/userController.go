package ddd

func GetUserControllerContent() []byte {
	return []byte(`
import { Request, Response } from "express";
import { CreateUserUseCase } from "../../../application/use-cases/CreateUserUseCase.js";
import { CreateUserDTO } from "../../../application/dtos/CreateUserDTO.js";

export class UserController {
  constructor(private createUserUseCase: CreateUserUseCase) {}

  public async create(req: Request, res: Response): Promise<void> {
    const dto: CreateUserDTO = req.body;
    const user = await this.createUserUseCase.execute(dto);
    res.status(201).json(user);
  }
}
`)
}
