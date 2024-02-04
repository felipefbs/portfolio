package portfolio

var (
	fullCycle = &EducationItemProps{
		Type:             "Curso Online",
		Title:            "Desenvolvedor Full Stack && Full Cycle",
		Course:           "Curso Full Cycle",
		Organization:     "Full Cycle",
		OrganizationLink: "https://curso.fullcycle.com.br/curso-fullcycle/",
		Location:         "Online",
		Duration:         "2024",
	}

	telematica = &EducationItemProps{
		Type:             "Formação Academica",
		Title:            "Tecnólogo em Telématica",
		Course:           "Curso superior de Tecnologia em Telématica",
		Organization:     "IFPB",
		OrganizationLink: "http://ifpb.com.br",
		Location:         "Campina Grande - PB",
		Duration:         "2017-2021",
	}

	EducationList = []*EducationItemProps{telematica, fullCycle}
)

type EducationItemProps struct {
	Type             string
	Title            string
	Course           string
	Organization     string
	OrganizationLink string
	Location         string
	Duration         string
}
