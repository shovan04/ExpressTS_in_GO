package expressts

import (
	"fmt"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/achitecture/ddd"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/achitecture/layered"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/achitecture/mvc"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func Init(project types.ProjectInitStruct) {

	fmt.Println("🚀 Creating project…")
	fmt.Println()
	// Implement project scaffolding based on user choices
	switch project.Options.ProjectArch {
	case "layered":
		//initialize layered architecture
		layered.InitLayeredArchitecture(project)
	case "ddd":
		//initialize ddd architecture
		ddd.InitDDDArchitecture(project)
	case "mvc":
		//initialize mvc architecture
		mvc.InitMVCArchitecture(project)
	default:
		fmt.Println("Unsupported architecture.")
	}
	
}
