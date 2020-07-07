describe 'Commands' do
  describe 'cd', command: true do
    it 'should emit a finisher' do
      loon %w(cd andremedeiros/loon)

      assert_finalizer 'chdir', "#{ENV['HOME']}/src/github.com/andremedeiros/loon"
      assert_status 0
    end
  end
end
