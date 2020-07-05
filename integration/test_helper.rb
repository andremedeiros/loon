require 'minitest/autorun'
require 'minitest/pride'
require 'open3'
require 'pathname'

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
end

module Loon
  class Test < Minitest::Test
    include Assertions

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
      out, err, status = Open3.capture3(ENV, cmd)

      @last_stdout = out
      @last_stderr = err
      @last_status = status.exitstatus
    end
  end
end
