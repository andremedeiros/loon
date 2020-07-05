require 'test_helper'

class TestRuby < Loon::Test
  def test_ruby_versions
    dependency_test :ruby
    %w(2.6.6 2.7.1).each do |version|
      test_ruby_dep version: version, match: version
    end
  end

  def test_default
    dependency_test :ruby
    test_ruby_dep match: '2.7.1'
  end

  private

  def test_ruby_dep(version: nil, match:)
    dep = if version
            {'ruby' => version}
          else
            'ruby'
          end

    with_payload(deps: dep) do
      loon %w(exec ruby --version)

      assert_status 0
      assert_stdout match
    end
  end
end
