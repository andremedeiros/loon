require 'byebug'
require 'tempfile'
require 'fileutils'
require 'open3'
require 'pathname'
require 'yaml'

module Assertions
  def assert_status(status)
    expect(status).to eq(@last_status)
  end

  def refute_status(status)
    expect(status).not_to eq(@last_status)
  end

  def assert_stderr(str)
    expect(@last_stderr).to include(str)
  end

  def assert_stderr_empty
    expect(@last_stderr).to be_empty
  end

  def assert_stdout(str)
    expect(@last_stdout).to include(str)
  end

  def assert_finalizer(type, content = nil)
    expect(@last_finalizer).to start_with(type)
    expect(@last_finalizer).to end_with(content) if content
  end

  def assert_path(path)
    expect(Dir.exist?(path)).to be_truthy
  end
end

module Helpers
  ROOT = Pathname(__FILE__).dirname.dirname
  LOON = ROOT.join('loon')

  def with_payload(name: "Test", url: "Test", deps: [], tasks: {})
    deps = deps.is_a?(Array) ? deps : [deps]

    yml = YAML.dump({
      'name' => name,
      'url' => url,
      'deps' => deps,
      'tasks' => tasks,
    })

    Dir.mktmpdir do |tmpdir|
      File.open(File.join(tmpdir, 'loon.yml'), 'w') do |f|
        f.write(yml)
      end
      yield tmpdir
    end
  end

  def with_config(cfg)
    home = Dir.mktmpdir
    with_environment(home: home) do
      FileUtils.mkdir_p File.join(home, '.config', 'loon')
      File.open(File.join(home, '.config', 'loon', 'config.yml'), 'w') do |f|
        cfg.each { |k, v| f.write("#{k}: #{v}\n") }
      end
      yield
    end
  ensure
    FileUtils.remove_entry home
  end

  def with_environment(env)
    previous_env = {}
    env.each do |k, v|
      k = k.to_s.upcase
      previous_env[k] = ENV[k]
      ENV[k] = v
    end
    yield
  ensure
    previous_env.each do |k, v|
      ENV[k] = v
    end
  end

  def with_env_path(dir)
    with_environment(path: "#{dir}:#{ENV['PATH']}") do
      yield
    end
  end

  def with_tmpdir
    dir = Dir.mktmpdir
    yield dir
  ensure
    FileUtils.remove_entry dir
  end

  def with_command_mock(cmd, mock)
    with_tmpdir do |dir|
      cmd = File.join(dir, cmd)
      open(cmd, "w") { |f| f.write(mock) }
      FileUtils.chmod("+x", cmd)
      with_env_path(dir) { yield }
    end
  end

  def loon(*args)
    opts = args.last.is_a?(Hash) ? args.pop : {}
    cmd = args.prepend(LOON).flatten.map(&:to_s).join(' ')
    finalizer = Tempfile.new("finalizer")
    script = Tempfile.new("script")
    script.write <<~SH
      __integration_test() {
        exec 9>"#{finalizer.path}"
        pushd #{opts[:dir] || Dir.pwd} >/dev/null
        #{cmd}
        popd >/dev/null
        ret=$?
        exec 9<&-
        return "$ret"
      }
      __integration_test
    SH
    script.close
    FileUtils.chmod("+x", script.path)
    out, err, status = Open3.capture3(ENV, script.path)
    @last_stdout = out
    @last_stderr = err
    @last_status = status.exitstatus
    @last_finalizer = IO.read(finalizer.path).chomp
  ensure
    File.unlink(script.path)
    File.unlink(finalizer.path)
  end
end

RSpec.configure do |config|
  config.expect_with :rspec do |expectations|
      expectations.include_chain_clauses_in_custom_matcher_descriptions = true
  end

  config.mock_with :rspec do |mocks|
    mocks.verify_partial_doubles = true
  end

  config.shared_context_metadata_behavior = :apply_to_host_groups
  config.filter_run_when_matching :focus
  config.warnings = true
  config.order = :random

  config.include Assertions
  config.include Helpers

  # Bundler does... things to the environment, so we want to get the original
  # environment before running tests so that the executions aren't tainted.
  config.around(:each) do |example|
    Bundler.with_original_env { example.run }
  end
end