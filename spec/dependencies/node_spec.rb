describe 'Dependencies' do
  describe 'Node.JS', dependency: true, node: true do
    versions, latest = versions_for :node

    versions.each do |version|
      it "installs version #{version} correctly" do
        test_dep 'node', version: version, match: version
      end
    end

    it "installs version #{latest} as the default" do
      test_dep 'node', match: latest
    end
  end
end
