#!/bin/bash

# set -a  # Automatically export all variables
# source .env.dev
# set +a  # Disable automatic exporting

run_tailwind() {
    ./tailwindcss -i tailwind.css -o static/css/styles.css --watch
}

run_server() {
    air -
}

(trap 'kill' SIGINT; run_server & run_tailwind)
