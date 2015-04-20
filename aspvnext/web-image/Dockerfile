FROM microsoft/aspnet

COPY project.json /app/

WORKDIR /app

RUN ["kpm", "restore"]

COPY . /app

EXPOSE 5004

ENTRYPOINT ["k", "kestrel"]