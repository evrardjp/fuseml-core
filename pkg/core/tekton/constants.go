package tekton

const (
	pipelineRunPrefix         = "fuseml-"
	pipelineRunServiceAccount = "staging-triggers-admin"
	workspaceAccessMode       = "ReadWriteOnce"
	workspaceSize             = "2Gi"
	inputTypeCodeset          = "codeset"
	codesetWorkspaceName      = "source"
	builderTaskName           = "kaniko"
	cloneTaskName             = "clone"
	codesetNameParam          = "codeset-name"
	codesetVersionParam       = "codeset-version"
	codesetURLParam           = "codeset-url"
	fuseMLRegistry            = "registry.fuseml-registry"
	fuseMLRegistryLocal       = "127.0.0.1:30500"
	imageParamName            = "IMAGE"
	stepOutputVarName         = "TASK_RESULT"
	inputsVarPrefix           = "FUSEML_"
	stepDefaultCmd            = "run"
)
