FROM golang

MAINTAINER Linda Zhou 

# from byoi build
ARG URL
ARG PAAS_BUILD_ID
ENV URL=${URL} \
    PAAS_BUILD_ID=${PAAS_BUILD_ID}
WORKDIR $GOPATH/src/lhmzhou/level-seven-matcha

RUN useradd -u 1000 -U -d $GOPATH/src app  && chown -R 1000:1000 $GOPATH

# do anything requiring ROOT above here
USER 1000
ENTRYPOINT ./app
ADD . .
RUN go build -o app -ldflags "-X main.appVersion=$(git rev-parse HEAD)"
