FROM node:19-alpine
WORKDIR /frontend
RUN apk update && apk add git && apk add vim
RUN npm install -g npm@latest && npm install uuid create-next-app @storybook/react swr @mui/material @mui/styled-engine-sc styled-components @fontsource/roboto @mui/icons-material