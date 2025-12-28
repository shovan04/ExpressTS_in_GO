package layered

func GetWellcomeMapperContent() []byte {
	return []byte(`import { ResponseDTO } from "../../DTOs/response.DTO.js";

export const mapSayWellcomeResponse = (data: string) => {
    const response = new ResponseDTO<string>();
    response.setStatus(true);
    response.setMessage("Success");
    response.setData(data);
    return response;
}`)
}
