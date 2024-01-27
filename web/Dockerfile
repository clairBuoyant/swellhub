FROM node:21-slim

WORKDIR /app
ADD . /app

ENV NODE_PATH=/web/node_modules
ENV PATH=$PATH:/web/node_modules/.bin

EXPOSE 8000
EXPOSE 35729

ENTRYPOINT ["/bin/bash", "/app/scripts/run"]
CMD ["start"]
