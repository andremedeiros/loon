describe 'Dependencies' do
  describe 'Postgres', dependency: true, postgres: true do
    versions, latest = versions_for :postgresql

    versions.each do |version|
      it "installs version #{version} correctly" do
        test_dep 'postgresql', cmd: 'psql --version', version: version, match: version
      end
    end

    it "installs version #{latest} as the default" do
      test_dep 'postgresql', cmd: 'psql --version', match: latest
    end
  end
end
