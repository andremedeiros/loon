describe 'Commands' do
  describe 'doctor', command: true do
    it 'should show error when sudo fails' do
      with_command_mock('sudo', 'exit 1') do
        loon %w(doctor)
      end

      assert_status 0
      assert_stderr 'sudo not enabled for this user'
    end

    it 'should show error when nix is not installed' do
      with_environment(path: '') do
        loon %w(doctor)
      end

      assert_status 0
      assert_stderr 'cannot find nix utility: nix'
    end

    it 'should show no errors when all is well' do
      loon %(doctor)

      assert_status 0
      assert_stdout "You're all good!"
    end
  end
end

