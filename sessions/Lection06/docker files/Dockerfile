#hub.masterofcode.com/ruby2.2:stable
FROM awesome8585/devopscourse2016_ruby2.3.1_passenger_latest

MAINTAINER Vlad Bolibruk "v.bolibruk@gmail.com"

# Set correct environment variables.
ENV HOME /root


# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]


ADD ./ci/moc_stage/docker/authorized_keys /root/.ssh/authorized_keys
RUN echo 'key' >> /root/.ssh/authorized_keys
RUN chmod 700 /root/.ssh/authorized_keys

RUN rm -f /etc/service/sshd/down



#RAILS setup

#USER root
#tell not to install docs with gems
RUN touch /root/.gemrc && echo "gem: --no-ri --no-rdoc" >> /root/.gemrc
  

# Start Nginx / Passenger
RUN rm -f /etc/service/nginx/down

# Enable the Redis service.
RUN rm -f /etc/service/redis/down

# Add the nginx site and config
ADD ./nginx.conf /etc/nginx/nginx.conf
ADD ./webapp.conf /etc/nginx/sites-enabled/webapp.conf
ADD ./rails-env.conf /etc/nginx/main.d/rails-env.conf

#Unlock app account
RUN passwd -u app

# Prepare folders

#add authorized_keys
#ADD ./ci/moc_stage/docker/authorized_keys /home/app/.ssh/authorized_keys

RUN touch /home/app/.gemrc  && echo "gem: --no-ri --no-rdoc" >> /home/app/.gemrc 

# Install bundle of gems
WORKDIR /tmp
ADD Gemfile /tmp/Gemfile
ADD Gemfile.lock /tmp/Gemfile.lock
RUN bundle install --jobs 20 --retry 5 



RUN mkdir -p /home/app/webapp/current
ENV APP_HOME /home/app/webapp/current
#add database.yml and secrets.yml
#ADD ./ci/moc_stage/docker/database.yml /home/app/webapp/current/config/database.yml
#ADD ./ci/moc_stage/docker/secrets.yml /home/app/webapp/current/config/secrets.yml

# Add the rails app
ADD . $APP_HOME
WORKDIR $APP_HOME

RUN chown -R app:app /home/app/webapp
RUN sudo -u app RAILS_ENV=stage bundle exec rake assets:precompile
#RUN sudo -u app RAILS_ENV=development bundle exec rake assets:email_asset


# Enable the Sidekiq service up to 4 processes
RUN rm /etc/service/sidekiq0/down
#RUN rm /etc/service/sidekiq1/down
#RUN rm /etc/service/sidekiq2/down
#RUN rm /etc/service/sidekiq3/down


#enable faye
#RUN mkdir /etc/service/faye && touch /etc/service/faye/down
#ADD faye.sh /etc/service/faye/run
#RUN rm /etc/service/faye/down


#enable memcached
#RUN rm /etc/service/memcached/down



#RUN chmod 400 /etc/nginx/ssl/server.key
#RUN chown root:root /etc/nginx/ssl/server.key

RUN rm -rf $APP_HOME/ci
#RUN rm -rf $APP_HOME/docker
#RUN rm -rf $APP_HOME/ansible

#Expose port for nginx, postgresql, sshd
EXPOSE 8080 8443  5432 2022 6379 9292
# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*