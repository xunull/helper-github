package gh_model

import "strconv"

func StatRepoListByLang(repos []*Repository) map[string]int {
	m := make(map[string]int)
	for _, repo := range repos {
		if repo.Language == "" {
			repo.Language = "Unknown"
		}
		if _, ok := m[repo.Language]; !ok {
			m[repo.Language] = 0
		}
		m[repo.Language] += 1
	}
	return m
}

func (r RepoList) StatRepoListByByCreateTime() map[string]int {
	m := make(map[string]int)
	for _, repo := range r {
		year := strconv.Itoa(repo.CreatedAt.Time.Year())

		if _, ok := m[year]; !ok {
			m[year] = 0
		}
		m[year] += 1
	}
	return m
}
