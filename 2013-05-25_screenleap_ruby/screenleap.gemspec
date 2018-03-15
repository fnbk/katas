# -*- encoding: utf-8 -*-
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'screenleap/version'

Gem::Specification.new do |gem|
  gem.name          = "screenleap"
  gem.version       = Screenleap::VERSION
  gem.authors       = ["Florian Boehmak"]
  gem.email         = ["florian.boehmak@googlemail.com"]
  gem.description   = %q{Simple screenleap API wrapper.}
  gem.summary       = %q{Simple screenleap API wrapper.}
  gem.homepage      = "https://github.com/fnbk/screenleap"

  gem.files         = `git ls-files`.split($/)
  gem.executables   = gem.files.grep(%r{^bin/}).map{ |f| File.basename(f) }
  gem.test_files    = gem.files.grep(%r{^(test|spec|features)/})
  gem.require_paths = ["lib"]

  gem.add_development_dependency('rspec', '2.13.0')
  gem.add_development_dependency('vcr', '2.5.0')
  gem.add_development_dependency('webmock', '1.11.0')
  gem.add_development_dependency('pry', '0.9.12.2')
  gem.add_dependency('activesupport', '>=3.0.0')
end
