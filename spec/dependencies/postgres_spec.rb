describe 'Dependencies' do
  describe 'Postgres', dependency: true, postgres: true do
    versions, default = versions_for :postgresql

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'postgresql', cmd: 'psql --version', version: version, match: version
      end
    end

    it "installs v#{default} as the default" do
      test_dep 'postgresql', cmd: 'psql --version', match: default
    end
  end
end
