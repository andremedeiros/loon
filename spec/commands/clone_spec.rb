describe 'Commands' do
  describe 'clone', command: true do
    around(:each) do |example|
      with_environment(home: Dir.mktmpdir) do
        example.run
      end
    end

    it 'should respect source tree setting' do
      with_config(source_tree: '$HOME/{owner}/{name}') do
        loon %w(clone andremedeiros/ruby-demo)

        assert_stderr_empty
        assert_path "#{ENV['HOME']}/andremedeiros/ruby-demo/.git"
        assert_status 0
      end
    end

    it 'should check out a repo' do
      loon %w(clone andremedeiros/ruby-demo)

      assert_stderr_empty
      assert_path "#{ENV['HOME']}/src/github.com/andremedeiros/ruby-demo/.git"
      assert_status 0
    end

    it 'should emit a finisher' do
      loon %w(clone andremedeiros/ruby-demo)

      assert_stderr_empty
      assert_finalizer 'chdir', "#{ENV['HOME']}/src/github.com/andremedeiros/ruby-demo"
      assert_status 0
    end
  end
end
