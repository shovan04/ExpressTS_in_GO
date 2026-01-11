package mvc

func GetUserModelContent() []byte {
	return []byte(`import { IsNotEmpty, IsEmail } from "class-validator";

export class User {
    @IsNotEmpty({ message: "Name should not be empty" })
    name!: string;
    
    @IsEmail({}, { message: "Invalid email format" })
    email!: string;

    constructor(name: string, email: string) {
        this.name = name;
        this.email = email;
    }
}
`)
}
