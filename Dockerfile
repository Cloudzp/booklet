FROM node
ADD .  .
RUN chmod +x startup.sh
CMD startup.sh