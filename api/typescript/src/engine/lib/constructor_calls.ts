import {
    CleanArgs,
    CreateEnclaveArgs,
    DestroyEnclaveArgs,
    GetServiceLogsArgs,
    LogLineFilter,
    LogLineOperator,
    StopEnclaveArgs,
} from "../kurtosis_engine_rpc_api_bindings/engine_service_pb";
import * as jspb from "google-protobuf";
import {ServiceGUID} from "../../core/lib/services/service";
import * as kurtosisCtx from "./kurtosis_context/log_line_filter";
import * as kurtosisLogLineOperator from "./kurtosis_context/log_line_operator";
import {err, ok, Result} from "neverthrow";

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

export function newGetServiceLogsArgs(
        enclaveID: string,
        serviceGUIDs: Set<ServiceGUID>,
        shouldFollowLogs: boolean,
        logLineFilter: kurtosisCtx.LogLineFilter|undefined,
): GetServiceLogsArgs {

    const result: GetServiceLogsArgs = new GetServiceLogsArgs();
    result.setEnclaveId(enclaveID);
    const serviceGUIDSetMap: jspb.Map<string, boolean> = result.getServiceGuidSetMap();
    const isServiceGUIDInSet: boolean = true;
    for (const serviceGUID of serviceGUIDs) {
        serviceGUIDSetMap.set(serviceGUID, isServiceGUIDInSet);
    }
    result.setFollowLogs(shouldFollowLogs)

    let grpcConjunctiveFilters: Array<LogLineFilter>;
    try {
        grpcConjunctiveFilters = newGRPCConjunctiveFilters(logLineFilter)
    } catch (error) {
        throw new Error(`An error occurred creating the GRPC conjunctive log line filters '${logLineFilter}'. Error:\n${error}`);
    }

    result.setConjunctiveFiltersList(grpcConjunctiveFilters)
    return result;
}

// ====================================================================================================
//                                       Private helper functions
// ====================================================================================================
//Even though the backend is prepared for receiving a list of conjunctive filters
//We allow users to send only one filter so far, because it covers the current supported use cases
function newGRPCConjunctiveFilters(logLineFilter: kurtosisCtx.LogLineFilter|undefined): Array<LogLineFilter> {
    const grpcLogLineFilters:Array<LogLineFilter> = new Array<LogLineFilter>();

    if(logLineFilter === undefined){
        return grpcLogLineFilters
    }

    const grpcLogLineFilter:LogLineFilter = new LogLineFilter();
    grpcLogLineFilter.setTextPattern(logLineFilter.getTextPattern())
    switch (logLineFilter.getOperator()) {
        case kurtosisLogLineOperator.LogLineOperator.DoesContainText:
            grpcLogLineFilter.setOperator(LogLineOperator.LOGLINEOPERATOR_DOES_CONTAIN_TEXT)
            break;
        case kurtosisLogLineOperator.LogLineOperator.DoesNotContainText:
            grpcLogLineFilter.setOperator(LogLineOperator.LOGLINEOPERATOR_DOES_NOT_CONTAIN_TEXT)
            break;
        case kurtosisLogLineOperator.LogLineOperator.DoesContainMatchRegex:
            grpcLogLineFilter.setOperator(LogLineOperator.LOGLINEOPERATOR_DOES_CONTAIN_MATCH_REGEX)
            break;
        case kurtosisLogLineOperator.LogLineOperator.DoesNotContainMatchRegex:
            grpcLogLineFilter.setOperator(LogLineOperator.LOGLINEOPERATOR_DOES_NOT_CONTAIN_MATCH_REGEX)
            break;
        default:
            throw new Error(`Unrecognized log line filter operator '${logLineFilter.getOperator()}' in filter '${logLineFilter}'; this is a bug in Kurtosis`);
            break;
    }
    grpcLogLineFilters.push(grpcLogLineFilter)

    return grpcLogLineFilters
}
