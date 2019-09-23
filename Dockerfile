FROM node
RUN cnpm install gitbook-cli -g \
&& gitbook -V \
ADD .  .
RUN chmod +x startup.sh
CMD startup.sh