# Jira CLI - Glossary of Terms

This glossary defines all Jira terms and concepts covered by the `jiracli` tool.

## Core Concepts

### Issue
A unit of work in Jira. Issues can be bugs, tasks, stories, epics, or any other type of work item. Each issue has a unique key (e.g., PROJ-123) and contains fields like summary, description, status, assignee, etc.

### Project
A container for organizing related issues. Projects have a unique key (e.g., PROJ) and contain components, versions, and issues. Projects can be software projects, business projects, or service desk projects.

### Component
A sub-section of a project. Components are used to group issues within a project by functional area (e.g., "Backend", "Frontend", "Database"). Each issue can have multiple components.

### Version
A point in time in a project's lifecycle. Versions represent releases (e.g., "1.0", "2.0", "Sprint 1"). Issues can be associated with affected versions and fix versions.

### Field
A piece of data on an issue. Fields include standard fields (summary, description, status) and custom fields (customfield_10001). Fields can be text, numbers, dates, users, or complex objects.

### Screen
A configuration that determines which fields are displayed when creating, editing, or viewing an issue. Screens contain tabs, and each tab contains fields.

### Workflow
The lifecycle of an issue. Workflows define the statuses an issue can be in and the transitions between them. For example: Open → In Progress → Done.

### Status
The current state of an issue in its workflow. Common statuses include Open, In Progress, Done, Closed, etc.

### Resolution
The reason why an issue was closed. Common resolutions include Fixed, Won't Fix, Duplicate, Cannot Reproduce, etc.

### Priority
The importance or urgency of an issue. Common priorities include Highest, High, Medium, Low, Lowest.

### Issue Type
The category of an issue. Common issue types include Bug, Story, Task, Epic, Sub-task, etc.

## People & Permissions

### User
A person who can access Jira. Users have usernames, email addresses, and can be assigned to issues, added as watchers, or mentioned in comments.

### Group
A collection of users. Groups are used to manage permissions and notifications. Users can belong to multiple groups.

### Assignee
The user currently responsible for working on an issue. An issue can have one assignee or be unassigned.

### Reporter
The user who created the issue. The reporter is typically notified of changes to the issue.

### Watcher
A user who receives notifications about an issue. Watchers are notified of all changes but are not necessarily working on the issue.

### Voter
A user who has voted for an issue. Votes are used to indicate the importance or demand for an issue.

### Administrator (Admin)
A user with elevated permissions. Admin mode in jiracli enables access to administrative operations like managing users, groups, permissions, and system configuration.

## Agile/Scrum Concepts

### Board
An agile board that displays issues in columns representing workflow statuses. Boards can be Scrum boards (with sprints) or Kanban boards (continuous flow).

### Sprint
A time-boxed period (usually 1-4 weeks) during which a set of issues are completed. Sprints have a start date, end date, and goal. Issues are moved into a sprint from the backlog.

### Epic
A large user story or feature that spans multiple sprints. Epics are broken down into smaller stories and tasks.

### Backlog
The list of issues that are not yet assigned to a sprint. The backlog is prioritized and issues are pulled into sprints during sprint planning.

### Sprint Goal
A short description of what the sprint will deliver. The goal helps the team focus on the outcome rather than just completing tasks.

## Communication & Tracking

### Comment
A note added to an issue. Comments are used for communication, updates, and discussions about an issue. Comments can be edited and deleted.

### Worklog
A record of time spent working on an issue. Worklogs track the time spent, the date, and a description of the work done.

### Attachment
A file attached to an issue. Attachments can be documents, images, logs, or any other file type. (Note: Attachment upload is not yet implemented in jiracli)

### Transition
A change in an issue's status. Transitions move an issue from one status to another (e.g., from "Open" to "In Progress"). Some transitions require additional fields like resolution.

### Changelog
A history of all changes made to an issue. The changelog tracks field changes, status transitions, comments, and worklogs.

## Search & Organization

### JQL (Jira Query Language)
A powerful query language used to search for issues. JQL supports operators like =, !=, IN, NOT IN, ~, and functions like currentUser(), now(), etc.

### Filter
A saved JQL search query. Filters can be shared with other users or groups and used in dashboards and boards.

### Dashboard
A customizable page that displays gadgets showing issue data, charts, and other information. Dashboards are used to monitor project health and progress.

### Gadget
A widget on a dashboard that displays specific information. Gadgets can show issue lists, charts, calendars, and other data visualizations.

## System & Configuration

### Permission Scheme
A set of permissions that define what users can do in a project. Permission schemes control access to operations like browsing projects, creating issues, editing issues, etc.

### Permission
A specific right to perform an operation. Examples include BROWSE_PROJECTS, CREATE_ISSUES, EDIT_ISSUES, DELETE_ISSUES, etc.

### Avatar
An image representing a user, project, or issue type. Avatars can be system-provided or custom-uploaded.

### Application Property
A system-wide configuration setting. Application properties control Jira's behavior and can only be modified by administrators.

### Reindex
The process of rebuilding Jira's search index. Reindexing is necessary after certain configuration changes or when search results are inconsistent.

### Webhook
A mechanism for sending real-time notifications when events occur in Jira. Webhooks can notify external systems when issues are created, updated, or deleted.

## Relationships

### Issue Link
A connection between two issues. Issue links define relationships like "blocks", "is blocked by", "relates to", "duplicates", etc.

### Sub-task
A child issue of a parent issue. Sub-tasks are used to break down a parent issue into smaller work items. Sub-tasks are typically used with Stories or Tasks.

### Linked Issue
An issue that is connected to another issue through an issue link. Linked issues can be in the same project or different projects.

## Common Abbreviations

- **JQL**: Jira Query Language
- **API**: Application Programming Interface
- **REST**: Representational State Transfer (type of web API)
- **DC**: Data Center (Jira Data Center deployment)
- **TOML**: Tom's Obvious, Minimal Language (configuration file format)
- **TLS**: Transport Layer Security (encryption protocol)
- **CA**: Certificate Authority
- **JIRA**: Just a Rather Intelligent Acronym (original name, now just "Jira")

## Command-Specific Terms

### Issue Key
The unique identifier for an issue, consisting of the project key and a number (e.g., PROJ-123). Issue keys are case-insensitive.

### Issue ID
The internal numeric identifier for an issue. Issue IDs are unique across all projects but are less user-friendly than issue keys.

### Project Key
The short code identifying a project (e.g., PROJ). Project keys are uppercase and used in issue keys.

### Transition ID
The numeric identifier for a workflow transition. Transition IDs are used with the `issue transition` command.

### Comment ID
The unique identifier for a comment. Comment IDs are used with the `issue comment update` and `issue comment delete` commands.

### Worklog ID
The unique identifier for a worklog entry. Worklog IDs are used with the `issue worklog update` and `issue worklog delete` commands.

### Board ID
The numeric identifier for an agile board. Board IDs are used with board-related commands.

### Sprint ID
The numeric identifier for a sprint. Sprint IDs are used with sprint-related commands.

### Filter ID
The unique identifier for a saved filter. Filter IDs are used with filter-related commands.

### Dashboard ID
The unique identifier for a dashboard. Dashboard IDs are used with dashboard-related commands.

### Screen ID
The numeric identifier for a screen. Screen IDs are used with screen-related commands.

### Component ID
The unique identifier for a component. Component IDs are used with component-related commands.

### Version ID
The unique identifier for a version. Version IDs are used with version-related commands.

### Field ID
The identifier for a field. Standard fields use names like "summary", "description", etc. Custom fields use IDs like "customfield_10001".

### Avatar ID
The numeric identifier for an avatar. Avatar IDs are used with avatar-related commands.

### Permission Scheme ID
The numeric identifier for a permission scheme. Permission scheme IDs are used with permission-related commands.
