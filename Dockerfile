# Create docker image for backend
FROM golang:1.19-alpine as go_builder

# forces using Go modules even if the project is in your GOPATH. Requires go.mod to work.
ENV GO111MODULE=on

# set the go api folder
ENV API_FOLDER "/personality-test-api"

# copy backend files to /personality-test-api - folder 
COPY api "${API_FOLDER}"

WORKDIR "${API_FOLDER}"

RUN go build -o personality-test-api .

EXPOSE 8080

# call docker run . (in api folder)
CMD [ "go", "run", "." ]

# Create docker image for frontend
FROM node:16-alpine as node_builder

ENV APP_FOLDER "/personality-test-app"

COPY app "${APP_FOLDER}"

WORKDIR "${APP_FOLDER}"

RUN npm install

EXPOSE 3000

CMD ["npm", "run", "start"]


