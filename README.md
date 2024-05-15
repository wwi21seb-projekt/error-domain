# Error Codes

This is the shared error domain for the project. If you want to contribute, you need to add the error to `errors.json` and commit it using conventional commits.

| Title | Code | Message | HTTP Status |
| --- | --- | --- | --- |
| BadRequest | ERR-001 | The request body is invalid. Please check the request body and try again. | 400 |
| UsernameTaken | ERR-002 | The username is already taken. Please try another username. | 409 |
| EmailTaken | ERR-003 | The email is already taken. Please try another email. | 409 |
| UserNotFound | ERR-004 | The user was not found. Please check the username and try again. | 404 |
| UserNotActivated | ERR-005 | The user is not activated. Please activate the user and try again. | 403 |
| ActivationTokenExpired | ERR-006 | The token has expired. Please check your mail for a new token and try again. | 401 |
| InvalidToken | ERR-007 | The token is invalid. Please check the token and try again. | 401 |
| InvalidCredentials | ERR-008 | The credentials are invalid. Please check the credentials and try again. | 404 |
| InternalServerError | ERR-009 | An internal server error occurred. Please try again later. | 500 |
| DatabaseError | ERR-010 | A database error occurred. Please try again later. | 500 |
| EmailUnreachable | ERR-011 | The email is unreachable. Please check the email and try again. | 400 |
| EmailNotSent | ERR-012 | The email could not be sent. Please try again later. | 500 |
| UserAlreadyActivated | ERR-013 | The user is already activated. Please login to your account. | 400 |
| Unauthorized | ERR-014 | The request is unauthorized. Please login to your account. | 401 |
| SubscriptionNotFound | ERR-015 | The subscription was not found. Please check the username and try again. | 404 |
| SubscriptionAlreadyExists | ERR-016 | The subscription already exists. Please check the username and try again. | 409 |
| SubscriptionSelfFollow | ERR-017 | You cannot follow yourself. Please check the username and try again. | 406 |
| UnsubscribeForbidden | ERR-018 | You can only delete your own subscriptions. | 403 |
| DeletePostForbidden | ERR-019 | You can only delete your own posts. | 403 |
| PostNotFound | ERR-020 | The post was not found. Please check the post ID and try again. | 404 |
| AlreadyLiked | ERR-021 | You have already liked this post. | 409 |
| NotLiked | ERR-022 | You can't unlike a post you haven't liked. | 409 |
| NotificationNotFound | ERR-023 | The notification was not found. Please check the notification ID and try again. | 404 |
| DeleteNotificationForbidden | ERR-024 | You can only delete your own notifications. | 403 |
