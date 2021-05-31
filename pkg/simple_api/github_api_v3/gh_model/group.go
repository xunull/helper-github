package gh_model

import "strings"

func GroupRepoListByLang(repos []*Repository) map[string][]*Repository {
	m := make(map[string][]*Repository)
	for _, repo := range repos {

		if repo.Language == "" {
			repo.Language = "Unknown"
		}
		rpLang := strings.ToLower(repo.Language)
		if _, ok := m[rpLang]; !ok {
			m[rpLang] = make([]*Repository, 0)
		}
		m[rpLang] = append(m[rpLang], repo)
	}
	return m
}
