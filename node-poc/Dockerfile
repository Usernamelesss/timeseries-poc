FROM node:22-slim

COPY . /app

WORKDIR /app

RUN npm install

RUN npx tsc

#ENTRYPOINT ["npx", "esrun", "/app/src/index.ts"]
ENTRYPOINT ["node", "/app/dist/index.js"]