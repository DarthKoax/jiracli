package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/darkkoax/jiracli/internal/client"
)

type AttachmentService struct {
	client *client.Client
}

func NewAttachmentService(c *client.Client) *AttachmentService {
	return &AttachmentService{client: c}
}

type Attachment struct {
	ID       string            `json:"id"`
	Self     string            `json:"self"`
	Filename string            `json:"filename"`
	Author   *User             `json:"author"`
	Created  string            `json:"created"`
	Size     int64             `json:"size"`
	MimeType string            `json:"mimeType"`
	Content  string            `json:"content"`
	Thumbnail string           `json:"thumbnail,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type AttachmentMeta struct {
	Enabled                  bool  `json:"enabled"`
	UploadLimit              int64 `json:"uploadLimit"`
	AllowAttachments         bool  `json:"allowAttachments"`
	AttachmentZipMaxSize     int64 `json:"attachmentZipMaxSize,omitempty"`
	AttachmentZipMaxUnzippedSize int64 `json:"attachmentZipMaxUnzippedSize,omitempty"`
}

func (s *AttachmentService) Get(ctx context.Context, attachmentID string) (*Attachment, error) {
	if err := s.client.CheckEndpoint("attachments"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/rest/api/2/attachment/%s", url.PathEscape(attachmentID))
	var attachment Attachment
	if err := s.client.Get(ctx, path, &attachment); err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (s *AttachmentService) Delete(ctx context.Context, attachmentID string) error {
	if err := s.client.CheckEndpoint("attachments"); err != nil {
		return err
	}
	path := fmt.Sprintf("/rest/api/2/attachment/%s", url.PathEscape(attachmentID))
	return s.client.Delete(ctx, path)
}

func (s *AttachmentService) GetMeta(ctx context.Context) (*AttachmentMeta, error) {
	if err := s.client.CheckEndpoint("attachments"); err != nil {
		return nil, err
	}
	var meta AttachmentMeta
	if err := s.client.Get(ctx, "/rest/api/2/attachment/meta", &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}
