require 'test_helper'

class TestDoctor < Loon::Test
  def test_that_it_shows_error_on_sudo
    with_command_mock('sudo', 'exit 1') do
      loon %w(doctor)
    end

    assert_status 0
    assert_stderr "sudo not enabled for this user"
  end

  def test_that_it_shows_error_on_nix
    with_environment(path: '') do
      loon %w(doctor)
    end

    assert_status 0
    assert_stderr "cannot find nix utility: nix"
  end

  def test_it_shows_no_errors_when_all_is_good
    loon %(doctor)

    assert_status 0
    assert_stdout "You're all good!"
  end
end
