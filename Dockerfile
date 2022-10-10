FROM golang:1.19

ENV UID=10000
ENV GID=10000

RUN addgroup \
    --g "$GID" \
    "$USER"


RUN adduser \
    --disabled-password \
    --gecos ""\
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"


COPY --chown=10000:10000 stock-tracker /

EXPOSE 7443

USER "$UID"

ENTRYPOINT["/stock-tracker"]
