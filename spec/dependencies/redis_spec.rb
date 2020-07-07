describe 'Dependencies' do
  describe 'Redis', dependency: true, redis: true do
    versions, default = versions_for :redis

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'redis', cmd: 'redis-cli --version', version: version, match: version
      end
    end

    it "installs v#{default} as the default" do
      test_dep 'redis', cmd: 'redis-cli --version', match: default
    end
  end
end
