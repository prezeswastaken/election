FROM oven/bun:1.0.6-debian

ENV NODE_ENV=development

WORKDIR /app

EXPOSE 3000

COPY package.json ./
COPY bun.lockb ./
RUN bun install

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000

ENV PATH=./node_modules/.bin/:$PATH

COPY . .
