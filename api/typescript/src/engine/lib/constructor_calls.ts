import {
    CleanArgs,
    CreateEnclaveArgs,
    DestroyEnclaveArgs,
    GetUserServiceLogsArgs,
    StopEnclaveArgs
} from "../kurtosis_engine_rpc_api_bindings/engine_service_pb";
import * as jspb from "google-protobuf";
import {ServiceGUID} from "../../core/lib/services/service";

// ====================================================================================================
//                                    Kurtosis Context
// ====================================================================================================


export function newCreateEnclaveArgs(
        enclaveId: string,
        apiContainerImageVersionTag: string,
        apiContainerLogLevel: string,
        isPartitioningEnabled: boolean): CreateEnclaveArgs {
    const result: CreateEnclaveArgs = new CreateEnclaveArgs();
    result.setEnclaveId(enclaveId);
    result.setApiContainerVersionTag(apiContainerImageVersionTag);
    result.setApiContainerLogLevel(apiContainerLogLevel);
    result.setIsPartitioningEnabled(isPartitioningEnabled);

    return result;
}

export function newStopEnclaveArgs(enclaveId:string): DestroyEnclaveArgs {
    const result: StopEnclaveArgs = new StopEnclaveArgs();
    result.setEnclaveId(enclaveId);
    return result;
}

export function newDestroyEnclaveArgs(enclaveId:string): DestroyEnclaveArgs {
    const result: DestroyEnclaveArgs = new DestroyEnclaveArgs();
    result.setEnclaveId(enclaveId);
    return result;
}

export function newCleanArgs(shouldCleanAll:boolean): CleanArgs {
    const result: CleanArgs = new CleanArgs();
    result.setShouldCleanAll(shouldCleanAll);
    return result;
}

export function newGetUserServiceLogsArgs(
        enclaveID: string,
        userServiceGUIDs: Set<ServiceGUID>,
        shouldFollowLogs: boolean,
): GetUserServiceLogsArgs {

    const result: GetUserServiceLogsArgs = new GetUserServiceLogsArgs();
    result.setEnclaveId(enclaveID);
    const serviceGUIDSetMap: jspb.Map<string, boolean> = result.getServiceGuidSetMap();
    const isServiceGUIDInSet: boolean = true;
    for (const serviceGUID of userServiceGUIDs) {
        serviceGUIDSetMap.set(serviceGUID, isServiceGUIDInSet);
    }
    result.setFollowLogs(shouldFollowLogs)
    return result;
}
