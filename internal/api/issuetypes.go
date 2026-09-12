package api

import (
	"context"

	"github.com/darkkoax/jiracli/internal/client"
)

type IssueTypeService struct {
	client *client.Client
}

func NewIssueTypeService(c *client.Client) *IssueTypeService {
	return &IssueTypeService{client: c}
}

type IssueType struct {
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	Subtask     bool   `json:"subtask,omitempty"`
	AvatarID    int    `json:"avatarId,omitempty"`
	UntranslatedName string `json:"untranslatedName,omitempty"`
}

func (s *IssueTypeService) GetAll(ctx context.Context) ([]IssueType, error) {
	if err := s.client.CheckEndpoint("issuetypes"); err != nil {
		return nil, err
	}
	var issueTypes []IssueType
	if err := s.client.Get(ctx, "/rest/api/2/issuetype", &issueTypes); err != nil {
		return nil, err
	}
	return issueTypes, nil
}

func (s *IssueTypeService) Get(ctx context.Context, issueTypeID string) (*IssueType, error) {
	if err := s.client.CheckEndpoint("issuetypes"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/issuetype/" + issueTypeID
	var issueType IssueType
	if err := s.client.Get(ctx, path, &issueType); err != nil {
		return nil, err
	}
	return &issueType, nil
}

func (s *IssueTypeService) Create(ctx context.Context, issueType *IssueType) (*IssueType, error) {
	if err := s.client.CheckEndpoint("issuetypes"); err != nil {
		return nil, err
	}
	var result IssueType
	if err := s.client.Post(ctx, "/rest/api/2/issuetype", issueType, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *IssueTypeService) Update(ctx context.Context, issueTypeID string, issueType *IssueType) error {
	if err := s.client.CheckEndpoint("issuetypes"); err != nil {
		return err
	}
	path := "/rest/api/2/issuetype/" + issueTypeID
	return s.client.Put(ctx, path, issueType, nil)
}

func (s *IssueTypeService) Delete(ctx context.Context, issueTypeID string, alternativeIssueTypeID string) error {
	if err := s.client.CheckEndpoint("issuetypes"); err != nil {
		return err
	}
	path := "/rest/api/2/issuetype/" + issueTypeID
	if alternativeIssueTypeID != "" {
		path += "?alternativeIssueTypeId=" + alternativeIssueTypeID
	}
	return s.client.Delete(ctx, path)
}

func (s *IssueTypeService) GetAlternativeIssueTypes(ctx context.Context, issueTypeID string) ([]IssueType, error) {
	if err := s.client.CheckEndpoint("issuetypes"); err != nil {
		return nil, err
	}
	path := "/rest/api/2/issuetype/" + issueTypeID + "/alternatives"
	var issueTypes []IssueType
	if err := s.client.Get(ctx, path, &issueTypes); err != nil {
		return nil, err
	}
	return issueTypes, nil
}
