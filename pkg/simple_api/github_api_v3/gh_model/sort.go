package gh_model

import "sort"

func SortRepoListByLang(repos []*Repository) {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Language > repos[j].Language
	})
}

func (r RepoList) SortRepoListByCreateTime() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].CreatedAt.After(r[j].CreatedAt.Time)
	})
}
