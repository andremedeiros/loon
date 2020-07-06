describe 'Dependencies' do
  describe 'Crystal', dependency: true, crystal: true do
    versions, latest = versions_for :crystal

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'crystal', version: version, match: version
      end
    end

    it "installs v#{latest} as the default" do
      test_dep 'crystal', match: latest
    end
  end
end
