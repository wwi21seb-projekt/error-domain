/** GENERATED FILE - DO NOT MODIFY DIRECTLY **/
export declare class CustomError extends Error {
    code: string;
    httpStatus: number;
    constructor(message: string, code: string, httpStatus: number);
}
export declare const BadRequest: CustomError;
export declare const UsernameTaken: CustomError;
export declare const EmailTaken: CustomError;
export declare const UserNotFound: CustomError;
export declare const UserNotActivated: CustomError;
export declare const ActivationTokenExpired: CustomError;
export declare const InvalidToken: CustomError;
export declare const InvalidCredentials: CustomError;
export declare const InternalServerError: CustomError;
export declare const DatabaseError: CustomError;
export declare const EmailUnreachable: CustomError;
export declare const EmailNotSent: CustomError;
export declare const UserAlreadyActivated: CustomError;
export declare const Unauthorized: CustomError;
export declare const SubscriptionNotFound: CustomError;
export declare const SubscriptionAlreadyExists: CustomError;
export declare const SubscriptionSelfFollow: CustomError;
export declare const UnsubscribeForbidden: CustomError;
export declare const DeletePostForbidden: CustomError;
export declare const PostNotFound: CustomError;
//# sourceMappingURL=errors.d.ts.map