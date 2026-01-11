package ddd

func GetCreateUserUseCaseContent() []byte {
	return []byte(`
import { User } from "../../domain/entities/User.js";
import { IUserRepository } from "../../domain/repositories/IUserRepository.js";
import { CreateUserDTO } from "../dtos/CreateUserDTO.js";

export class CreateUserUseCase {
  constructor(private userRepository: IUserRepository) {}

  public async execute(dto: CreateUserDTO): Promise<User> {
    const user = new User(
      crypto.randomUUID(),
      dto.name,
      dto.email,
      new Date()
    );

    await this.userRepository.save(user);
    return user;
  }
}
`)
}
