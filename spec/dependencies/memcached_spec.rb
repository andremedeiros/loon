describe 'Dependencies' do
  describe 'Memcached', dependency: true, memcached: true do
    versions, latest = versions_for :memcached

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'memcached', version: version, match: version
      end
    end

    it "installs v#{latest} as the default" do
      test_dep 'memcached', match: latest
    end
  end
end
