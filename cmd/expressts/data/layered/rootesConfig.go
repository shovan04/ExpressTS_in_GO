package layered

func GetRoutesConfigContent() []byte {
return []byte(`export class WellcomeRoutes {
    static readonly BASE_PATH: string = "/hello";
    static readonly WELLCOME: string = "/wellcome";
}`)}