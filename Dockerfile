FROM alpine:latest

ENV name=go-shopping-cart

COPY ./out/build/go-shopping-cart-linux /${name}

CMD [ "/go-shopping-cart" ]
