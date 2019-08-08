# Mailer

## TODO
- limit line length to 78 or 998 characters
	- be careful with 78 limit when sending token, could make trouble when copy/pasting
		- tokens are about 900 characters long
- make sure CRLF (\r\n) is used in body
	- maybe do a replace?

- Impl support for multiple recipients.
- // BCC support
- make logger injectable and do not log by default, thus maybe use a dummy logger
