package main

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/a-h/templ"
	"github.com/felipefbs/portfolio/experience"
	"github.com/felipefbs/portfolio/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	theia = &experience.ExperienceItemProps{
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
	apoio = &experience.ExperienceItemProps{
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
	lsd = &experience.ExperienceItemProps{
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
	resilia = &experience.ExperienceItemProps{
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
	gcompi = &experience.ExperienceItemProps{
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

var xpList = []*experience.ExperienceItemProps{theia, apoio, lsd, resilia, gcompi}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	fileHandler := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileHandler))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/about", http.StatusPermanentRedirect)
	})

	router.Handle("/about", templ.Handler(templates.About()))

	languageList := make([]string, 0)
	for _, item := range xpList {
		languageList = append(languageList, item.Languages...)
	}

	router.Get("/experience", func(w http.ResponseWriter, r *http.Request) {
		skills := r.URL.Query().Get("skill")
		templates.Experience(getXpListFiltered(skills), removeDuplicate[string](languageList), skills).Render(r.Context(), w)
	})

	router.Handle("/projects", templ.Handler(templates.Projects()))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getXpListFiltered(skill string) []*experience.ExperienceItemProps {
	if skill == "" {
		return xpList
	}

	l := make([]*experience.ExperienceItemProps, 0)
	for _, v := range xpList {
		if slices.ContainsFunc[[]string, string](v.Languages, func(s string) bool {
			return strings.ToLower(s) == strings.ToLower(skill)
		}) {
			l = append(l, v)
		}
	}

	return l
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
