require 'fileutils'
require 'open3'
require 'pathname'
require 'yaml'

require 'minitest/autorun'
require 'minitest/hooks/default'
require 'minitest/pride'

ROOT = Pathname(__FILE__).dirname.dirname
LOON = ROOT.join('loon')

module Assertions
  def assert_status(status)
    assert_equal status, @last_status
  end

  def refute_status(status)
    refute_equal status, @last_status
  end

  def assert_stderr(str)
    assert @last_stderr.include?(str), "Expected\n\n#{@last_stderr}\n\nto include\n\n#{str}\n"
  end

  def assert_stdout(str)
    assert @last_stdout.include?(str), "Expected\n\n#{@last_stdout}\n\nto include\n\n#{str}\n"
  end

  def assert_finalizer(type, content = nil)
    assert @last_finalizer.start_with?("#{type}:"), "Expected finalizer to be #{type}"
    assert @last_finalizer.end_with?(":#{content}"), "Expected finalizer to contain #{content} but instead it was #{@last_finalizer}" if content
  end

  def assert_path(path)
    assert Dir.exist?(path), "Expected #{path} to exist"
  end
end

module Loon
  class Test < Minitest::Test
    include Minitest::Hooks
    include Assertions

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
      cmd = args.prepend(LOON).flatten.map(&:to_s).join(' ')
      finalizer = Tempfile.new("finalizer")
      script = Tempfile.new("script")
      script.write <<~SH
        __integration_test() {
          exec 9>"#{finalizer.path}"
          #{cmd}
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
end
