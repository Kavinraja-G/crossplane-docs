FROM gcr.io/distroless/static-debian12

LABEL org.opencontainers.image.source="https://github.com/Kavinraja-G/crossplane-docs" \
      org.opencontainers.image.title="crossplane-docs" \
      org.opencontainers.image.description="Docs (XDocs) generator for Crossplane" \
      org.opencontainers.image.licenses="Apache-2.0"

# Copy the pre-built binary from GoReleaser
COPY crossplane-docs /usr/local/bin/crossplane-docs

ENTRYPOINT ["/usr/local/bin/crossplane-docs"]
CMD ["--help"]
