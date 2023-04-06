FROM node
WORKDIR /frontend
COPY entrypoint.sh /usr/bin
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT [ "entrypoint.sh" ]

CMD ["npm","run","dev"]