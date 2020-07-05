require 'test_helper'

class TestRuby < Loon::Test
  def test_versions
    %w(2.6.6 2.7.1).each do |version|
      with_payload(deps: {'ruby' => version}) do
        loon %w(exec ruby --version)

        assert_status 0
        assert_stdout version
      end
    end
  end
end

