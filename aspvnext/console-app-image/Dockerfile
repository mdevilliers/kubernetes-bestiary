FROM microsoft/aspnet

COPY project.json /app/

WORKDIR /app

RUN ["dnu", "restore"]

COPY . /app

ENTRYPOINT dnx project.json run
