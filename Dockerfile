FROM golang:1.19-alpine as builder

# forces using Go modules even if the project is in your GOPATH. Requires go.mod to work.
ENV GO111MODULE=on

# set the go api folder
#ENV APP_HOME /go/src/github.com/saku-kaarakainen/personality-test-app/api
ENV API_FOLDER "/personality-test-api"

# copy backend files to /personality-test-api - folder 
COPY api "${API_FOLDER}"

WORKDIR "${API_FOLDER}"

# https://stackoverflow.com/questions/66356034/what-is-the-difference-between-go-get-command-and-go-mod-download-command
# go mod download does not add new requirements or update existing requirements 
# RUN go mod download

# -o: output file name
RUN go build -o personality-test-api .

EXPOSE 8080

# call docker run . (in api folder)
CMD [ "go", "run", "." ]

# docker build -t personality-test-api .
# docker run -it --rm -p 8080:8080 personality-test-api