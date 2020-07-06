describe 'Dependencies' do
  describe 'MySQL', dependency: true, mysql: true do
    versions, latest = versions_for :mysql

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'mysql', version: version, match: version
      end
    end

    it "installs v#{latest} as the default" do
      test_dep 'mysql', match: latest
    end
  end
end
