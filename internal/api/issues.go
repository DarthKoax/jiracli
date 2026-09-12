package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darthkoax/jiracli/internal/client"
)

type IssueService struct {
	client *client.Client
}

func NewIssueService(c *client.Client) *IssueService {
	return &IssueService{client: c}
}

type Issue struct {
	ID          string                 `json:"id,omitempty"`
	Key         string                 `json:"key,omitempty"`
	Self        string                 `json:"self,omitempty"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
	Expand      string                 `json:"expand,omitempty"`
	Transitions []Transition           `json:"transitions,omitempty"`
}

type Transition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	To   Status `json:"to"`
}

type Status struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Comment struct {
	ID      string `json:"id,omitempty"`
	Body    string `json:"body"`
	Author  *User  `json:"author,omitempty"`
	Created string `json:"created,omitempty"`
}

type Worklog struct {
	ID           string `json:"id,omitempty"`
	TimeSpent    string `json:"timeSpent,omitempty"`
	TimeSpentSec int    `json:"timeSpentSeconds,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Started      string `json:"started,omitempty"`
	Author       *User  `json:"author,omitempty"`
}

type Watchers struct {
	Self       string `json:"self,omitempty"`
	IsWatching bool   `json:"isWatching"`
	WatchCount int    `json:"watchCount"`
	Watchers   []User `json:"watchers"`
}

type Vote struct {
	Self     string `json:"self,omitempty"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}

type IssueSearchResult struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

func (s *IssueService) Get(ctx context.Context, issueKeyOrID string, fields []string, expand []string) (*Issue, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s", url.PathEscape(issueKeyOrID))
	params := url.Values{}
	if len(fields) > 0 {
		params.Set("fields", joinStrings(fields))
	}
	if len(expand) > 0 {
		params.Set("expand", joinStrings(expand))
	}
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var issue Issue
	if err := s.client.Get(ctx, path, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (s *IssueService) Create(ctx context.Context, issue *Issue) (*Issue, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	var result Issue
	if err := s.client.Post(ctx, "/rest/api/2/issue", issue, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *IssueService) Update(ctx context.Context, issueKeyOrID string, issue *Issue) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s", url.PathEscape(issueKeyOrID))
	return s.client.Put(ctx, path, issue, nil)
}

func (s *IssueService) Delete(ctx context.Context, issueKeyOrID string, deleteSubtasks bool) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s", url.PathEscape(issueKeyOrID))
	if deleteSubtasks {
		path += "?deleteSubtasks=true"
	}
	return s.client.Delete(ctx, path)
}

func (s *IssueService) GetTransitions(ctx context.Context, issueKeyOrID string) ([]Transition, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/transitions", url.PathEscape(issueKeyOrID))
	var result struct {
		Transitions []Transition `json:"transitions"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Transitions, nil
}

func (s *IssueService) DoTransition(ctx context.Context, issueKeyOrID string, transitionID string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/transitions", url.PathEscape(issueKeyOrID))
	body := map[string]interface{}{
		"transition": map[string]interface{}{
			"id": transitionID,
		},
	}
	return s.client.Post(ctx, path, body, nil)
}

func (s *IssueService) GetComments(ctx context.Context, issueKeyOrID string) ([]Comment, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/comment", url.PathEscape(issueKeyOrID))
	var result struct {
		Comments []Comment `json:"comments"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Comments, nil
}

func (s *IssueService) AddComment(ctx context.Context, issueKeyOrID string, comment *Comment) (*Comment, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/comment", url.PathEscape(issueKeyOrID))
	var result Comment
	if err := s.client.Post(ctx, path, comment, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *IssueService) UpdateComment(ctx context.Context, issueKeyOrID, commentID string, comment *Comment) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/comment/%s", url.PathEscape(issueKeyOrID), url.PathEscape(commentID))
	return s.client.Put(ctx, path, comment, nil)
}

func (s *IssueService) DeleteComment(ctx context.Context, issueKeyOrID, commentID string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/comment/%s", url.PathEscape(issueKeyOrID), url.PathEscape(commentID))
	return s.client.Delete(ctx, path)
}

func (s *IssueService) GetWorklogs(ctx context.Context, issueKeyOrID string) ([]Worklog, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/worklog", url.PathEscape(issueKeyOrID))
	var result struct {
		Worklogs []Worklog `json:"worklogs"`
	}
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Worklogs, nil
}

func (s *IssueService) AddWorklog(ctx context.Context, issueKeyOrID string, worklog *Worklog) (*Worklog, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/worklog", url.PathEscape(issueKeyOrID))
	var result Worklog
	if err := s.client.Post(ctx, path, worklog, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *IssueService) UpdateWorklog(ctx context.Context, issueKeyOrID, worklogID string, worklog *Worklog) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/worklog/%s", url.PathEscape(issueKeyOrID), url.PathEscape(worklogID))
	return s.client.Put(ctx, path, worklog, nil)
}

func (s *IssueService) DeleteWorklog(ctx context.Context, issueKeyOrID, worklogID string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/worklog/%s", url.PathEscape(issueKeyOrID), url.PathEscape(worklogID))
	return s.client.Delete(ctx, path)
}

func (s *IssueService) AddAttachment(ctx context.Context, issueKeyOrID string, filePath string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	return fmt.Errorf("multipart attachment upload not yet implemented - requires multipart/form-data")
}

func (s *IssueService) GetWatchers(ctx context.Context, issueKeyOrID string) (*Watchers, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/watchers", url.PathEscape(issueKeyOrID))
	var watchers Watchers
	if err := s.client.Get(ctx, path, &watchers); err != nil {
		return nil, err
	}
	return &watchers, nil
}

func (s *IssueService) AddWatcher(ctx context.Context, issueKeyOrID, username string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/watchers", url.PathEscape(issueKeyOrID))
	return s.client.Post(ctx, path, username, nil)
}

func (s *IssueService) RemoveWatcher(ctx context.Context, issueKeyOrID, username string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/watchers?username=%s", url.PathEscape(issueKeyOrID), url.QueryEscape(username))
	return s.client.Delete(ctx, path)
}

func (s *IssueService) GetVotes(ctx context.Context, issueKeyOrID string) (*Vote, error) {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/votes", url.PathEscape(issueKeyOrID))
	var vote Vote
	if err := s.client.Get(ctx, path, &vote); err != nil {
		return nil, err
	}
	return &vote, nil
}

func (s *IssueService) AddVote(ctx context.Context, issueKeyOrID string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/votes", url.PathEscape(issueKeyOrID))
	return s.client.Post(ctx, path, nil, nil)
}

func (s *IssueService) RemoveVote(ctx context.Context, issueKeyOrID string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/votes", url.PathEscape(issueKeyOrID))
	return s.client.Delete(ctx, path)
}

func (s *IssueService) Assign(ctx context.Context, issueKeyOrID, assignee string) error {
	if err := s.client.CheckEndpoint("issues"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/issue/%s/assignee", url.PathEscape(issueKeyOrID))
	body := map[string]interface{}{"name": assignee}
	return s.client.Put(ctx, path, body, nil)
}

func joinStrings(strs []string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += ","
		}
		result += s
	}
	return result
}
