FROM scratch
LABEL authors="guillaume"

COPY oncall-to-prowl /

ENTRYPOINT ["/oncall-to-prowl"]