FROM microsoft/aspnet

COPY project.json /app/

WORKDIR /app

RUN ["dnu", "restore"]

COPY . /app

EXPOSE 5004

ENTRYPOINT sleep 9999999999 | dnx project.json kestrel
