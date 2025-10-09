FROM node:21-alpine
WORKDIR /city_letterbox
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY . .
EXPOSE 5173
CMD ["npm", "run", "dev"]