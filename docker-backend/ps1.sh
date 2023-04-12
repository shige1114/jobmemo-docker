#!/bin/bash
set -e
toch ~/.bash_profile
echo 'export PS1="\w\n\[\e[35m\]\u\[\e[0m\]:\[\e[32m\]\W\[\e[0m\]$ "' >> ~/.bash_profile
source ~/.bash_profile

exec "$@"