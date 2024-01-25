package models

type Workflow struct {
	WorkFlowId       string           `json:"workFlowId"`
	WorkFlowName     string           `json:"workFlowName"`
	WorkFlowCategory WorkFlowCategory `json:"workFlowCategory"`
	WorkFlowTask     []WorkflowTask   `json:"workFlowTasks"`
	WorkFlowOwner    WorkFlowOwner    `json:"workFlowOwner"`
	WorkFlowManager  WorkFlowManager  `json:"workFlowManagedGroup"`
}

type WorkflowTask struct {
	WorkflowTaskId          string `json:"workflowTaskId"`
	WorkFlowTaskName        string `json:"workflowTaskName"`
	WorkflowTaskDescription string `json:"workflowTaskDescription"`
}

type WorkFlowCategory struct {
	WorkFlowCategoryId   string `json:"workFlowCategoryId"`
	WorkFlowCategoryName string `json:"workFlowCategoryName"`
}

type WorkFlowOwner struct {
	WorkFlowOwnerId    string `json:"workFlowOwnerId"`
	WorkFlowOwnerUrl   string `json:"workFlowOwnerUrl"`
	WorkFlowOwnerEmail string `json:"workFlowOwnerEmail"`
	WorkFlowOwnerPhone string `json:"workFlowOwnerPhone"`
}

type WorkFlowManager struct {
	WorkFlowManagerId    string `json:"workFlowManagerId"`
	WorkFlowManagerUrl   string `json:"workFlowManagerUrl"`
	WorkFlowManagerEmail string `json:"workFlowManagerEmail"`
	WorkFlowManagerPhone string `json:"workFlowManagerPhone"`
}
