package api

import (
	"github.com/darthkoax/jiracli/internal/client"
)

type Services struct {
	Issues                  *IssueService
	Projects                *ProjectService
	Users                   *UserService
	Groups                  *GroupService
	Search                  *SearchService
	Filters                 *FilterService
	Dashboards              *DashboardService
	Boards                  *BoardService
	Sprints                 *SprintService
	Fields                  *FieldService
	Components              *ComponentService
	Versions                *VersionService
	Priorities              *PriorityService
	Statuses                *StatusService
	Resolutions             *ResolutionService
	IssueTypes              *IssueTypeService
	Permissions             *PermissionService
	Myself                  *MyselfService
	Screens                 *ScreenService
	Reindex                 *ReindexService
	ApplicationProperties   *ApplicationPropertiesService
	Configuration           *ConfigurationService
	Avatars                 *AvatarService
	Attachments             *AttachmentService
	Comments                *StandaloneCommentService
	Changelogs              *ChangelogService
	Notify                  *NotifyService
	IssueLinks              *IssueLinkService
	Epics                   *EpicService
	Backlog                 *BacklogService
	HealthCheck             *HealthCheckService
	Webhooks                *WebhookService
}

func NewServices(c *client.Client) *Services {
	return &Services{
		Issues:                NewIssueService(c),
		Projects:              NewProjectService(c),
		Users:                 NewUserService(c),
		Groups:                NewGroupService(c),
		Search:                NewSearchService(c),
		Filters:               NewFilterService(c),
		Dashboards:            NewDashboardService(c),
		Boards:                NewBoardService(c),
		Sprints:               NewSprintService(c),
		Fields:                NewFieldService(c),
		Components:            NewComponentService(c),
		Versions:              NewVersionService(c),
		Priorities:            NewPriorityService(c),
		Statuses:              NewStatusService(c),
		Resolutions:           NewResolutionService(c),
		IssueTypes:            NewIssueTypeService(c),
		Permissions:           NewPermissionService(c),
		Myself:                NewMyselfService(c),
		Screens:               NewScreenService(c),
		Reindex:               NewReindexService(c),
		ApplicationProperties: NewApplicationPropertiesService(c),
		Configuration:         NewConfigurationService(c),
		Avatars:               NewAvatarService(c),
		Attachments:           NewAttachmentService(c),
		Comments:              NewStandaloneCommentService(c),
		Changelogs:            NewChangelogService(c),
		Notify:                NewNotifyService(c),
		IssueLinks:            NewIssueLinkService(c),
		Epics:                 NewEpicService(c),
		Backlog:               NewBacklogService(c),
		HealthCheck:           NewHealthCheckService(c),
		Webhooks:              NewWebhookService(c),
	}
}
