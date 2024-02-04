package portfolio

var ProjectList = []*ProjectItemProps{
	{
		Title:          "Portfolio",
		Type:           "Projeto Pessoal",
		Description:    []string{"Um portfolio mostrando minhas principais experiências e conquistas na carreira"},
		RepositoryLink: "https://github.com/felipefbs/portfolio",
		Year:           "2024",
		Technologies:   []string{"HTMX", "Templ"},
		Languages:      []string{"Go"},
	},
	{
		Title:          "Teste de integração com testcontainers",
		Type:           "Projeto Pessoal",
		Description:    []string{"Simulando um sistema de comunicação assíncrona utilizando Apache Kafka, criei testes de validação que garantem o desenvolvimento do software de maneira robusta utilizando containers docker."},
		RepositoryLink: "https://github.com/felipefbs/testcontainers",
		Year:           "2024",
		Technologies:   []string{"testcontainers", " Apache Kafka", " PostgreSQL"},
		Languages:      []string{"Go", "SQL"},
	},
	{
		Title:          "Impacto das Modulações do IEEE 802.15.4g na Qualidade de comunicação em ambiente de Smart Building",
		Type:           "Projeto de Graduação",
		Description:    []string{"Utilizando Sistemas Embarcados para verificar parâmetros de telecomunicações para analise de dados da qualidade do enlace sem fio."},
		RepositoryLink: "https://github.com/GComPI-IFPB/openmote-fw",
		Year:           "2020-2021",
		Technologies:   []string{"testcontainers", "InfluxDB", "GNU/Linux"},
		Languages:      []string{"Python", "C++"},
	},
	{
		Title:          "Coronavírus BR",
		Type:           "Projeto pessoal",
		Description:    []string{"Site desenvolvido em conjunto com amigos para ajudar e informar o povo brasileiro sobre a pandemia de COVID-19."},
		RepositoryLink: "https://github.com/henry-ns/coronavirusbr",
		Year:           "2020",
		Technologies:   []string{"ReactJS", "Gatsby.js", "GraphQL"},
		Languages:      []string{"Typescript"},
	},
	{
		Title:          "Tim Maia Bot",
		Type:           "Desafio",
		Description:    []string{"Bot musical para Discord que toca músicas do cantor Tim Maia desenvolvido para o Desafio333 da comunidade do Código Falado."},
		RepositoryLink: "https://github.com/felipefbs/desafio333/tree/master/2020-Bot-Discord/felipefbs",
		Year:           "2020",
		Technologies:   []string{"Discord.js"},
		Languages:      []string{"Javascript"},
	},
	{
		Title:          "Tetris333",
		Type:           "Desafio",
		Description:    []string{"O clássico jogo tetris desenvolvido em Typescript para o Desafio333 da comunidade do Código Falado."},
		RepositoryLink: "https://github.com/henry-ns/tetris333",
		Year:           "2020",
		Technologies:   []string{"p5.js", " ReactJS"},
		Languages:      []string{"Typescript"},
	},
	{
		Title:          "Manipulação de Braço Robótico Usando Acelerômetro e Microcontrolador ESP32 Aplicado ao Ensino da Robótica",
		Type:           "Projeto Acadêmico",
		Description:    []string{"Utilização de um braço robótico controlado por microcontroladores para explicação de conceitos físicos como aceleração da gravidade aplicadas ao ensino básico."},
		RepositoryLink: " ",
		Year:           "2019",
		Technologies:   []string{"Arduino"},
		Languages:      []string{"C"},
	},
}

type ProjectItemProps struct {
	Title          string
	Type           string
	Description    []string
	RepositoryLink string
	Year           string
	Technologies   []string
	Languages      []string
}
