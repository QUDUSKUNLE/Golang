package templates
// Create send registration email template
// This template is used to send a registration email to the user.
// It includes a link to verify the user's email address and a button to complete the registration process.
// 	return cfg, ctx, stop
// }

func RegistrationEmailTemplate(name, email, token string) string {
	return `
	<!DOCTYPE html>
	<html lang="en">
	<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
		</head>
		<body>
			<p>Hello, ` + name + `!</p>
			<p>Please verify your email address by clicking the link below:</p>
			<a href="https://s3records.com/verify?token=` + token + `">Verify Email</a>
		</body>
	</html>
	`
}

// This function generates a registration email template with the user's name, email, and token.

