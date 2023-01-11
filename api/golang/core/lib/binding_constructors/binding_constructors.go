package binding_constructors

import (
	"github.com/kurtosis-tech/kurtosis-sdk/api/golang/core/kurtosis_core_rpc_api_bindings"
)

// The generated bindings don't come with constructors (leaving it up to the user to initialize all the fields), so we
// add them so that our code is safer

// ==============================================================================================
//
//	Shared Objects (Used By Multiple Endpoints)
//
// ==============================================================================================

func NewPort(number uint32, protocol kurtosis_core_rpc_api_bindings.Port_TransportProtocol, maybeApplicationProtocol string) *kurtosis_core_rpc_api_bindings.Port {
	return &kurtosis_core_rpc_api_bindings.Port{
		Number:                   number,
		TransportProtocol:        protocol,
		MaybeApplicationProtocol: maybeApplicationProtocol,
	}
}

func NewServiceConfig(
	containerImageName string,
	privatePorts map[string]*kurtosis_core_rpc_api_bindings.Port,
	publicPorts map[string]*kurtosis_core_rpc_api_bindings.Port, //TODO this is a huge hack to temporarily enable static ports for NEAR until we have a more productized solution
	entrypointArgs []string,
	cmdArgs []string,
	envVars map[string]string,
	filesArtifactMountDirpaths map[string]string,
	cpuAllocationMillicpus uint64,
	memoryAllocationMegabytes uint64,
	privateIPAddrPlaceholder string,
	subnetwork string) *kurtosis_core_rpc_api_bindings.ServiceConfig {
	return &kurtosis_core_rpc_api_bindings.ServiceConfig{
		ContainerImageName:        containerImageName,
		PrivatePorts:              privatePorts,
		PublicPorts:               publicPorts,
		EntrypointArgs:            entrypointArgs,
		CmdArgs:                   cmdArgs,
		EnvVars:                   envVars,
		FilesArtifactMountpoints:  filesArtifactMountDirpaths,
		CpuAllocationMillicpus:    cpuAllocationMillicpus,
		MemoryAllocationMegabytes: memoryAllocationMegabytes,
		PrivateIpAddrPlaceholder:  privateIPAddrPlaceholder,
		Subnetwork:                &subnetwork,
	}
}

func NewUpdateServiceConfig(subnetwork string) *kurtosis_core_rpc_api_bindings.UpdateServiceConfig {
	return &kurtosis_core_rpc_api_bindings.UpdateServiceConfig{
		Subnetwork: &subnetwork,
	}
}

// ==============================================================================================
//
//	Execute Starlark Arguments
//
// ==============================================================================================
func NewRunStarlarkScriptArgs(serializedString string, serializedParams string, dryRun bool) *kurtosis_core_rpc_api_bindings.RunStarlarkScriptArgs {
	return &kurtosis_core_rpc_api_bindings.RunStarlarkScriptArgs{
		SerializedScript: serializedString,
		SerializedParams: serializedParams,
		DryRun:           &dryRun,
	}
}

func NewRunStarlarkPackageArgs(packageId string, compressedPackage []byte, serializedParams string, dryRun bool) *kurtosis_core_rpc_api_bindings.RunStarlarkPackageArgs {
	return &kurtosis_core_rpc_api_bindings.RunStarlarkPackageArgs{
		PackageId:              packageId,
		StarlarkPackageContent: &kurtosis_core_rpc_api_bindings.RunStarlarkPackageArgs_Local{Local: compressedPackage},
		SerializedParams:       serializedParams,
		DryRun:                 &dryRun,
	}
}

func NewRunStarlarkRemotePackageArgs(packageId string, serializedParams string, dryRun bool) *kurtosis_core_rpc_api_bindings.RunStarlarkPackageArgs {
	return &kurtosis_core_rpc_api_bindings.RunStarlarkPackageArgs{
		PackageId:              packageId,
		StarlarkPackageContent: &kurtosis_core_rpc_api_bindings.RunStarlarkPackageArgs_Remote{Remote: true},
		SerializedParams:       serializedParams,
		DryRun:                 &dryRun,
	}
}

// ==============================================================================================
//
//	Startosis Execution Response
//
// ==============================================================================================
func NewStarlarkRunResponseLineFromInstruction(instruction *kurtosis_core_rpc_api_bindings.StarlarkInstruction) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_Instruction{
			Instruction: instruction,
		},
	}
}

func NewStarlarkRunResponseLineFromInstructionResult(serializedInstructionResult string) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_InstructionResult{
			InstructionResult: &kurtosis_core_rpc_api_bindings.StarlarkInstructionResult{
				SerializedInstructionResult: serializedInstructionResult,
			},
		},
	}
}

func NewStarlarkRunResponseLineFromInterpretationError(interpretationError *kurtosis_core_rpc_api_bindings.StarlarkInterpretationError) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_Error{
			Error: &kurtosis_core_rpc_api_bindings.StarlarkError{
				Error: &kurtosis_core_rpc_api_bindings.StarlarkError_InterpretationError{
					InterpretationError: interpretationError,
				},
			},
		},
	}
}

func NewStarlarkRunResponseLineFromValidationError(validationError *kurtosis_core_rpc_api_bindings.StarlarkValidationError) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_Error{
			Error: &kurtosis_core_rpc_api_bindings.StarlarkError{
				Error: &kurtosis_core_rpc_api_bindings.StarlarkError_ValidationError{
					ValidationError: validationError,
				},
			},
		},
	}
}

func NewStarlarkRunResponseLineFromExecutionError(executionError *kurtosis_core_rpc_api_bindings.StarlarkExecutionError) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_Error{
			Error: &kurtosis_core_rpc_api_bindings.StarlarkError{
				Error: &kurtosis_core_rpc_api_bindings.StarlarkError_ExecutionError{
					ExecutionError: executionError,
				},
			},
		},
	}
}

func NewStarlarkRunResponseLineFromSinglelineProgressInfo(currentStepInfo string, currentStepNumber uint32, totalSteps uint32) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_ProgressInfo{
			ProgressInfo: &kurtosis_core_rpc_api_bindings.StarlarkRunProgress{
				CurrentStepInfo:   []string{currentStepInfo},
				TotalSteps:        totalSteps,
				CurrentStepNumber: currentStepNumber,
			},
		},
	}
}

func NewStarlarkRunResponseLineFromMultilineProgressInfo(currentStepInfoMultiline []string, currentStepNumber uint32, totalSteps uint32) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_ProgressInfo{
			ProgressInfo: &kurtosis_core_rpc_api_bindings.StarlarkRunProgress{
				CurrentStepInfo:   currentStepInfoMultiline,
				TotalSteps:        totalSteps,
				CurrentStepNumber: currentStepNumber,
			},
		},
	}
}

func NewStarlarkRunResponseLineFromRunFailureEvent() *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_RunFinishedEvent{
			RunFinishedEvent: &kurtosis_core_rpc_api_bindings.StarlarkRunFinishedEvent{
				IsRunSuccessful:  false,
				SerializedOutput: nil,
			},
		},
	}
}

func NewStarlarkRunResponseLineFromRunSuccessEvent(serializedOutputObject string) *kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine {
	return &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine{
		RunResponseLine: &kurtosis_core_rpc_api_bindings.StarlarkRunResponseLine_RunFinishedEvent{
			RunFinishedEvent: &kurtosis_core_rpc_api_bindings.StarlarkRunFinishedEvent{
				IsRunSuccessful:  true,
				SerializedOutput: &serializedOutputObject,
			},
		},
	}
}

func NewStarlarkInstruction(position *kurtosis_core_rpc_api_bindings.StarlarkInstructionPosition, name string, executableInstruction string, arguments []*kurtosis_core_rpc_api_bindings.StarlarkInstructionArg) *kurtosis_core_rpc_api_bindings.StarlarkInstruction {
	return &kurtosis_core_rpc_api_bindings.StarlarkInstruction{
		InstructionName:       name,
		Position:              position,
		ExecutableInstruction: executableInstruction,
		Arguments:             arguments,
	}
}

func NewStarlarkInstructionPosition(filename string, line int32, column int32) *kurtosis_core_rpc_api_bindings.StarlarkInstructionPosition {
	return &kurtosis_core_rpc_api_bindings.StarlarkInstructionPosition{
		Filename: filename,
		Line:     line,
		Column:   column,
	}
}

func NewStarlarkInstructionKwarg(serializedArgValue string, argName string, isRepresentative bool) *kurtosis_core_rpc_api_bindings.StarlarkInstructionArg {
	return &kurtosis_core_rpc_api_bindings.StarlarkInstructionArg{
		SerializedArgValue: serializedArgValue,
		ArgName:            &argName,
		IsRepresentative:   isRepresentative,
	}
}

func NewStarlarkInstructionArg(serializedArgValue string, isRepresentative bool) *kurtosis_core_rpc_api_bindings.StarlarkInstructionArg {
	return &kurtosis_core_rpc_api_bindings.StarlarkInstructionArg{
		SerializedArgValue: serializedArgValue,
		ArgName:            nil,
		IsRepresentative:   isRepresentative,
	}
}

func NewStarlarkInterpretationError(errorMessage string) *kurtosis_core_rpc_api_bindings.StarlarkInterpretationError {
	return &kurtosis_core_rpc_api_bindings.StarlarkInterpretationError{
		ErrorMessage: errorMessage,
	}
}

func NewStarlarkValidationError(errorMessage string) *kurtosis_core_rpc_api_bindings.StarlarkValidationError {
	return &kurtosis_core_rpc_api_bindings.StarlarkValidationError{
		ErrorMessage: errorMessage,
	}
}

func NewStarlarkExecutionError(errorMessage string) *kurtosis_core_rpc_api_bindings.StarlarkExecutionError {
	return &kurtosis_core_rpc_api_bindings.StarlarkExecutionError{
		ErrorMessage: errorMessage,
	}
}

// ==============================================================================================
//
//	Start Service
//
// ==============================================================================================

func NewStartServicesArgs(serviceConfigs map[string]*kurtosis_core_rpc_api_bindings.ServiceConfig) *kurtosis_core_rpc_api_bindings.StartServicesArgs {
	return &kurtosis_core_rpc_api_bindings.StartServicesArgs{
		ServiceIdsToConfigs: serviceConfigs,
	}
}

func NewStartServicesResponse(
	successfulServicesInfo map[string]*kurtosis_core_rpc_api_bindings.ServiceInfo,
	failedServicesErrors map[string]string) *kurtosis_core_rpc_api_bindings.StartServicesResponse {
	return &kurtosis_core_rpc_api_bindings.StartServicesResponse{
		SuccessfulServiceIdsToServiceInfo: successfulServicesInfo,
		FailedServiceIdsToError:           failedServicesErrors,
	}
}

// ==============================================================================================
//
//	Get Service Info
//
// ==============================================================================================

func NewGetServicesArgs(serviceIds map[string]bool) *kurtosis_core_rpc_api_bindings.GetServicesArgs {
	return &kurtosis_core_rpc_api_bindings.GetServicesArgs{
		ServiceIds: serviceIds,
	}
}

func NewGetServicesResponse(
	serviceInfo map[string]*kurtosis_core_rpc_api_bindings.ServiceInfo,
) *kurtosis_core_rpc_api_bindings.GetServicesResponse {
	return &kurtosis_core_rpc_api_bindings.GetServicesResponse{
		ServiceInfo: serviceInfo,
	}
}

func NewServiceInfo(
	guid string,
	privateIpAddr string,
	privatePorts map[string]*kurtosis_core_rpc_api_bindings.Port,
	maybePublicIpAddr string,
	maybePublicPorts map[string]*kurtosis_core_rpc_api_bindings.Port,
) *kurtosis_core_rpc_api_bindings.ServiceInfo {
	return &kurtosis_core_rpc_api_bindings.ServiceInfo{
		ServiceGuid:       guid,
		PrivateIpAddr:     privateIpAddr,
		PrivatePorts:      privatePorts,
		MaybePublicIpAddr: maybePublicIpAddr,
		MaybePublicPorts:  maybePublicPorts,
	}
}

// ==============================================================================================
//
//	Remove Service
//
// ==============================================================================================

func NewRemoveServiceArgs(serviceId string) *kurtosis_core_rpc_api_bindings.RemoveServiceArgs {
	return &kurtosis_core_rpc_api_bindings.RemoveServiceArgs{
		ServiceId: serviceId,
	}
}

func NewRemoveServiceResponse(serviceGuid string) *kurtosis_core_rpc_api_bindings.RemoveServiceResponse {
	return &kurtosis_core_rpc_api_bindings.RemoveServiceResponse{
		ServiceGuid: serviceGuid,
	}
}

// ==============================================================================================
//
//	Repartition
//
// ==============================================================================================

func NewRepartitionArgs(
	partitionServices map[string]*kurtosis_core_rpc_api_bindings.PartitionServices,
	partitionConnections map[string]*kurtosis_core_rpc_api_bindings.PartitionConnections,
	defaultConnection *kurtosis_core_rpc_api_bindings.PartitionConnectionInfo) *kurtosis_core_rpc_api_bindings.RepartitionArgs {
	return &kurtosis_core_rpc_api_bindings.RepartitionArgs{
		PartitionServices:    partitionServices,
		PartitionConnections: partitionConnections,
		DefaultConnection:    defaultConnection,
	}
}

func NewPartitionServices(serviceIdSet map[string]bool) *kurtosis_core_rpc_api_bindings.PartitionServices {
	return &kurtosis_core_rpc_api_bindings.PartitionServices{
		ServiceIdSet: serviceIdSet,
	}
}

func NewPartitionConnections(connectionInfo map[string]*kurtosis_core_rpc_api_bindings.PartitionConnectionInfo) *kurtosis_core_rpc_api_bindings.PartitionConnections {
	return &kurtosis_core_rpc_api_bindings.PartitionConnections{
		ConnectionInfo: connectionInfo,
	}
}

func NewPartitionConnectionInfo(packetLossPercentage float32) *kurtosis_core_rpc_api_bindings.PartitionConnectionInfo {
	return &kurtosis_core_rpc_api_bindings.PartitionConnectionInfo{
		PacketLossPercentage: packetLossPercentage,
	}
}

// ==============================================================================================
//                                          Pause/Unpause Service
// ==============================================================================================

func NewPauseServiceArgs(serviceId string) *kurtosis_core_rpc_api_bindings.PauseServiceArgs {
	return &kurtosis_core_rpc_api_bindings.PauseServiceArgs{
		ServiceId: serviceId,
	}
}

func NewUnpauseServiceArgs(serviceId string) *kurtosis_core_rpc_api_bindings.UnpauseServiceArgs {
	return &kurtosis_core_rpc_api_bindings.UnpauseServiceArgs{
		ServiceId: serviceId,
	}
}

// ==============================================================================================
//
//	Exec Command
//
// ==============================================================================================

func NewExecCommandArgs(serviceId string, commandArgs []string) *kurtosis_core_rpc_api_bindings.ExecCommandArgs {
	return &kurtosis_core_rpc_api_bindings.ExecCommandArgs{
		ServiceId:   serviceId,
		CommandArgs: commandArgs,
	}
}

func NewExecCommandResponse(exitCode int32, logOutput string) *kurtosis_core_rpc_api_bindings.ExecCommandResponse {
	return &kurtosis_core_rpc_api_bindings.ExecCommandResponse{
		ExitCode:  exitCode,
		LogOutput: logOutput,
	}
}

// ==============================================================================================
//
//	Upload Files Artifact
//
// ==============================================================================================

func NewUploadFilesArtifactArgs(data []byte, name string) *kurtosis_core_rpc_api_bindings.UploadFilesArtifactArgs {
	return &kurtosis_core_rpc_api_bindings.UploadFilesArtifactArgs{Data: data, Name: name}
}

// ==============================================================================================
//
//	Store Web Files Artifact
//
// ==============================================================================================

func NewStoreWebFilesArtifactArgs(url string, name string) *kurtosis_core_rpc_api_bindings.StoreWebFilesArtifactArgs {
	return &kurtosis_core_rpc_api_bindings.StoreWebFilesArtifactArgs{Url: url, Name: name}
}

// ==============================================================================================
//
//	Download Files Artifact
//
// ==============================================================================================

func DownloadFilesArtifactArgs(fileIdentifier string) *kurtosis_core_rpc_api_bindings.DownloadFilesArtifactArgs {
	return &kurtosis_core_rpc_api_bindings.DownloadFilesArtifactArgs{
		Identifier: fileIdentifier,
	}
}

// ==============================================================================================
//
//	Render Templates To Files Artifact
//
// ==============================================================================================

func NewTemplateAndData(template string, dataAsJson string) *kurtosis_core_rpc_api_bindings.RenderTemplatesToFilesArtifactArgs_TemplateAndData {
	return &kurtosis_core_rpc_api_bindings.RenderTemplatesToFilesArtifactArgs_TemplateAndData{
		Template:   template,
		DataAsJson: dataAsJson,
	}
}

func NewRenderTemplatesToFilesArtifactResponse(filesArtifactUuid string) *kurtosis_core_rpc_api_bindings.RenderTemplatesToFilesArtifactResponse {
	return &kurtosis_core_rpc_api_bindings.RenderTemplatesToFilesArtifactResponse{
		Uuid: filesArtifactUuid,
	}
}
