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
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
