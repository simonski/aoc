FROM alpine
# RUN apk update
EXPOSE 8000
COPY aoc_linux /aoc
ADD api/ /html
CMD [ "/aoc", "server", "-p", "8000", "-fs", "/html" ]

