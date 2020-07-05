describe 'Dependencies' do
  describe 'Memcached', dependency: true, memcached: true do
    def test_memcached_dep(version: nil, match:)
      dep = if version
              {'memcached' => version}
            else
              'memcached'
            end

      with_payload(deps: dep) do |project|
        loon %w(up), dir: project
        loon %w(exec memcached --version), dir: project

        assert_stderr_empty
        assert_stdout match
        assert_status 0
      end
    end

    %w(1.6.5 1.6.6).each do |version|
      it "installs version #{version} correctly" do
        test_memcached_dep version: version, match: version
      end
    end

    it 'installs version 1.6.6 as the default' do
      test_memcached_dep match: '1.6.6'
    end
  end
end
