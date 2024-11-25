
###################################
#Build stage
FROM golang:1.12-alpine3.9 AS build-env

ARG NXGIT_VERSION
ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata $TAGS"

#Build deps
RUN apk --no-cache add build-base git

#Setup repo
COPY . ${GOPATH}/src/go.khulnasoft.com/nxgit
WORKDIR ${GOPATH}/src/go.khulnasoft.com/nxgit

#Checkout version if set
RUN if [ -n "${NXGIT_VERSION}" ]; then git checkout "${NXGIT_VERSION}"; fi \
 && make clean generate build

FROM alpine:3.9
LABEL maintainer="maintainers@nxgit.io"

EXPOSE 22 3000

RUN apk --no-cache add \
    bash \
    ca-certificates \
    curl \
    gettext \
    git \
    linux-pam \
    openssh \
    s6 \
    sqlite \
    su-exec \
    tzdata

RUN addgroup \
    -S -g 1000 \
    git && \
  adduser \
    -S -H -D \
    -h /data/git \
    -s /bin/bash \
    -u 1000 \
    -G git \
    git && \
  echo "git:$(dd if=/dev/urandom bs=24 count=1 status=none | base64)" | chpasswd

ENV USER git
ENV NXGIT_CUSTOM /data/nxgit

VOLUME ["/data"]

ENTRYPOINT ["/usr/bin/entrypoint"]
CMD ["/bin/s6-svscan", "/etc/s6"]

COPY docker /
COPY --from=build-env /go/src/go.khulnasoft.com/nxgit/nxgit /app/nxgit/nxgit
RUN ln -s /app/nxgit/nxgit /usr/local/bin/nxgit
