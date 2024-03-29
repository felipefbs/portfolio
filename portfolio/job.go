package portfolio

var (
	theia = &JobExperienceItemProps{
		JobTitle:         "Desenvolvedor Back-end Pleno",
		Organization:     "Theia",
		OrganizationLink: "https://www.theia.com.br/",
		Location:         "São Paulo - SP (Remoto)",
		Duration:         "Jun, 2022 - Dez, 2023",
		Responsibilities: []string{
			"Implementação de diversos novos serviços em Go renovando os projetos legados, implementando em uma arquitetura distribuída e garantido seguridade de evolução através de testes robustos",
			"Automatização de processos manuais garantindo maior consistência e padronização",
			"Aplicação de padrões de mercado para rastreabilidade, documentação (utilizando Swagger 2.0) e diversas outras melhorias e padronização de código a fim de melhorar legibilidade e manutenibilidade do código",
		},
		Technologies: []string{"AWS(S3, ECR, SQS, DynamoDB)", "PostgreSQL", "testcontainers", "Github", "Swagger", "JWT", " OAUTH2.0"},
		Languages:    []string{"Go", "SQL", "Kotlin"},
	}
	apoio = &JobExperienceItemProps{
		JobTitle:         "Desenvolvedor Back-end",
		Organization:     "Apoio Ecolimp",
		OrganizationLink: "https://apoioecolimp.com/",
		Location:         "Rio de Janeiro - RJ (Remoto)",
		Duration:         "Fev, 2022 - Jun, 2022",
		Responsibilities: []string{
			"Implementação serviços utilizando Go em conjunto a diversos padrões de mercado para garantir rastreabilidade, consistência e manutenibilidade de código",
			"Ampliação e manutenção de diversos serviços em Typescript utilizados para o gerenciamento de limpeza de salas de hospitais",
		},
		Technologies: []string{"NestJS", "Swagger", "AWS (API Gateway)"},
		Languages:    []string{"Go", "Typescript", "MongoDB"},
	}
	lsd = &JobExperienceItemProps{
		JobTitle:         "Desenvolvedor e Pesquisador",
		Organization:     "LSD - UFCG",
		OrganizationLink: "https://www.lsd.ufcg.edu.br/#/",
		Location:         "Campina Grande - PB",
		Duration:         "Ago, 2021 - Jun, 2022",
		Technologies:     []string{"VMWare", "TPM 2.0", "GNU/Linux"},
		Responsibilities: []string{
			"Pesquisa de soluções de segurança da informação para a plataforma da VMWare baseado na virtualização do chip TPM2.0",
			"Desenvolvimento e aplicação das soluções para garantir melhor integridade de maquinas virtuais utilizando a solução do vTPM utilizando Go+VMWare SDK"},
		Languages: []string{"Go", "Shell Script"},
	}
	resilia = &JobExperienceItemProps{
		JobTitle:         "Facilitador Tech",
		Organization:     "Resilia Educação",
		OrganizationLink: "https://www.resilia.com.br/",
		Location:         "Rio de Janeiro - RJ (Remoto)",
		Duration:         "Mai, 2021 - Dez, 2021",
		Technologies:     []string{"Docker", "express.js", "node.js", "MySQL"},
		Responsibilities: []string{
			"Atuação como Facilitador Tech, transmitindo conhecimento e auxiliando o entendimento de conceitos chave de desenvolvimento web",
			"Compartilhamento de conhecimentos de desenvolvimento backend utilizando Express.JS e Javascript",
			"Orientação na modelagem de bancos de dados utilizando boas práticas para manipulação e organização de dados usando o MySQL",
			"Auxilio sobre o entendimento de gerenciamento de containers Docker de maneira eficiente",
		},
		Languages: []string{"Javascript", "SQL"},
	}
	gcompi = &JobExperienceItemProps{
		JobTitle:         "Discente Bolsista",
		Organization:     "GCOMPI - IFPB",
		OrganizationLink: "https://www.ifpb.edu.br/polodeinovacao/laboratorios/gcompi",
		Location:         "Campina Grande - PB",
		Duration:         "Jan, 2020 - Mai, 2021",
		Technologies:     []string{"InfluxDB", "SCons", "FreeRTOS", "GNU/Linux"},
		Responsibilities: []string{
			"Pesquisa e desenvolvimento no campo de Redes de Sensores Sem Fio, aplicando conhecimentos adquiridos na graduação",
			"Desenvolvimento do firmware utilizando plataforma OpenMote B, ampliando a comunicação via rádio para várias modulações do protocolo IEEE 802.15.4g",
			"Analise dos dados, utilizando Python, coletados que foram utilizados para a escrita do Trabalho de Conclusão de Curso",
		},
		Languages: []string{"C++", "Python", "SQL"},
	}
)

var JobList = []*JobExperienceItemProps{theia, apoio, lsd, resilia, gcompi}

type JobExperienceItemProps struct {
	JobTitle         string
	Organization     string
	OrganizationLink string
	Location         string
	Duration         string
	Technologies     []string
	Responsibilities []string
	Languages        []string
}
