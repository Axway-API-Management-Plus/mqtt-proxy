FROM node
COPY package.json .
RUN npm install
COPY . .
CMD [ "npm" , "run", "test"]
