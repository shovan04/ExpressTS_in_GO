package mvc

func GetResponseDTOContent() []byte {
	return []byte(`export class ResponseDTO<T> {
  private status!: boolean;
  private message!: string;
  private data!: T;

  setStatus(status: boolean) {
    this.status = status;
  }

  setMessage(message: string) {
    this.message = message;
  }

  setData(data: T) {
    this.data = data;
  }

  getStatus(): boolean {
    return this.status;
  }

  getMessage(): string {
    return this.message;
  }

  getData(): T {
    return this.data;
  }
}
`)
}
