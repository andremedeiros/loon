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

        assert_status 0
        assert_path "#{ENV['HOME']}/andremedeiros/ruby-demo/.git"
      end
    end

    it 'should check out a repo' do
      loon %w(clone andremedeiros/ruby-demo)

      assert_status 0
      assert_path "#{ENV['HOME']}/src/github.com/andremedeiros/ruby-demo/.git"
    end

    it 'should emit a finisher' do
      loon %w(clone andremedeiros/ruby-demo)

      assert_status 0
      assert_finalizer 'chdir', "#{ENV['HOME']}/src/github.com/andremedeiros/ruby-demo"
    end
  end
end
