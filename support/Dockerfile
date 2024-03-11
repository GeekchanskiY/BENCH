FROM --platform=$BUILDPLATFORM python:3.10-alpine AS builder
EXPOSE 3003
WORKDIR /app 
COPY . .
RUN pip3 install -r requirements.txt --no-cache-dir

ENTRYPOINT ["python3"] 
CMD ["/app/app/manage.py", "runserver", "0.0.0.0:3003"]

FROM builder as dev-envs
RUN <<EOF
apk update
apk add git
EOF

RUN <<EOF
addgroup -S docker
adduser -S --shell /bin/bash --ingroup docker vscode
EOF
# install Docker tools (cli, buildx, compose)
COPY --from=gloursdocker/docker / /
CMD ["/app/app/manage.py", "runserver", "0.0.0.0:3003"]