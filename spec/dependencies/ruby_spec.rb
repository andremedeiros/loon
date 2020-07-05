describe 'Dependencies' do
  describe 'Ruby' do
    def test_ruby_dep(version: nil, match:)
      dep = if version
              {'ruby' => version}
            else
              'ruby'
            end

      with_payload(deps: dep) do |project|
        loon %w(up), dir: project
        loon %w(exec ruby --version), dir: project

        assert_stderr_empty
        assert_stdout match
        assert_status 0
      end
    end

    %w(2.6.6 2.7.1).each do |version|
      it "installs version #{version} correctly", dependency: 'ruby' do
        test_ruby_dep version: version, match: version
      end
    end

    it 'installs version 2.7.1 as the default', dependency: 'ruby' do
      test_ruby_dep match: '2.7.1'
    end
  end
end
