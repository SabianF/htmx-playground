# Set environment variables for bash
export $(grep -v '^#' .env | xargs -d '\n')

# Watch .templ file changes
templ generate --watch\
 --proxy=http://localhost:${APP_PORT}\
 --proxyport=${TEMPL_PROXY_PORT}\
 --open-browser=false\
 --proxybind="0.0.0.0" & \
# Watch .go file changes
air
