package ddd

func GetUserRepositoryInterfaceContent() []byte {
	return []byte(`
import { User } from "../entities/User.js";

export interface IUserRepository {
  save(user: User): Promise<void>;
  findById(id: string): Promise<User | null>;
  findByEmail(email: string): Promise<User | null>;
}
`)
}
