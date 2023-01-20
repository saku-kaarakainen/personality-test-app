# Create docker image for frontend
FROM node:16-alpine as node_builder

ENV APP_FOLDER "/personality-test-app"

COPY app "${APP_FOLDER}"

WORKDIR "${APP_FOLDER}"

# TODO: Use production build instead
RUN npm install

EXPOSE 3000

CMD ["npm", "run", "start"]
