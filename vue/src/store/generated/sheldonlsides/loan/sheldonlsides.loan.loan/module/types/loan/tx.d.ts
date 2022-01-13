import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "sheldonlsides.loan.loan";
export interface MsgRequestLoan {
    creator: string;
    amount: string;
    fee: string;
    collateral: string;
    deadline: string;
}
export interface MsgRequestLoanResponse {
}
export interface MsgApproveLoan {
    creator: string;
    id: number;
}
export interface MsgApproveLoanResponse {
}
export declare const MsgRequestLoan: {
    encode(message: MsgRequestLoan, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRequestLoan;
    fromJSON(object: any): MsgRequestLoan;
    toJSON(message: MsgRequestLoan): unknown;
    fromPartial(object: DeepPartial<MsgRequestLoan>): MsgRequestLoan;
};
export declare const MsgRequestLoanResponse: {
    encode(_: MsgRequestLoanResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRequestLoanResponse;
    fromJSON(_: any): MsgRequestLoanResponse;
    toJSON(_: MsgRequestLoanResponse): unknown;
    fromPartial(_: DeepPartial<MsgRequestLoanResponse>): MsgRequestLoanResponse;
};
export declare const MsgApproveLoan: {
    encode(message: MsgApproveLoan, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveLoan;
    fromJSON(object: any): MsgApproveLoan;
    toJSON(message: MsgApproveLoan): unknown;
    fromPartial(object: DeepPartial<MsgApproveLoan>): MsgApproveLoan;
};
export declare const MsgApproveLoanResponse: {
    encode(_: MsgApproveLoanResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveLoanResponse;
    fromJSON(_: any): MsgApproveLoanResponse;
    toJSON(_: MsgApproveLoanResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveLoanResponse>): MsgApproveLoanResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    ApproveLoan(request: MsgApproveLoan): Promise<MsgApproveLoanResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
    ApproveLoan(request: MsgApproveLoan): Promise<MsgApproveLoanResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
