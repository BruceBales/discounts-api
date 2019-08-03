from ubuntu:14.04

EXPOSE 8080

ADD ./bin/discounts /bin/discounts

CMD ["/bin/discounts"]