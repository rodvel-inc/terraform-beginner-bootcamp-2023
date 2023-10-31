# Terraform Beginner Bootcamp 2023 - Week 2

### Working with Ruby

## Bundler

It is a packet manager for Ruby.

It is used to install Ruby packages, called "gems".

### Install Gems

You need to create a gemfile and define your gems in there.

```rb
source "https://rubygems.org"

gem 'sinatra'
gem 'rake'
gem 'pry'
gem 'puma'
gem 'activerecord'
```

Then, you need to run the `bundle install` command.

This will install the gems globally on the system.

A gemfile.lock file will be used to lock down the versions of the gems being used in a project.

### Executing Ruby scripts in the context of bundler

We have to use `bundle exec` to tell future Ruby scripts to use the gems already installed. This way, future scripts will have a context to work in.

### Using Sinatra

Sinatra is a micro web framework for Ruby. You can build web apps with Sinatra.

It is a good option for development environments. You can create a web appp in a single file.

https://sinatrarb.com/

## Terratowns Mock Server

### Running the web server

We can run the web server by executing the following commands:

```rb
bundle install
bundle exec ruby server.rb
```

All of the code of our server is stored in the `server.rb` file.




