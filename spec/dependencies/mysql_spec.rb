describe 'Dependencies' do
  describe 'MySQL', dependency: true, mysql: true do
    versions, default = versions_for :mysql

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'mysql', version: version, match: version
      end
    end

    it "installs v#{default} as the default" do
      test_dep 'mysql', match: default
    end
  end
end
