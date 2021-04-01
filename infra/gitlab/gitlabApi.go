package gitlab

import (
	"fmt"

	constants "fastshop.com.br/create_pipelines/domain/Constants"

	gl "github.com/xanzy/go-gitlab"
)

// Gitlab strut
type Gitlab struct {
	token     string
	Client    *gl.Client
	idProject int
}

// Commit Commit
func (g *Gitlab) Commit(nameFile string, action gl.FileAction, content string, commitMessage string, nameBranch string) *gl.Branch {
	branch, _, _ := g.Client.Branches.GetBranch(g.idProject, nameBranch)
	_, _, err := g.Client.Commits.CreateCommit(g.idProject, &gl.CreateCommitOptions{
		Branch:        &branch.Name,
		CommitMessage: &commitMessage,
		Actions: []*gl.CommitAction{
			&gl.CommitAction{
				Action:   action,
				FilePath: nameFile,
				Content:  content,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	return branch
}

// CreateMergeRequest create Merge
func (g *Gitlab) CreateMergeRequest(sourceBranch string, targetBranch string, title string) {
	_, _, err := g.Client.MergeRequests.CreateMergeRequest(g.idProject, &gl.CreateMergeRequestOptions{
		Title:        &title,
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
	})

	if err != nil {
		fmt.Println(err)
	}
}

// CreateBranch create new branch
func (g *Gitlab) CreateBranch(branchName string, ref string) {

	_, _, err := g.Client.Branches.CreateBranch(g.idProject, &gl.CreateBranchOptions{
		Branch: &branchName,
		Ref:    &ref,
	})

	fmt.Println(err)
}

// StartClient client start
func (g *Gitlab) StartClient(idProject int) *Gitlab {
	git, err := gl.NewClient(g.token, gl.WithBaseURL("https://code.fastshopdev.com/api/v4"))
	if err != nil {
		fmt.Println(err)
	}

	g.Client = git
	g.idProject = idProject
	return g
}

// NewGitlab Ioc
func NewGitlab() *Gitlab {
	return &Gitlab{
		token: constants.GitlabToken,
	}
}
