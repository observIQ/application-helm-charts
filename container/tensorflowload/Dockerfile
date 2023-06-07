FROM curlimages/curl:8.1.1
USER root
COPY load.sh /load.sh
RUN \
    chmod +x /load.sh && \
    chown curl_user: /load.sh
USER curl_user
ENTRYPOINT [ "/load.sh" ]
