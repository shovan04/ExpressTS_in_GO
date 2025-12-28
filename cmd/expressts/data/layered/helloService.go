package layered


func GetHelloServiceContent() []byte {
	return []byte(`
export class HelloService {
    public static sayWellcome(): string {
        return "Hello, welcome to our service!";
    }
}`)}