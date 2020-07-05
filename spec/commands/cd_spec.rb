describe 'Commands' do
  describe 'cd' do
    it 'should emit a finisher' do
      loon %w(cd andremedeiros/loon)

      assert_status 0
      assert_finalizer 'chdir', "#{ENV['HOME']}/src/github.com/andremedeiros/loon"
    end
  end
end
