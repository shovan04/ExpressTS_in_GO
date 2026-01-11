package ddd

func GetCreateUserDTOContent() []byte {
	return []byte(`
export interface CreateUserDTO {
  name: string;
  email: string;
}
`)
}
