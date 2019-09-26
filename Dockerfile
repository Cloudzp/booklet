FROM node
WORKDIR /booklet
ADD .  /booklet
RUN npm install -g gitbook-cli \
&& chmod +x startup.sh
EXPOSE 8080
CMD ./startup.sh