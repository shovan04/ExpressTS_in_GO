package ddd

func GetUserEntityContent() []byte {
	return []byte(`
export class User {
  constructor(
    public readonly id: string,
    public name: string,
    public email: string,
    public readonly createdAt: Date
  ) {}

  public changeName(newName: string): void {
    if (newName.length < 3) {
      throw new Error("Name is too short");
    }
    this.name = newName;
  }
}
`)
}
