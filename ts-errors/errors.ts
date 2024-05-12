export class CustomError extends Error {
    code: string;
    httpStatus: number;

    constructor(message: string, code: string, httpStatus: number) {
        super(message);
        this.code = code;
        this.httpStatus = httpStatus;
    }
}

export const BadRequest = new CustomError(
	"The request body is invalid. Please check the request body and try again.", 
	"BadRequest", 
	400
);
export const UsernameTaken = new CustomError(
	"The username is already taken. Please try another username.", 
	"UsernameTaken", 
	409
);
export const EmailTaken = new CustomError(
	"The email is already taken. Please try another email.", 
	"EmailTaken", 
	409
);
export const UserNotFound = new CustomError(
	"The user was not found. Please check the username and try again.", 
	"UserNotFound", 
	404
);
export const UserNotActivated = new CustomError(
	"The user is not activated. Please activate the user and try again.", 
	"UserNotActivated", 
	403
);
export const ActivationTokenExpired = new CustomError(
	"The token has expired. Please check your mail for a new token and try again.", 
	"ActivationTokenExpired", 
	401
);
export const InvalidToken = new CustomError(
	"The token is invalid. Please check the token and try again.", 
	"InvalidToken", 
	401
);
export const InvalidCredentials = new CustomError(
	"The credentials are invalid. Please check the credentials and try again.", 
	"InvalidCredentials", 
	401
);
export const InternalServerError = new CustomError(
	"An internal server error occurred. Please try again later.", 
	"InternalServerError", 
	500
);
export const DatabaseError = new CustomError(
	"A database error occurred. Please try again later.", 
	"DatabaseError", 
	500
);
export const EmailUnreachable = new CustomError(
	"The email is unreachable. Please check the email and try again.", 
	"EmailUnreachable", 
	400
);
export const EmailNotSent = new CustomError(
	"The email could not be sent. Please try again later.", 
	"EmailNotSent", 
	500
);
export const UserAlreadyActivated = new CustomError(
	"The user is already activated. Please login to your account.", 
	"UserAlreadyActivated", 
	400
);
export const Unauthorized = new CustomError(
	"The request is unauthorized. Please login to your account.", 
	"Unauthorized", 
	401
);
export const SubscriptionNotFound = new CustomError(
	"The subscription was not found. Please check the username and try again.", 
	"SubscriptionNotFound", 
	404
);
export const SubscriptionAlreadyExists = new CustomError(
	"The subscription already exists. Please check the username and try again.", 
	"SubscriptionAlreadyExists", 
	409
);
export const SubscriptionSelfFollow = new CustomError(
	"You cannot follow yourself. Please check the username and try again.", 
	"SubscriptionSelfFollow", 
	400
);
export const UnsubscribeForbidden = new CustomError(
	"You can only delete your own subscriptions.", 
	"UnsubscribeForbidden", 
	403
);
export const DeletePostForbidden = new CustomError(
	"You can only delete your own posts.", 
	"DeletePostForbidden", 
	403
);
export const PostNotFound = new CustomError(
	"The post was not found. Please check the post ID and try again.", 
	"PostNotFound", 
	404
);
