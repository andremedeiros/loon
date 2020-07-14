describe 'Commands' do
  describe 'exec', command: true do
    it 'should execute things in the project environment' do
      with_payload do
        loon %w(exec env)
        assert_stdout %r(LOON_PROJECT_IP=\d+.\d+.\d+.\d+)
      end
    end
  end
end

