package gh_model

import "github.com/xunull/goc/enhance/mapx"

func (r RepoList) GetRepoFullNameMap() map[string]*Repository {
	res := make(map[string]*Repository)
	for _, item := range r {
		res[item.FullName] = item
	}
	return res
}

func (r RepoList) GetRepoFullNameTopic() map[string][]string {
	res := make(map[string][]string)
	for _, item := range r {
		res[item.FullName] = item.Topics
	}
	return res
}

func (r RepoList) GetRepoWithTopic(topic string) map[string][]*Repository {
	fnm := r.GetRepoFullNameMap()
	fnt := r.GetRepoFullNameTopic()
	rmmm := mapx.GetStrRmmmap(fnt)

	res := make(map[string][]*Repository)

	for topic, nameMap := range rmmm {
		res[topic] = make([]*Repository, 0, len(nameMap))
		for name, _ := range nameMap {
			res[topic] = append(res[topic], fnm[name])
		}
	}
	return res
}
