package mvc

func GetCreateUserDTOContent() []byte {
	return []byte(`import { IsNotEmpty, IsEmail } from "class-validator";

export class CreateUserDTO {
    @IsNotEmpty({ message: "Name should not be empty" })
    name!: string;
    
    @IsEmail({}, { message: "Invalid email format" })
    email!: string;
}
`)
}
