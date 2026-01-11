package ddd

func GetInMemoryUserRepositoryContent() []byte {
	return []byte(`
import { User } from "../../domain/entities/User.js";
import { IUserRepository } from "../../domain/repositories/IUserRepository.js";

export class InMemoryUserRepository implements IUserRepository {
  private users: Map<string, User> = new Map();

  public async save(user: User): Promise<void> {
    this.users.set(user.id, user);
  }

  public async findById(id: string): Promise<User | null> {
    return this.users.get(id) || null;
  }

  public async findByEmail(email: string): Promise<User | null> {
    for (const user of this.users.values()) {
      if (user.email === email) {
        return user;
      }
    }
    return null;
  }
}
`)
}
