require_relative '../test_helper'

class TestClone < Loon::Test
  # Ensure we run these tests in temporary home folders
  def around
    with_environment(home: Dir.mktmpdir) do
      super
    end
  end

  def test_that_clone_respects_source_tree
    with_config(source_tree: '$HOME/{owner}/{name}') do
      loon %w(clone andremedeiros/ruby-demo)

      assert_status 0
      assert_path "#{ENV['HOME']}/andremedeiros/ruby-demo/.git"
    end
  end

  def test_that_repo_gets_checked_out
    loon %w(clone andremedeiros/ruby-demo)

    assert_status 0
    assert_path "#{ENV['HOME']}/src/github.com/andremedeiros/ruby-demo/.git"
  end

  def test_that_we_get_a_finisher
    loon %w(clone andremedeiros/ruby-demo)

    assert_status 0
    assert_finalizer 'chdir', "#{ENV['HOME']}/src/github.com/andremedeiros/ruby-demo"
  end
end
