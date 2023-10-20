## Basic Cloudlfare bypass

I've had trouble with Go in the past making requests that do basic
TLS cipher checks so I figured out a basic solution for this situation.
This mainly just bypasses the "Just a Moment" check that cloudflare does
but may not bypass more complex browser checks. In some cases it may be a
good idea to set the referrer using a request taken from the browser as a
reference.
