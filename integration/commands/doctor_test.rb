require_relative '../test_helper'

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
end
