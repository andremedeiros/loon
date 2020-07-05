require 'test_helper'

class TestCd < Loon::Test
  def test_that_we_get_a_finisher
    loon %w(cd andremedeiros/loon)

    assert_status 0
    assert_finalizer 'chdir', "#{ENV['HOME']}/src/github.com/andremedeiros/loon"
  end
end
