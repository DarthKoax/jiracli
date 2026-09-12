package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type VersionService struct {
	client *client.Client
}

func NewVersionService(c *client.Client) *VersionService {
	return &VersionService{client: c}
}

type Version struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	Archived      bool   `json:"archived"`
	Released      bool   `json:"released"`
	ReleaseDate   string `json:"releaseDate,omitempty"`
	Project       string `json:"project,omitempty"`
	ProjectID     int    `json:"projectId,omitempty"`
	Self          string `json:"self,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	ReleaseDescription string `json:"releaseDescription,omitempty"`
	Expired       bool   `json:"expired,omitempty"`
	Overdue       bool   `json:"overdue,omitempty"`
	StartDateMillis   int64 `json:"startDateMillis,omitempty"`
	ReleaseDateMillis int64 `json:"releaseDateMillis,omitempty"`
	UserStartDate     string `json:"userStartDate,omitempty"`
	UserReleaseDate   string `json:"userReleaseDate,omitempty"`
	IsReleased        bool   `json:"isReleased,omitempty"`
	MoveUnfixedIssuesTo string `json:"moveUnfixedIssuesTo,omitempty"`
}

type VersionIssueCount struct {
	Self              string `json:"self"`
	IssuesFixedCount  int    `json:"issuesFixedCount"`
	IssuesAffectedCount int  `json:"issuesAffectedCount"`
	IssueCountWithCustomFieldsChecked int `json:"issueCountWithCustomFieldsChecked"`
}

type VersionMoveInfo struct {
	Self      string    `json:"self"`
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Archived  bool      `json:"archived"`
	Released  bool      `json:"released"`
}

func (s *VersionService) Get(ctx context.Context, versionID string, expand []string) (*Version, error) {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s", url.PathEscape(versionID))
	if len(expand) > 0 {
		path += "?expand=" + joinStrings(expand)
	}
	var version Version
	if err := s.client.Get(ctx, path, &version); err != nil {
		return nil, err
	}
	return &version, nil
}

func (s *VersionService) Create(ctx context.Context, version *Version) (*Version, error) {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return nil, err
	}
	var result Version
	if err := s.client.Post(ctx, "/rest/api/2/version", version, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *VersionService) Update(ctx context.Context, versionID string, version *Version) (*Version, error) {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s", url.PathEscape(versionID))
	var result Version
	if err := s.client.Put(ctx, path, version, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *VersionService) Delete(ctx context.Context, versionID string, moveFixIssuesTo, moveAffectedIssuesTo string) error {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s", url.PathEscape(versionID))
	params := url.Values{}
	if moveFixIssuesTo != "" {
		params.Set("moveFixIssuesTo", moveFixIssuesTo)
	}
	if moveAffectedIssuesTo != "" {
		params.Set("moveAffectedIssuesTo", moveAffectedIssuesTo)
	}
	if len(params) > 0 {
		path += "?" + params.Encode()
	}
	return s.client.Delete(ctx, path)
}

func (s *VersionService) GetRelatedIssueCounts(ctx context.Context, versionID string) (*VersionIssueCount, error) {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s/relatedIssueCounts", url.PathEscape(versionID))
	var result VersionIssueCount
	if err := s.client.Get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *VersionService) GetMoveUnresolvedIssuesVersion(ctx context.Context, versionID string) ([]VersionMoveInfo, error) {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s/moveUnresolvedIssues", url.PathEscape(versionID))
	var versions []VersionMoveInfo
	if err := s.client.Get(ctx, path, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

func (s *VersionService) MoveUnresolvedIssues(ctx context.Context, versionID, moveToVersionID string) error {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s/moveUnresolvedIssues", url.PathEscape(versionID))
	body := map[string]string{"moveTo": moveToVersionID}
	return s.client.Post(ctx, path, body, nil)
}

func (s *VersionService) Merge(ctx context.Context, versionID, moveToVersionID string) error {
	if err := s.client.CheckEndpoint("versions"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/version/%s/mergeto/%s", url.PathEscape(versionID), url.PathEscape(moveToVersionID))
	return s.client.Post(ctx, path, nil, nil)
}
