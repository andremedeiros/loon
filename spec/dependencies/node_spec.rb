describe 'Dependencies' do
  describe 'Node.JS', dependency: true, node: true do
    versions, latest = versions_for :node

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'node', version: version, match: version
      end
    end

    it "installs v#{latest} as the default" do
      test_dep 'node', match: latest
    end
  end
end
