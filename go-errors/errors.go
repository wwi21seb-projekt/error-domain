package goerrors

type CustomError struct {
	Message string
	Code string
	HttpStatus int
}

var (
	BadRequest = &CustomError{
		Message: "The request body is invalid. Please check the request body and try again.",
		Code: "BadRequest",
		HttpStatus: 400,
	}
	UsernameTaken = &CustomError{
		Message: "The username is already taken. Please try another username.",
		Code: "UsernameTaken",
		HttpStatus: 409,
	}
	EmailTaken = &CustomError{
		Message: "The email is already taken. Please try another email.",
		Code: "EmailTaken",
		HttpStatus: 409,
	}
	UserNotFound = &CustomError{
		Message: "The user was not found. Please check the username and try again.",
		Code: "UserNotFound",
		HttpStatus: 404,
	}
	UserNotActivated = &CustomError{
		Message: "The user is not activated. Please activate the user and try again.",
		Code: "UserNotActivated",
		HttpStatus: 403,
	}
	ActivationTokenExpired = &CustomError{
		Message: "The token has expired. Please check your mail for a new token and try again.",
		Code: "ActivationTokenExpired",
		HttpStatus: 401,
	}
	InvalidToken = &CustomError{
		Message: "The token is invalid. Please check the token and try again.",
		Code: "InvalidToken",
		HttpStatus: 401,
	}
	InvalidCredentials = &CustomError{
		Message: "The credentials are invalid. Please check the credentials and try again.",
		Code: "InvalidCredentials",
		HttpStatus: 401,
	}
	InternalServerError = &CustomError{
		Message: "An internal server error occurred. Please try again later.",
		Code: "InternalServerError",
		HttpStatus: 500,
	}
	DatabaseError = &CustomError{
		Message: "A database error occurred. Please try again later.",
		Code: "DatabaseError",
		HttpStatus: 500,
	}
	EmailUnreachable = &CustomError{
		Message: "The email is unreachable. Please check the email and try again.",
		Code: "EmailUnreachable",
		HttpStatus: 400,
	}
	EmailNotSent = &CustomError{
		Message: "The email could not be sent. Please try again later.",
		Code: "EmailNotSent",
		HttpStatus: 500,
	}
	UserAlreadyActivated = &CustomError{
		Message: "The user is already activated. Please login to your account.",
		Code: "UserAlreadyActivated",
		HttpStatus: 400,
	}
	Unauthorized = &CustomError{
		Message: "The request is unauthorized. Please login to your account.",
		Code: "Unauthorized",
		HttpStatus: 401,
	}
	SubscriptionNotFound = &CustomError{
		Message: "The subscription was not found. Please check the username and try again.",
		Code: "SubscriptionNotFound",
		HttpStatus: 404,
	}
	SubscriptionAlreadyExists = &CustomError{
		Message: "The subscription already exists. Please check the username and try again.",
		Code: "SubscriptionAlreadyExists",
		HttpStatus: 409,
	}
	SubscriptionSelfFollow = &CustomError{
		Message: "You cannot follow yourself. Please check the username and try again.",
		Code: "SubscriptionSelfFollow",
		HttpStatus: 400,
	}
	UnsubscribeForbidden = &CustomError{
		Message: "You can only delete your own subscriptions.",
		Code: "UnsubscribeForbidden",
		HttpStatus: 403,
	}
	DeletePostForbidden = &CustomError{
		Message: "You can only delete your own posts.",
		Code: "DeletePostForbidden",
		HttpStatus: 403,
	}
	PostNotFound = &CustomError{
		Message: "The post was not found. Please check the post ID and try again.",
		Code: "PostNotFound",
		HttpStatus: 404,
	}
)
