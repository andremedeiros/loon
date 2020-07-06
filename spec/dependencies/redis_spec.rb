describe 'Dependencies' do
  describe 'Redis', dependency: true, redis: true do
    versions, latest = versions_for :redis

    versions.each do |version|
      it "installs version #{version} correctly" do
        test_dep 'redis', cmd: 'redis-cli --version', version: version, match: version
      end
    end

    it "installs version #{latest} as the default" do
      test_dep 'redis', cmd: 'redis-cli --version', match: latest
    end
  end
end
