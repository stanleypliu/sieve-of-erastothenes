FROM ruby:2.6.5

WORKDIR /sieve-of-eratosthenes

COPY . .

RUN gem install benchmark-ips

CMD ["ruby", "sieve_benchmark.rb"]