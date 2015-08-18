package bbs

import "github.com/tedsuo/rata"

const (
	// Domains
	DomainsRoute      = "Domains"
	UpsertDomainRoute = "UpsertDomain"

	// Actual LRPs
	ActualLRPGroupsRoute                     = "ActualLRPGroups"
	ActualLRPGroupsByProcessGuidRoute        = "ActualLRPGroupsByProcessGuid"
	ActualLRPGroupByProcessGuidAndIndexRoute = "ActualLRPGroupsByProcessGuidAndIndex"

	// Actual LRP Lifecycle
	ClaimActualLRPRoute  = "ClaimActualLRP"
	StartActualLRPRoute  = "StartActualLRP"
	CrashActualLRPRoute  = "CrashActualLRP"
	FailActualLRPRoute   = "FailActualLRP"
	RemoveActualLRPRoute = "RemoveActualLRP"
	RetireActualLRPRoute = "RetireActualLRP"

	// Evacuation
	RemoveEvacuatingActualLRPRoute = "RemoveEvacuatingActualLRP"
	EvacuateClaimedActualLRPRoute  = "EvacuateClaimedActualLRP"
	EvacuateCrashedActualLRPRoute  = "EvacuateCrashedActualLRP"
	EvacuateStoppedActualLRPRoute  = "EvacuateStoppedActualLRP"
	EvacuateRunningActualLRPRoute  = "EvacuateRunningActualLRP"

	// Desired LRPs
	DesiredLRPsRoute             = "DesiredLRPs"
	DesiredLRPByProcessGuidRoute = "DesiredLRPByProcessGuid"

	// Tasks
	TasksRoute         = "Tasks"
	TaskByGuidRoute    = "TaskByGuid"
	DesireTaskRoute    = "DesireTask"
	StartTaskRoute     = "StartTask"
	CancelTaskRoute    = "CancelTask"
	FailTaskRoute      = "FailTask"
	CompleteTaskRoute  = "CompleteTask"
	ResolvingTaskRoute = "ResolvingTask"
	ResolveTaskRoute   = "ResolveTask"
	ConvergeTasksRoute = "ConvergeTasks"

	// Event Streaming
	EventStreamRoute = "EventStream"
)

var Routes = rata.Routes{
	// Domains
	{Path: "/v1/domains/list", Method: "POST", Name: DomainsRoute},
	{Path: "/v1/domains/upsert", Method: "POST", Name: UpsertDomainRoute},

	// Actual LRPs
	{Path: "/v1/actual_lrp_groups", Method: "GET", Name: ActualLRPGroupsRoute},
	{Path: "/v1/actual_lrp_groups/:process_guid", Method: "GET", Name: ActualLRPGroupsByProcessGuidRoute},
	{Path: "/v1/actual_lrp_groups/:process_guid/index/:index", Method: "GET", Name: ActualLRPGroupByProcessGuidAndIndexRoute},

	// Actual LRP Lifecycle
	{Path: "/v1/actual_lrps/claim", Method: "POST", Name: ClaimActualLRPRoute},
	{Path: "/v1/actual_lrps/start", Method: "POST", Name: StartActualLRPRoute},
	{Path: "/v1/actual_lrps/crash", Method: "POST", Name: CrashActualLRPRoute},
	{Path: "/v1/actual_lrps/fail", Method: "POST", Name: FailActualLRPRoute},
	{Path: "/v1/actual_lrps/:process_guid/index/:index", Method: "DELETE", Name: RemoveActualLRPRoute},
	{Path: "/v1/actual_lrps/retire", Method: "POST", Name: RetireActualLRPRoute},

	// Evacuation
	{Path: "/v1/evacuating_actual_lrps/remove", Method: "POST", Name: RemoveEvacuatingActualLRPRoute},
	{Path: "/v1/actual_lrps/evacuate_claimed", Method: "POST", Name: EvacuateClaimedActualLRPRoute},
	{Path: "/v1/actual_lrps/evacuate_crashed", Method: "POST", Name: EvacuateCrashedActualLRPRoute},
	{Path: "/v1/actual_lrps/evacuate_stopped", Method: "POST", Name: EvacuateStoppedActualLRPRoute},
	{Path: "/v1/actual_lrps/evacuate_running", Method: "POST", Name: EvacuateRunningActualLRPRoute},

	// Desired LRPs
	{Path: "/v1/desired_lrps", Method: "GET", Name: DesiredLRPsRoute},
	{Path: "/v1/desired_lrps/:process_guid", Method: "GET", Name: DesiredLRPByProcessGuidRoute},

	// Tasks
	{Path: "/v1/tasks", Method: "POST", Name: DesireTaskRoute},
	{Path: "/v1/tasks", Method: "GET", Name: TasksRoute},
	{Path: "/v1/tasks/:task_guid", Method: "GET", Name: TaskByGuidRoute},

	// Task Lifecycle
	{Path: "/v1/tasks/start", Method: "POST", Name: StartTaskRoute},
	{Path: "/v1/tasks/cancel", Method: "POST", Name: CancelTaskRoute},
	{Path: "/v1/tasks/fail", Method: "POST", Name: FailTaskRoute},
	{Path: "/v1/tasks/complete", Method: "POST", Name: CompleteTaskRoute},
	{Path: "/v1/tasks/resolving", Method: "POST", Name: ResolvingTaskRoute},
	{Path: "/v1/tasks/resolve", Method: "POST", Name: ResolveTaskRoute},

	// Task Convergence
	{Path: "/v1/tasks/converge", Method: "POST", Name: ConvergeTasksRoute},

	// Event Streaming
	{Path: "/v1/events", Method: "GET", Name: EventStreamRoute},
}
