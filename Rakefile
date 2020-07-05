require "rake/testtask"

Rake::TestTask.new(:test) do |t|
  t.libs << "integration"
  t.test_files = FileList["integration/**/*_test.rb"]
end

task :default => :test

